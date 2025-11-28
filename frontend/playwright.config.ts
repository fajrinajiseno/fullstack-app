import { defineConfig } from 'playwright/test'
import dotenv from 'dotenv'

dotenv.config({ path: '.env' })

export default defineConfig({
  testDir: 'tests/e2e',
  timeout: 30_000,
  retries: process.env.CI ? 2 : 0,

  use: {
    baseURL: process.env.E2E_BASE_URL || 'http://localhost:3000',
    trace: 'on-first-retry',
    screenshot: 'only-on-failure',
    headless: process.env.E2E_HEADLESS === 'true'
  },

  webServer: {
    command: 'pnpm run build && pnpm run preview',
    url: process.env.E2E_BASE_URL || 'http://localhost:3000',
    reuseExistingServer: !process.env.CI,
    timeout: 120_000
  }
})
