import type {
  V1AuthLoginPostRequest,
  ModelError
} from '../../generated/openapi-client'
import { useGeneratedClientServer } from '../utils/openapicliet'
import { getJwtExpiresIn } from '../utils/getJwtExpiresIn'

export default defineEventHandler(async (event) => {
  const body = await readBody<V1AuthLoginPostRequest>(event)
  const { api } = useGeneratedClientServer(event)

  try {
    const res = await api.v1AuthLoginPost({
      v1AuthLoginPostRequest: body
    })

    const token = res.token
    const maxAge = getJwtExpiresIn(token!)

    if (maxAge <= 0) {
      throw createError({
        statusCode: 401,
        statusMessage: 'Invalid token'
      })
    }

    // Set token in httpOnly cookie (same as before)
    setCookie(event, 'auth_token', token!, {
      httpOnly: process.env.NODE_ENV === 'production',
      secure: process.env.NODE_ENV === 'production',
      sameSite: 'lax',
      path: '/',
      maxAge
    })
    return res
  } catch (error) {
    return createError({
      statusCode: (error as ModelError).code,
      statusMessage: (error as ModelError).message
    })
  }
})
