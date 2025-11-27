export default defineNuxtRouteMiddleware(async (to) => {
  const auth = useAuthStore()
  let token = useCookie<string | null>(TOKEN_KEY).value
  if (import.meta.client) {
    token = auth.getUser()?.token || null
  }

  if (!token && to.path === '/dashboard') {
    return navigateTo({ path: '/login' })
  } else if (token) {
    if (isJwtExpired(token)) {
      auth.logout()
    } else if (to.path === '/login') {
      return navigateTo({ path: '/dashboard' })
    }
  }
})

function isJwtExpired(token: string): boolean {
  try {
    const [, payload] = token.split('.')
    const decoded = JSON.parse(atob(payload || ''))
    return decoded.exp * 1000 < Date.now()
  } catch {
    return true
  }
}
