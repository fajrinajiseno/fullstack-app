<template>
  <div class="flex flex-col items-center justify-center gap-4 p-4">
    <UPageCard class="w-full max-w-md">
      <UAuthForm
        :schema="schema"
        title="Login"
        description="Enter your credentials to access dashboard."
        icon="i-lucide-user"
        :fields="fields"
        :submit="{ label: 'Submit', block: true }"
        @submit="onSubmit"
      >
        <template #description>
          Don't have an account?
          <ULink to="/register" class="text-primary font-medium">Sign up</ULink
          >.
        </template>
      </UAuthForm>
    </UPageCard>
  </div>
</template>

<script setup lang="ts">
import type { ModelError, User } from '../../generated/openapi-client'
import * as z from 'zod'
import type { FormSubmitEvent, AuthFormField } from '@nuxt/ui'

definePageMeta({
  middleware: 'auth'
})

const auth = useAuthStore()
const toast = useToast()

const fields: AuthFormField[] = [
  {
    name: 'email',
    type: 'email',
    label: 'Email',
    placeholder: 'Enter your email',
    required: true
  },
  {
    name: 'password',
    label: 'Password',
    type: 'password',
    placeholder: 'Enter your password',
    required: true
  }
]

const schema = z.object({
  email: z.email('Invalid email'),
  password: z.string('Password is required')
})

type Schema = z.output<typeof schema>

async function onSubmit(payload: FormSubmitEvent<Schema>) {
  // const { api } = useGeneratedClient()
  try {
    // const data = await api.v1AuthLoginPost({
    //   v1AuthLoginPostRequest: {
    //     email: payload.data.email,
    //     password: payload.data.password
    //   }
    // })
    const res = await $fetch<User>('/api/login', {
      method: 'POST',
      body: {
        email: payload.data.email,
        password: payload.data.password
      }
    })
    auth.setUser({
      id: res.id!,
      email: res.email!,
      token: res.token!
    })
    toast.add({ title: 'success login' })
    window.location.assign('/dashboard')
  } catch (error) {
    toast.add({
      title: `${(error as ModelError).code}`,
      description: (error as ModelError).message,
      color: 'error'
    })
  }
}
</script>
