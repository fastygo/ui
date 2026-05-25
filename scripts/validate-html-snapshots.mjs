/**
 * Validates every *.html under .validate/html-snapshots/nu/ with W3C Nu HTML Checker (network).
 * Only that subtree is checked so draft outputs in html-snapshots/generated/ do not fail CI.
 * Skips with success when no HTML files are present.
 *
 * Uses native fetch (not html-validator/axios) — W3C blocks axios with Cloudflare 403.
 */
import { readdir, readFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const root = path.resolve(__dirname, "..");
const snapshotRoot = path.join(root, ".validate", "html-snapshots", "nu");
const NU_URL = "https://validator.w3.org/nu/?out=json";
const REQUEST_DELAY_MS = 1000;
const MAX_ATTEMPTS = 5;

const sleep = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

async function validateWithNu(data) {
  let lastStatus = 0;
  for (let attempt = 1; attempt <= MAX_ATTEMPTS; attempt++) {
    const response = await fetch(NU_URL, {
      method: "POST",
      headers: {
        "Content-Type": "text/html; charset=utf-8",
        "User-Agent": "FastyGoUI-validate/1.0",
      },
      body: data,
    });
    lastStatus = response.status;
    if (response.ok) {
      return response.json();
    }
    if ((response.status === 403 || response.status === 429) && attempt < MAX_ATTEMPTS) {
      await sleep(2000 * attempt);
      continue;
    }
    throw new Error(`Validator returned unexpected statuscode: ${response.status}`);
  }
  throw new Error(`Validator returned unexpected statuscode: ${lastStatus}`);
}

async function collectHTMLFiles(dir, acc = []) {
  let entries;
  try {
    entries = await readdir(dir, { withFileTypes: true });
  } catch (e) {
    if (e && e.code === "ENOENT") {
      return acc;
    }
    throw e;
  }
  for (const ent of entries) {
    const full = path.join(dir, ent.name);
    if (ent.isDirectory()) {
      await collectHTMLFiles(full, acc);
    } else if (ent.isFile() && ent.name.endsWith(".html")) {
      acc.push(full);
    }
  }
  return acc;
}

async function main() {
  const files = await collectHTMLFiles(snapshotRoot);
  if (files.length === 0) {
    console.log("validate-html-snapshots: no .html files under .validate/html-snapshots/nu — skip.");
    return;
  }

  const errors = [];
  let infoCount = 0;
  const sorted = files.sort();
  for (let i = 0; i < sorted.length; i++) {
    const file = sorted[i];
    const data = await readFile(file, "utf8");
    const rel = path.relative(root, file);
    try {
      const body = await validateWithNu(data);
      const messages = body.messages ?? [];
      infoCount += messages.filter((m) => (m.type || "").toLowerCase() === "info").length;
      const real = messages.filter((m) => {
        const t = (m.type || "").toLowerCase();
        return t === "error" || t === "non-document-error";
      });
      if (real.length > 0) {
        errors.push({ file: rel, messages: real });
      } else {
        console.log(`OK Nu HTML: ${rel}`);
      }
    } catch (e) {
      errors.push({ file: rel, error: String(e?.message || e) });
    }
    if (i + 1 < sorted.length) {
      await sleep(REQUEST_DELAY_MS);
    }
  }

  console.log(
    `validate-html-snapshots: ${sorted.length} file(s), ${infoCount} info (non-failing), ${errors.length} failed.`,
  );

  if (errors.length) {
    for (const err of errors) {
      console.error(`\nFAILED: ${err.file}`);
      if (err.messages) {
        for (const m of err.messages) {
          console.error(`  L${m.lastLine ?? "?"}:${m.lastColumn ?? "?"} ${m.message}`);
        }
      } else {
        console.error(`  ${err.error}`);
      }
    }
    process.exit(1);
  }
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
