type JwtPayload = {
  exp: number
}

export function getJwtExpiresIn(token: string): number {
  try {
    const [, payload] = token.split('.')
    const decoded = JSON.parse(
      Buffer.from(payload, 'base64').toString('utf-8')
    ) as JwtPayload

    const expiresAtMs = decoded.exp * 1000
    const nowMs = Date.now()

    const secondsLeft = Math.floor((expiresAtMs - nowMs) / 1000)

    // Prevent negative or zero maxAge
    return Math.max(secondsLeft, 0)
  } catch {
    return 0
  }
}
