import { Configuration, DefaultApi } from '../../generated/openapi-client'
import { handleOpenapiError } from '#shared/utils/handleOpenapiError'
import { ERROR_UNAUTHORIZED } from '#shared/types/api'
import { TOKEN_KEY } from '#shared/types/auth'
import type { H3Event } from 'h3'
import { deleteCookie } from 'h3'
import { useRuntimeConfig } from '#imports'

export function useGeneratedClientServer(event: H3Event) {
  const config = useRuntimeConfig()
  const apiBase = config.backendApiBase
  const cfg = new Configuration({
    basePath: apiBase,
    middleware: [
      {
        async post(context) {
          if (!context.response.ok) {
            const errorParsed = await handleOpenapiError(context)
            if (errorParsed.message.includes(ERROR_UNAUTHORIZED)) {
              deleteCookie(event, TOKEN_KEY, { path: '/' })
            }
            throw errorParsed
          }
          return Promise.resolve(context.response)
        }
      }
    ]
  })
  const api = new DefaultApi(cfg)

  return { api }
}
