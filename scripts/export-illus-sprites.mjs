/**
 * Export transparent PNG sprite sheets for docs index illustrations.
 *
 * Requires built CSS (bun run build:css) and a running server or auto-starts one.
 *
 * Output:
 *   web/static/img/docs-illus/primitives.png
 *   web/static/img/docs-illus/components.png
 */
import fs from "node:fs";
import path from "node:path";
import { spawn, spawnSync } from "node:child_process";
import { chromium } from "@playwright/test";

const root = process.cwd();
const outDir = path.join(root, "web", "static", "img", "docs-illus");
const baseURL = (process.env.ILLUS_EXPORT_BASE_URL ?? "http://127.0.0.1:8080").replace(
  /\/$/,
  "",
);
const deviceScaleFactor = Number(process.env.ILLUS_EXPORT_DPR ?? "2");

const sections = [
  { section: "primitives", file: "primitives.png" },
  { section: "components", file: "components.png" },
];

async function healthzOk(url) {
  try {
    const res = await fetch(`${url}/healthz`, { signal: AbortSignal.timeout(3000) });
    return res.ok;
  } catch {
    return false;
  }
}

function startServer(url) {
  const parsed = new URL(url);
  const bind = `${parsed.hostname}:${parsed.port || "8080"}`;
  const child = spawn("go", ["run", "./cmd/server"], {
    cwd: root,
    stdio: "pipe",
    env: { ...process.env, APP_BIND: bind },
    windowsHide: true,
  });

  return { child, bind };
}

function stopServer(child) {
  if (!child || child.exitCode !== null || child.signalCode) {
    return;
  }
  if (process.platform === "win32") {
    spawnSync("taskkill", ["/PID", String(child.pid), "/T", "/F"], {
      stdio: "ignore",
      windowsHide: true,
    });
    return;
  }
  child.kill("SIGTERM");
}

async function waitForServer(url, timeoutMs = 120_000) {
  const started = Date.now();
  while (Date.now() - started < timeoutMs) {
    if (await healthzOk(url)) {
      return;
    }
    await new Promise((resolve) => setTimeout(resolve, 500));
  }
  throw new Error(`server did not become ready at ${url}/healthz within ${timeoutMs}ms`);
}

async function exportSection(page, section, outFile) {
  const url = `${baseURL}/lab/docs-index-illus/export/${section}`;
  await page.goto(url, { waitUntil: "networkidle" });
  const grid = page.locator(".docs-illus-export-grid");
  await grid.waitFor({ state: "visible" });

  const cellCount = await page.locator(".docs-illus-export-cell").count();
  if (cellCount !== 26) {
    throw new Error(`expected 26 export cells for ${section}, got ${cellCount}`);
  }

  const box = await grid.boundingBox();
  if (!box) {
    throw new Error(`export grid has no bounding box for ${section}`);
  }

  await grid.screenshot({
    path: outFile,
    omitBackground: true,
  });

  console.log(
    `wrote ${path.relative(root, outFile)} (${Math.round(box.width)}×${Math.round(box.height)} css px, dpr=${deviceScaleFactor})`,
  );
}

async function main() {
  const cssPath = path.join(root, "web", "static", "css", "app.css");
  if (!fs.existsSync(cssPath)) {
    console.error("missing web/static/css/app.css — run: bun run build:css");
    process.exit(1);
  }

  fs.mkdirSync(outDir, { recursive: true });

  let server = null;
  let ownedServer = false;

  if (!(await healthzOk(baseURL))) {
    console.log(`starting server at ${baseURL} ...`);
    server = startServer(baseURL);
    ownedServer = true;
    server.child.on("exit", (code, signal) => {
      if (code && code !== 0) {
        console.error(`server exited with code ${code}${signal ? ` signal ${signal}` : ""}`);
      }
    });
    await waitForServer(baseURL);
  } else {
    console.log(`reusing server at ${baseURL}`);
  }

  const browser = await chromium.launch();
  const context = await browser.newContext({
    deviceScaleFactor,
    colorScheme: "light",
  });
  const page = await context.newPage();

  try {
    for (const { section, file } of sections) {
      await exportSection(page, section, path.join(outDir, file));
    }
  } finally {
    await browser.close();
    if (ownedServer && server?.child) {
      stopServer(server.child);
    }
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
