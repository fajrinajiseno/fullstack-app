<template>
  <UDashboardGroup>
    <UDashboardSidebar
      collapsible
      resizable
      :ui="{ footer: 'border-t border-default' }"
      open
      toggle-side="left"
    >
      <template #header>
        <UIcon
          name="i-simple-icons-nuxtdotjs"
          class="size-5 text-primary mx-auto"
        />
      </template>

      <template #default="{ collapsed }">
        <UNavigationMenu
          :collapsed="collapsed"
          :items="items[0]"
          orientation="vertical"
        />

        <ClientOnly>
          <UNavigationMenu
            v-if="auth.getUser()"
            :collapsed="collapsed"
            :items="items[1]"
            orientation="vertical"
            class="mt-auto"
          />
        </ClientOnly>
      </template>
    </UDashboardSidebar>
    <UDashboardPanel resizable>
      <template #header>
        <UDashboardNavbar title="Dashboard" />
      </template>
      <template #body>
        <slot />
        <ClientOnly>
          <div class="text-sm text-gray-600">
            Auto logout in
            <strong>
              {{ String(minutes).padStart(2, '0') }}:
              {{ String(seconds).padStart(2, '0') }}
            </strong>
          </div>
        </ClientOnly>
      </template>
    </UDashboardPanel>
  </UDashboardGroup>
</template>

<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'

const auth = useAuthStore()
const { minutes, seconds } = useIdleLogout()

const items: NavigationMenuItem[][] = [
  [
    {
      label: 'Home',
      icon: 'i-lucide-home',
      active: true
    }
  ],
  [
    {
      label: auth.getUser()?.email,
      icon: 'i-lucide-user'
    },
    {
      label: `User ID: ${auth.getUser()?.id}`,
      icon: 'i-lucide-id-card-lanyard'
    },
    {
      label: 'Log out',
      icon: 'i-lucide-log-out',
      class: 'pointer',
      onSelect: () => {
        auth.logout()
      }
    }
  ]
]
</script>
