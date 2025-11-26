<template>
  <div>
    <ClientOnly>
      <h1 class="text-lg mb-4">
        Hello {{ auth.getUser()?.email }}, welcome back
      </h1>
    </ClientOnly>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'dashboard',
  middleware: 'auth'
})

const auth = useAuthStore()
</script>
