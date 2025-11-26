export const useIdleLogout = () => {
  const config = useRuntimeConfig()
  const { getUser, logout } = useAuthStore()

  const DEFAULT_INACTIVITY_LIMIT = 15 * 60 * 1000 // 15 minutes in ms
  const INACTIVITY_LIMIT_MS =
    config.public.inactivityLimit || DEFAULT_INACTIVITY_LIMIT

  const remainingMs = useState<number>(
    'idle-remaining-ms',
    () => INACTIVITY_LIMIT_MS
  )

  let timeoutId: ReturnType<typeof setTimeout> | null = null
  let intervalId: ReturnType<typeof setInterval> | null = null

  const clearTimers = () => {
    if (timeoutId) {
      clearTimeout(timeoutId)
      timeoutId = null
    }
    if (intervalId) {
      clearInterval(intervalId)
      intervalId = null
    }
  }

  const logoutForInactivity = async () => {
    await logout()
  }

  const startCountdownInterval = (startAt: number) => {
    if (intervalId) clearInterval(intervalId)

    intervalId = setInterval(() => {
      const elapsed = Date.now() - startAt
      const left = INACTIVITY_LIMIT_MS - elapsed
      remainingMs.value = left > 0 ? left : 0

      if (left <= 0 && intervalId) {
        clearInterval(intervalId)
        intervalId = null
      }
    }, 1000)
  }

  const resetTimer = () => {
    if (!getUser()?.token) {
      clearTimers()
      remainingMs.value = INACTIVITY_LIMIT_MS
      return
    }

    clearTimers()

    const startedAt = Date.now()
    remainingMs.value = INACTIVITY_LIMIT_MS

    timeoutId = setTimeout(logoutForInactivity, INACTIVITY_LIMIT_MS)
    startCountdownInterval(startedAt)
  }

  if (import.meta.client) {
    const events = ['mousemove', 'keydown', 'click', 'scroll', 'touchstart']

    const onActivity = () => {
      resetTimer()
    }

    onMounted(() => {
      // attach listeners only when this layout is in use
      events.forEach((evt) => window.addEventListener(evt, onActivity))
      resetTimer()
    })

    onUnmounted(() => {
      events.forEach((evt) => window.removeEventListener(evt, onActivity))
      clearTimers()
    })
  }

  const secondsLeft = computed(() => Math.floor(remainingMs.value / 1000))
  const minutes = computed(() => Math.floor(secondsLeft.value / 60))
  const seconds = computed(() => secondsLeft.value % 60)

  return {
    remainingMs,
    secondsLeft,
    minutes,
    seconds
  }
}
