import fs from "node:fs";
import path from "node:path";
import { defineConfig, devices } from "@playwright/test";

const root = process.cwd();
const serverBin = path.join(
  root,
  ".internal",
  "bin",
  process.platform === "win32" ? "server.exe" : "server",
);
const reuseOnly = process.env.E2E_REUSE_SERVER === "1";
const serverCommand = fs.existsSync(serverBin)
  ? serverBin
  : "go run ./cmd/server";

export default defineConfig({
  testDir: "tests/e2e",
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 1 : 0,
  reporter: [["list"]],
  use: {
    baseURL: "http://127.0.0.1:18081",
    trace: "on-first-retry",
  },
  projects: [{ name: "chromium", use: { ...devices["Desktop Chrome"] } }],
  webServer: reuseOnly
    ? undefined
    : {
        command: serverCommand,
        cwd: root,
        url: "http://127.0.0.1:18081/healthz",
        reuseExistingServer: !process.env.CI,
        timeout: 120_000,
        env: {
          ...process.env,
          APP_BIND: "127.0.0.1:18081",
        },
      },
});
