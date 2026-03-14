<template>
  <div class="min-h-screen flex items-center justify-center bg-slate-200">
    <UiCard class="w-80" padding="lg">
      <h1 class="text-xl font-bold mb-6 text-slate-800">Вход</h1>
      <form @submit.prevent="login" class="space-y-4">
        <UiInput v-model="username" type="text" label="Логин" required />
        <UiInput v-model="password" type="password" label="Пароль" required />
        <p v-if="error" class="text-red-500 text-sm">{{ error }}</p>
        <UiButton type="submit" class="w-full">Войти</UiButton>
      </form>
    </UiCard>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { UiButton, UiCard, UiInput } from '@ui'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const auth = useAuthStore()
const username = ref('')
const password = ref('')
const error = ref('')

async function login() {
  error.value = ''
  try {
    await auth.login(username.value, password.value)
    router.push('/')
  } catch {
    error.value = 'Неверный логин или пароль'
  }
}
</script>
