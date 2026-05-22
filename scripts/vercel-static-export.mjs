/**
 * Maps docgen output (web/static/docs/{en,ru}/…) to Vercel public/ URL layout:
 *   /docs/…       ← en/
 *   /ru/docs/…    ← ru/
 *   /static/…     ← web/static/{css,js,img,fonts}
 */
import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const root = path.join(path.dirname(fileURLToPath(import.meta.url)), "..");
const staticRoot = path.join(root, "web", "static");
const docsRoot = path.join(staticRoot, "docs");
const publicRoot = path.join(root, "public");

const assetDirs = ["css", "js", "img", "fonts"];
const docArtifacts = ["search-index.json", "sitemap.xml", "registry-manifest.json"];

function rmrf(dir) {
  if (fs.existsSync(dir)) {
    fs.rmSync(dir, { recursive: true, force: true });
  }
}

function copyDir(src, dest) {
  fs.mkdirSync(path.dirname(dest), { recursive: true });
  fs.cpSync(src, dest, { recursive: true });
}

function copyFile(src, dest) {
  fs.mkdirSync(path.dirname(dest), { recursive: true });
  fs.copyFileSync(src, dest);
}

rmrf(publicRoot);
fs.mkdirSync(publicRoot, { recursive: true });

for (const dir of assetDirs) {
  const src = path.join(staticRoot, dir);
  if (!fs.existsSync(src)) {
    continue;
  }
  copyDir(src, path.join(publicRoot, "static", dir));
}

const enDocs = path.join(docsRoot, "en");
if (!fs.existsSync(enDocs)) {
  console.error(
    "vercel-static-export: missing web/static/docs/en — run `bun run docs:build` locally and commit web/static/docs/",
  );
  process.exit(1);
}
copyDir(enDocs, path.join(publicRoot, "docs"));

const ruDocs = path.join(docsRoot, "ru");
if (fs.existsSync(ruDocs)) {
  copyDir(ruDocs, path.join(publicRoot, "ru", "docs"));
}

for (const name of docArtifacts) {
  const src = path.join(docsRoot, name);
  if (fs.existsSync(src)) {
    copyFile(src, path.join(publicRoot, "docs", name));
  }
}

console.log("vercel-static-export: wrote", publicRoot);
