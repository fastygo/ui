/**
 * Run Playwright against an already-running dev server (no webServer boot).
 */
import { spawnSync } from "node:child_process";

const result = spawnSync("playwright", ["test"], {
  stdio: "inherit",
  env: { ...process.env, E2E_REUSE_SERVER: "1" },
  shell: process.platform === "win32",
});

process.exit(result.status ?? 1);
