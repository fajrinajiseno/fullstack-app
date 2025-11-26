<template>
  <div class="flex flex-col items-center justify-center gap-4 p-4">
    <UPageCard class="w-full max-w-md">
      <UAuthForm
        :schema="schema"
        title="Login"
        description="Enter your Email and Password to create account."
        icon="i-lucide-user"
        :fields="fields"
        :submit="{ label: 'Submit', block: true }"
        @submit="onSubmit"
      >
        <template #description>
          Have an account?
          <ULink to="/login" class="text-primary font-medium">Login</ULink>.
        </template>
      </UAuthForm>
    </UPageCard>
  </div>
</template>

<script setup lang="ts">
import type { ModelError } from '../../generated/openapi-client'
import * as z from 'zod'
import type { FormSubmitEvent, AuthFormField } from '@nuxt/ui'

definePageMeta({
  middleware: 'auth'
})

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
  },
  {
    name: 'confirmPassword',
    label: 'Confirm Password',
    type: 'password',
    placeholder: 'Enter password confirmation',
    required: true
  }
]

const schema = z
  .object({
    email: z.email('Invalid email'),
    password: z.string('Password is required'),
    confirmPassword: z.string('Confirmation Password is required')
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords don't match",
    path: ['confirmPassword'] // path of error
  })

type Schema = z.output<typeof schema>

async function onSubmit(payload: FormSubmitEvent<Schema>) {
  const { api } = useGeneratedClient()
  try {
    const data = await api.v1AuthRegisterPost({
      v1AuthRegisterPostRequest: {
        email: payload.data.email,
        password: payload.data.password,
        confirmPassword: payload.data.confirmPassword
      }
    })
    toast.add({
      title: data.message,
      description: 'Please login using your new account'
    })
    navigateTo({
      path: '/login'
    })
  } catch (error) {
    toast.add({
      title: `${(error as ModelError).code}`,
      description: (error as ModelError).message,
      color: 'error'
    })
  }
}
</script>
