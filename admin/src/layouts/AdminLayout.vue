<template>
  <div class="min-h-screen bg-slate-100 flex">
    <!-- Sidebar: как в клиенте, но всегда развёрнуто -->
    <nav class="fixed top-4 left-4 bottom-4 w-72 flex-col bg-white/95 backdrop-blur-xl text-slate-700 border border-slate-200 z-50 rounded-xl overflow-hidden shadow-xl flex">
      <div class="shrink-0 pt-6 px-4 pb-5 border-b border-slate-200 flex items-center min-h-[4.5rem]">
        <router-link to="/" class="flex items-center">
          <img
            src="/admin/logo-admin.png"
            alt="Урал Металл"
            class="h-10 w-auto object-contain"
          />
        </router-link>
      </div>
      <div class="flex-1 flex flex-col gap-2 pt-4 px-2 overflow-y-auto min-h-0">
        <router-link v-if="auth.isEditor" to="/" class="flex items-center gap-2 h-12 w-full min-w-12 px-3 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" exact-active-class="!bg-slate-200 !text-slate-900">
          <HomeIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap">Дашборд</span>
        </router-link>
        <router-link v-if="auth.isEditor" to="/catalog" class="flex items-center gap-2 h-12 w-full min-w-12 px-3 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" active-class="!bg-slate-200 !text-slate-900">
          <FolderIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap">Каталог</span>
        </router-link>
        <router-link v-if="auth.isEditor" to="/catalog/import" class="flex items-center gap-2 h-12 w-full min-w-12 px-3 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" active-class="!bg-slate-200 !text-slate-900">
          <ArrowUpTrayIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap">Импорт</span>
        </router-link>
        <router-link v-if="auth.isEditor" to="/categories" class="flex items-center gap-2 h-12 w-full min-w-12 px-3 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" active-class="!bg-slate-200 !text-slate-900">
          <TagIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap">Категории</span>
        </router-link>
        <router-link v-if="auth.isEditor" to="/content" class="flex items-center gap-2 h-12 w-full min-w-12 px-3 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" active-class="!bg-slate-200 !text-slate-900">
          <DocumentTextIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap">Контент</span>
        </router-link>
        <router-link v-if="auth.isManager" to="/leads" class="flex items-center gap-2 h-12 w-full min-w-12 px-3 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" active-class="!bg-slate-200 !text-slate-900">
          <EnvelopeIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap">Заявки</span>
        </router-link>
        <router-link v-if="auth.isAdmin" to="/users" class="flex items-center gap-2 h-12 w-full min-w-12 px-3 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" active-class="!bg-slate-200 !text-slate-900">
          <UsersIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap">Пользователи</span>
        </router-link>
      </div>
      <div class="shrink-0 py-4 px-3 border-t border-slate-200 flex flex-col gap-2">
        <span class="text-sm text-slate-600 truncate">{{ auth.user?.username }}</span>
        <UiButton variant="secondary" size="sm" class="w-full" @click="logout">Выйти</UiButton>
      </div>
    </nav>
    <main class="flex-1 ml-[19rem] p-6">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import {
  HomeIcon,
  FolderIcon,
  ArrowUpTrayIcon,
  TagIcon,
  DocumentTextIcon,
  EnvelopeIcon,
  UsersIcon,
} from '@heroicons/vue/24/outline'
import { UiButton } from '@ui'
import { useAuthStore } from '../stores/auth'

const auth = useAuthStore()

function logout() {
  auth.logout()
  const base = (import.meta.env.BASE_URL || '/').replace(/\/?$/, '/')
  window.location.href = base + 'login'
}
</script>
