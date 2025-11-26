import { describe, vi, it, expect, beforeEach, afterEach } from 'vitest'
import { mountSuspended, mockNuxtImport } from '@nuxt/test-utils/runtime'
import Register from '~/pages/register.vue'
import AuthForm from '@nuxt/ui/components/AuthForm.vue'

const v1AuthRegisterPostMock = vi.fn()
const toastAddMock = vi.fn()

describe('Register Page', () => {
  beforeEach(() => {
    const { useGeneratedClientMock, useToastMock } = vi.hoisted(() => {
      return {
        useGeneratedClientMock: vi.fn(() => {
          return {
            api: { v1AuthRegisterPost: v1AuthRegisterPostMock }
          }
        }),
        useToastMock: vi.fn(() => {
          return {
            add: toastAddMock
          }
        })
      }
    })
    mockNuxtImport('useGeneratedClient', () => {
      return useGeneratedClientMock
    })
    mockNuxtImport('useToast', () => {
      return useToastMock
    })
  })

  afterEach(() => {
    vi.clearAllMocks()
    vi.clearAllMocks()
  })

  it('Success', async () => {
    v1AuthRegisterPostMock.mockResolvedValueOnce({
      message: 'Success Register'
    })
    const page = await mountSuspended(Register, { route: '/register' })
    const UAuthForm = page.findComponent(AuthForm)
    UAuthForm.vm.$emit('submit', {
      data: {
        email: 'operation@test.com',
        password: 'password',
        confirmPassword: 'password'
      }
    })
    await page.vm.$nextTick()
    expect(v1AuthRegisterPostMock).toBeCalledWith({
      v1AuthRegisterPostRequest: {
        email: 'operation@test.com',
        password: 'password',
        confirmPassword: 'password'
      }
    })
    expect(toastAddMock).toBeCalledWith({
      title: 'Success Register',
      description: 'Please login using your new account'
    })
  })

  it('Error', async () => {
    v1AuthRegisterPostMock.mockRejectedValueOnce({
      code: '500',
      message: 'error'
    })
    const page = await mountSuspended(Register, { route: '/register' })
    const UAuthForm = page.findComponent(AuthForm)
    UAuthForm.vm.$emit('submit', {
      data: {
        email: 'operation@test.com',
        password: 'password',
        confirmPassword: 'password'
      }
    })
    await page.vm.$nextTick()
    expect(v1AuthRegisterPostMock).toBeCalledWith({
      v1AuthRegisterPostRequest: {
        email: 'operation@test.com',
        password: 'password',
        confirmPassword: 'password'
      }
    })
    await page.vm.$nextTick()
    expect(toastAddMock).toBeCalledWith({
      title: '500',
      description: 'error',
      color: 'error'
    })
  })
})
