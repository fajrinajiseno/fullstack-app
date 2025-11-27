export default defineEventHandler(async (event) => {
  deleteCookie(event, TOKEN_KEY, { path: '/' })
  return { success: true }
})
