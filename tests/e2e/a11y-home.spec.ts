import { test, expect } from "@playwright/test";
import AxeBuilder from "@axe-core/playwright";

test.describe("a11y (axe-core)", () => {
  test("home shell has no serious or critical violations (wireframe scope)", async ({
    page,
  }) => {
    await page.goto("/");
    const results = await new AxeBuilder({ page })
      .disableRules(["color-contrast"])
      .analyze();
    const bad = results.violations.filter(
      (v) => v.impact === "critical" || v.impact === "serious",
    );
    expect(bad).toEqual([]);
  });
});
