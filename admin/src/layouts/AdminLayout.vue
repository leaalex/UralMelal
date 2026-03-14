<template>
  <div class="min-h-screen bg-slate-100 flex">
    <!-- Sidebar: сворачивается до иконок, кнопка на границе -->
    <div class="fixed top-4 left-4 bottom-4 z-50 inline-flex">
      <nav
        class="flex-col bg-white/95 backdrop-blur-xl text-slate-700 border border-slate-200 rounded-xl overflow-hidden shadow-xl flex transition-[width] duration-200 ease-out"
        :class="collapsed ? 'w-20' : 'w-72'"
      >
        <div class="shrink-0 pt-6 pb-5 border-b border-slate-200 flex items-center justify-center min-h-[4.5rem] min-w-0" :class="collapsed ? 'px-2' : 'px-4'">
          <router-link to="/" class="flex items-center justify-center min-w-0 shrink-0 w-full">
            <img
              :src="collapsed ? '/admin/logo.png' : '/admin/logo-admin.png'"
              alt="Урал Металл"
              class="h-10 object-contain"
              :class="collapsed ? 'w-10 min-w-10 min-h-10' : 'w-auto max-w-full'"
            />
          </router-link>
        </div>
      <div class="flex-1 flex flex-col gap-2 pt-4 overflow-y-auto min-h-0 min-w-0" :class="collapsed ? 'px-2 items-center' : 'px-2'">
        <router-link v-if="auth.isEditor" to="/" class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" :class="[collapsed ? 'justify-center px-2 gap-0' : 'gap-2 px-3']" exact-active-class="!bg-slate-200 !text-slate-900">
          <HomeIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap overflow-hidden" :class="{ 'opacity-0 w-0': collapsed }">Дашборд</span>
        </router-link>
        <router-link v-if="auth.isEditor" to="/catalog" class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" :class="[collapsed ? 'justify-center px-2 gap-0' : 'gap-2 px-3']" active-class="!bg-slate-200 !text-slate-900">
          <FolderIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap overflow-hidden" :class="{ 'opacity-0 w-0': collapsed }">Каталог</span>
        </router-link>
        <router-link v-if="auth.isEditor" to="/catalog/import" class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" :class="[collapsed ? 'justify-center px-2 gap-0' : 'gap-2 px-3']" active-class="!bg-slate-200 !text-slate-900">
          <ArrowUpTrayIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap overflow-hidden" :class="{ 'opacity-0 w-0': collapsed }">Импорт</span>
        </router-link>
        <router-link v-if="auth.isEditor" to="/categories" class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" :class="[collapsed ? 'justify-center px-2 gap-0' : 'gap-2 px-3']" active-class="!bg-slate-200 !text-slate-900">
          <TagIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap overflow-hidden" :class="{ 'opacity-0 w-0': collapsed }">Категории</span>
        </router-link>
        <router-link v-if="auth.isEditor" to="/content" class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" :class="[collapsed ? 'justify-center px-2 gap-0' : 'gap-2 px-3']" active-class="!bg-slate-200 !text-slate-900">
          <DocumentTextIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap overflow-hidden" :class="{ 'opacity-0 w-0': collapsed }">Контент</span>
        </router-link>
        <router-link v-if="auth.isManager" to="/leads" class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" :class="[collapsed ? 'justify-center px-2 gap-0' : 'gap-2 px-3']" active-class="!bg-slate-200 !text-slate-900">
          <EnvelopeIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap overflow-hidden" :class="{ 'opacity-0 w-0': collapsed }">Заявки</span>
        </router-link>
        <router-link v-if="auth.isAdmin" to="/users" class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors" :class="[collapsed ? 'justify-center px-2 gap-0' : 'gap-2 px-3']" active-class="!bg-slate-200 !text-slate-900">
          <UsersIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap overflow-hidden" :class="{ 'opacity-0 w-0': collapsed }">Пользователи</span>
        </router-link>
      </div>
      <div class="shrink-0 py-4 border-t border-slate-200 flex flex-col gap-2 min-w-0" :class="collapsed ? 'px-2 items-center' : 'px-3'">
        <span class="text-sm text-slate-600 truncate overflow-hidden" :class="{ 'opacity-0 w-0 h-0 overflow-hidden': collapsed }">{{ auth.user?.username }}</span>
        <button
          type="button"
          class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors border border-slate-200"
          :class="[collapsed ? 'justify-center px-2 gap-0' : 'gap-2 px-3']"
          @click="logout"
        >
          <ArrowRightOnRectangleIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap overflow-hidden" :class="{ 'opacity-0 w-0': collapsed }">Выйти</span>
        </button>
      </div>
      </nav>
      <button
        type="button"
        class="absolute top-[2.25rem] -right-3 w-6 h-6 flex items-center justify-center rounded-lg bg-white border border-slate-200 shadow-md text-slate-500 hover:bg-slate-50 hover:text-slate-700 transition-colors z-10"
        :aria-label="collapsed ? 'Развернуть меню' : 'Свернуть меню'"
        @click="collapsed = !collapsed"
      >
        <ChevronLeftIcon v-if="!collapsed" class="h-4 w-4" />
        <ChevronRightIcon v-else class="h-4 w-4" />
      </button>
    </div>
    <main
      class="flex-1 min-h-0 flex flex-col transition-[margin-left] duration-200 ease-out"
      :class="[
        collapsed ? 'ml-24' : 'ml-[19rem]',
        isCategories ? 'py-4 px-6' : 'p-6',
      ]"
    >
      <div class="flex-1 min-h-0 flex flex-col min-w-0">
        <router-view />
      </div>
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
  ChevronLeftIcon,
  ChevronRightIcon,
  ArrowRightOnRectangleIcon,
} from '@heroicons/vue/24/outline'
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const auth = useAuthStore()
const collapsed = ref(false)
const route = useRoute()
const isCategories = computed(() => route.name === 'Categories')

function logout() {
  auth.logout()
  const base = (import.meta.env.BASE_URL || '/').replace(/\/?$/, '/')
  window.location.href = base + 'login'
}
</script>
