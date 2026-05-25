/**
 * Copies built static docs into the Nu HTML validation snapshot folder.
 * Run after `bun run docs:build` so docgen executes only once in verify.
 */
import fs from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const root = path.resolve(__dirname, "..");
const src = path.join(root, "web", "static", "docs");
const dest = path.join(root, ".validate", "html-snapshots", "nu", "docs");

async function rmDir(dir) {
  try {
    await fs.rm(dir, { recursive: true, force: true });
  } catch {
    /* ignore */
  }
}

async function copyDir(from, to) {
  await fs.mkdir(to, { recursive: true });
  const entries = await fs.readdir(from, { withFileTypes: true });
  for (const entry of entries) {
    const srcPath = path.join(from, entry.name);
    const destPath = path.join(to, entry.name);
    if (entry.isDirectory()) {
      await copyDir(srcPath, destPath);
    } else if (entry.isFile()) {
      await fs.copyFile(srcPath, destPath);
    }
  }
}

try {
  await fs.access(src);
} catch {
  console.error("docs:snapshot: run `bun run docs:build` first");
  process.exit(1);
}

await rmDir(dest);
await copyDir(src, dest);
console.log(`docs:snapshot: copied ${src} -> ${dest}`);
