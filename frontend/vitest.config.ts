import { defineConfig } from 'vitest/config'
import { defineVitestProject } from '@nuxt/test-utils/config'

export default defineConfig({
  test: {
    projects: [
      await defineVitestProject({
        test: {
          name: 'nuxt',
          include: ['tests/**/*.spec.ts', 'tests/**/*.test.ts'],
          exclude: ['tests/e2e/**', '**/*.e2e.{ts,js}'],
          environment: 'nuxt',
          globals: true
        }
      })
    ]
  }
})
