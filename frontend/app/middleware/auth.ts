export default defineNuxtRouteMiddleware(async (to) => {
  // if (import.meta.server) {
  //   return
  // }

  const auth = useAuthStore()

  // if (!auth.getUser()?.token && to.path === '/dashboard') {
  //   window.location.assign('/login')
  // } else if (auth.getUser()?.token && to.path === '/login') {
  //   window.location.assign('/dashboard')
  // }
  const token = useCookie<string | null>('auth_token').value

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
