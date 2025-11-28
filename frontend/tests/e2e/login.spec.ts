import { test, expect } from '@playwright/test'

test('Login Page', async ({ page }) => {
  await page.goto('/dashboard')

  // unauthenticated will redirect to login
  await page.waitForURL('/login')

  await page.waitForTimeout(3000)

  await expect(page.locator('[data-slot="title"]')).toContainText('Login')
  await page.fill('[data-slot=input]', 'operation@test.com')
  await page.fill('[data-slot=password]', 'password')
  await page.click('button[type=submit]')

  await page.waitForURL('/dashboard')

  await expect(
    page.locator('[data-testid="dashboard-welcome-text"]')
  ).toContainText('Hello operation@test.com, welcome back')

  // wait inactivity
  await page.waitForTimeout(
    Number(process.env.PUBLIC_INACTIVITY_LIMIT || 900000) + 1000 // +1000 for buffer
  )

  // redirected to login
  await expect(page.locator('[data-slot="title"]')).toContainText('Login')
})
