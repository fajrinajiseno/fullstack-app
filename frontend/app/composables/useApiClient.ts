import type { User } from '../../generated/openapi-client'

export function useApiClient() {
  const apiLogin = (email: string, password: string) => {
    return new Promise<User>((resolve, reject) => {
      $fetch<User>('/api/login', {
        method: 'POST',
        body: {
          email,
          password
        }
      })
        .then((response) => resolve(response))
        .catch((error) => reject(error))
    })
  }

  return { apiLogin }
}
