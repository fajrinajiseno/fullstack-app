export interface User {
  id: string
  email: string
  token: string
}

export function useAuthStore() {
  const USER_KEY = 'auth.user'

  function getUser(): User | null {
    if (import.meta.server) return null
    try {
      const rawUser = localStorage.getItem(USER_KEY)
      if (rawUser) return JSON.parse(rawUser) as User
    } catch (e) {
      console.warn('auth: getUser failed', e)
    }
    return null
  }

  function setUser(user: User | null) {
    if (import.meta.server) return
    try {
      if (user) localStorage.setItem(USER_KEY, JSON.stringify(user))
      else localStorage.removeItem(USER_KEY)
    } catch (e) {
      console.warn('auth: setUser failed', e)
    }
  }

  async function logout() {
    setUser(null)
    await $fetch<User>('/api/logout', { method: 'POST' })
    window.location.assign('/login')
  }

  return {
    logout,
    getUser,
    setUser
  }
}
