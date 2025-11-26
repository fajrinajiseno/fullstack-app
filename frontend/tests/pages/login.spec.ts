import { describe, vi, it, expect, beforeEach, afterEach } from 'vitest'
import { mountSuspended, mockNuxtImport } from '@nuxt/test-utils/runtime'
import Login from '~/pages/login.vue'
import AuthForm from '@nuxt/ui/components/AuthForm.vue'

const apiLoginMock = vi.fn()
const toastAddMock = vi.fn()
const setUserMock = vi.fn()

describe('Login Page', () => {
  beforeEach(() => {
    const { useApiClientMock, useToastMock, useAuthStoreMock } = vi.hoisted(
      () => {
        return {
          useApiClientMock: vi.fn(() => {
            return {
              apiLogin: apiLoginMock
            }
          }),
          useToastMock: vi.fn(() => {
            return {
              add: toastAddMock
            }
          }),
          useAuthStoreMock: vi.fn(() => {
            return {
              getUser: () => null,
              setUser: setUserMock
            }
          })
        }
      }
    )
    mockNuxtImport('useApiClient', () => {
      return useApiClientMock
    })
    mockNuxtImport('useToast', () => {
      return useToastMock
    })
    mockNuxtImport('useAuthStore', () => {
      return useAuthStoreMock
    })
  })

  afterEach(() => {
    vi.resetAllMocks()
    vi.clearAllMocks()
  })

  it('Success', async () => {
    apiLoginMock.mockResolvedValueOnce({
      id: '1',
      email: 'operation@test.com',
      token: 'token'
    })
    const page = await mountSuspended(Login, { route: '/login' })
    const UAuthForm = page.findComponent(AuthForm)
    UAuthForm.vm.$emit('submit', {
      data: { email: 'operation@test.com', password: 'password' }
    })
    await page.vm.$nextTick()
    expect(apiLoginMock).toBeCalledWith('operation@test.com', 'password')
    expect(setUserMock).toBeCalledWith({
      id: '1',
      email: 'operation@test.com',
      token: 'token'
    })
    expect(toastAddMock).toBeCalledWith({
      title: 'success login'
    })
    page.unmount()
  })

  it('Error', async () => {
    apiLoginMock.mockRejectedValueOnce({
      code: '500',
      message: 'error'
    })
    const page = await mountSuspended(Login, { route: '/login' })
    const UAuthForm = page.findComponent(AuthForm)
    UAuthForm.vm.$emit('submit', {
      data: { email: 'operation@test.com', password: 'password' }
    })
    await page.vm.$nextTick()
    expect(apiLoginMock).toBeCalledWith('operation@test.com', 'password')
    await page.vm.$nextTick()
    expect(toastAddMock).toBeCalledWith({
      title: '500',
      description: 'error',
      color: 'error'
    })
    page.unmount()
  })
})
