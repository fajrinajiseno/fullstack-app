import { Configuration, DefaultApi } from '../../generated/openapi-client'
import { useRuntimeConfig } from '#imports'

export function useGeneratedClientServer() {
  const config = useRuntimeConfig()
  const apiBase = config.public.apiBase || 'http://localhost:8080'
  const cfg = new Configuration({
    basePath: apiBase
  })
  const api = new DefaultApi(cfg)

  return { api }
}
