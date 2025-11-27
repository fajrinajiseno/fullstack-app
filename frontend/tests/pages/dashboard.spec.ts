import { describe, vi, it, expect, beforeEach, afterEach } from 'vitest'
import { mountSuspended, mockNuxtImport } from '@nuxt/test-utils/runtime'
import Dashboard from '~/pages/dashboard.vue'

describe('Dashboard Page', () => {
  beforeEach(() => {
    const { useAuthStoreMock } = vi.hoisted(() => {
      return {
        useAuthStoreMock: vi.fn(() => {
          return {
            getUser: () => ({
              id: '1',
              email: 'operation@test.com',
              token: 'token'
            }),
            logout: () => {}
          }
        })
      }
    })
    mockNuxtImport('useAuthStore', () => {
      return useAuthStoreMock
    })
  })

  afterEach(() => {
    vi.resetAllMocks()
    vi.clearAllMocks()
  })

  it('Success Render', async () => {
    const page = await mountSuspended(Dashboard, { route: '/dashboard' })
    expect(page.text()).toContain('Hello operation@test.com, welcome back')
    page.unmount()
  })
})
