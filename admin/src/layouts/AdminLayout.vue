<template>
  <div class="min-h-screen bg-slate-100 flex">
    <aside class="w-56 bg-slate-800 text-white fixed h-full">
      <div class="p-4 border-b border-slate-700">
        <span class="font-semibold">Урал Металл</span>
      </div>
      <nav class="p-2">
        <router-link v-if="auth.isEditor" to="/" class="flex items-center gap-2 px-3 py-2.5 rounded-lg transition-all duration-150 hover:bg-slate-700" active-class="bg-slate-700">
          <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" /></svg>
          Дашборд
        </router-link>
        <router-link v-if="auth.isEditor" to="/catalog" class="flex items-center gap-2 px-3 py-2.5 rounded-lg transition-all duration-150 hover:bg-slate-700" active-class="bg-slate-700">
          <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" /></svg>
          Каталог
        </router-link>
        <router-link v-if="auth.isEditor" to="/catalog/import" class="flex items-center gap-2 px-3 py-2.5 rounded-lg transition-all duration-150 hover:bg-slate-700" active-class="bg-slate-700">
          <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12" /></svg>
          Импорт
        </router-link>
        <router-link v-if="auth.isEditor" to="/categories" class="flex items-center gap-2 px-3 py-2.5 rounded-lg transition-all duration-150 hover:bg-slate-700" active-class="bg-slate-700">
          <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" /></svg>
          Категории
        </router-link>
        <router-link v-if="auth.isEditor" to="/content" class="flex items-center gap-2 px-3 py-2.5 rounded-lg transition-all duration-150 hover:bg-slate-700" active-class="bg-slate-700">
          <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" /></svg>
          Контент
        </router-link>
        <router-link v-if="auth.isManager" to="/leads" class="flex items-center gap-2 px-3 py-2.5 rounded-lg transition-all duration-150 hover:bg-slate-700" active-class="bg-slate-700">
          <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" /></svg>
          Заявки
        </router-link>
        <router-link v-if="auth.isAdmin" to="/users" class="flex items-center gap-2 px-3 py-2.5 rounded-lg transition-all duration-150 hover:bg-slate-700" active-class="bg-slate-700">
          <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" /></svg>
          Пользователи
        </router-link>
      </nav>
    </aside>
    <main class="flex-1 ml-56 p-6">
      <header class="flex justify-end items-center mb-6">
        <span class="text-slate-600 mr-2">{{ auth.user?.username }}</span>
        <UiButton variant="secondary" size="sm" @click="logout">Выйти</UiButton>
      </header>
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { UiButton } from '@ui'
import { useAuthStore } from '../stores/auth'

const auth = useAuthStore()

function logout() {
  auth.logout()
  const base = (import.meta.env.BASE_URL || '/').replace(/\/?$/, '/')
  window.location.href = base + 'login'
}
</script>
