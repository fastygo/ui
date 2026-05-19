/**
 * Validates every *.html under .validate/html-snapshots/nu/ with W3C Nu HTML Checker (network).
 * Only that subtree is checked so draft outputs in html-snapshots/generated/ do not fail CI.
 * Skips with success when no HTML files are present.
 */
import { createRequire } from "node:module";
import { readdir, readFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";

const require = createRequire(import.meta.url);
const htmlValidator = require("html-validator");

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const root = path.resolve(__dirname, "..");
const snapshotRoot = path.join(root, ".validate", "html-snapshots", "nu");

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
  for (const file of files.sort()) {
    const data = await readFile(file, "utf8");
    const rel = path.relative(root, file);
    try {
      const result = await htmlValidator({
        data,
        format: "json",
      });
      const body = typeof result === "string" ? JSON.parse(result) : result;
      const messages = body.messages ?? [];
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
  }

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
