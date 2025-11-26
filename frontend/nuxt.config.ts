// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: [
    '@nuxt/eslint',
    '@nuxt/hints',
    '@nuxt/image',
    '@nuxt/scripts',
    '@nuxt/test-utils/module',
    '@nuxt/ui',
    '@pinia/nuxt'
  ],
  devtools: { enabled: true },
  css: ['~/assets/css/main.css'],
  colorMode: {
    preference: 'light'
  },
  runtimeConfig: {
    public: {
      inactivityLimit: Number(process.env.PUBLIC_INACTIVITY_LIMIT || 900000),
      apiBase: ''
    }
  },
  compatibilityDate: '2025-07-15',
  typescript: {
    // Add globals from Vitest (describe, it, vi, etc.)
    // https://vitest.dev/config/#globals
    tsConfig: {
      compilerOptions: {
        types: ['vitest/globals']
      }
    }
  },
  eslint: {
    config: {
      stylistic: true
    }
  }
})
