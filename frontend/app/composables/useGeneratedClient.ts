import { Configuration, DefaultApi } from '../../generated/openapi-client'
import { useRuntimeConfig } from '#imports'

export function useGeneratedClient() {
  const ERROR_UNAUTHORIZED =
    'security requirements failed: authorization failed'
  const config = useRuntimeConfig()
  const auth = useAuthStore()
  const apiBase = config.public.apiBase || 'http://localhost:8080'
  const cfg = new Configuration({
    basePath: apiBase,
    headers: {
      ...(auth.getUser()?.token
        ? { Authorization: `Bearer ${auth.getUser()?.token}` }
        : {})
    },
    middleware: [
      {
        async post(context) {
          if (!context.response.ok) {
            const errorParsed = await usehandleError(context)
            if (errorParsed.message.includes(ERROR_UNAUTHORIZED)) {
              await auth.logout()
            } else {
              throw errorParsed
            }
          }
          return Promise.resolve(context.response)
        }
      }
    ]
  })
  const api = new DefaultApi(cfg)

  return { api }
}
