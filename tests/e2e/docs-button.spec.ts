import { test, expect } from "@playwright/test";
import AxeBuilder from "@axe-core/playwright";

test.describe("docs gallery", () => {
  test("button component page renders", async ({ page }) => {
    await page.goto("/docs/primitives/button");
    await expect(
      page.getByRole("main").getByRole("heading", { name: "Button", level: 1 }),
    ).toBeVisible();
    await expect(page.getByRole("button", { name: "Primary" }).first()).toBeVisible();
  });

  test("button docs pass axe (wireframe scope)", async ({ page }) => {
    await page.goto("/docs/primitives/button");
    const results = await new AxeBuilder({ page })
      .disableRules(["color-contrast", "list", "nested-interactive"])
      .analyze();
    const bad = results.violations.filter(
      (v) => v.impact === "critical" || v.impact === "serious",
    );
    expect(bad).toEqual([]);
  });
});
