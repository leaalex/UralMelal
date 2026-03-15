<template>
  <div class="min-h-screen flex flex-col bg-slate-50 text-slate-800 overflow-x-hidden">
    <!-- Mobile: logo in top-left over video -->
    <header class="lg:hidden fixed top-4 left-4 z-[60] drop-shadow-lg">
      <router-link to="/" class="block">
        <Logo size="lg" />
      </router-link>
    </header>

    <!-- Desktop: две отдельные панели (меню + контакты) одна под другой -->
    <div class="hidden lg:flex fixed top-4 left-4 bottom-4 z-50 flex-col gap-2">
      <!-- Панель 1: Меню -->
      <nav
        class="flex flex-col bg-white/50 backdrop-blur-2xl backdrop-saturate-150 text-slate-700 border border-white/20 rounded-xl overflow-hidden shadow-2xl min-w-0 flex-1 min-h-0"
        :class="isSidebarOpen ? 'w-[15.4rem] sidebar-expanded' : 'w-16'"
      >
        <div class="shrink-0 pt-6 pb-5 border-b border-white/20 flex items-center justify-center min-h-[4.5rem] min-w-0 px-2">
          <div class="sidebar-logo-compact flex justify-center items-center w-12 min-w-12 h-10">
            <Logo compact size="sm" class="sidebar-logo-img" />
          </div>
          <div class="sidebar-logo-full w-full flex justify-start items-center h-10">
            <Logo size="sm" class="sidebar-logo-img" />
          </div>
        </div>
        <div class="flex-1 flex flex-col gap-2 pt-4 overflow-y-auto min-h-0 min-w-0 px-2" :class="isSidebarOpen ? '' : 'items-center'">
          <router-link
            to="/"
            exact-active-class="!bg-slate-500 !text-white"
            class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-400 hover:text-slate-900 nav-link"
          >
            <HomeIcon class="h-5 w-5 shrink-0" />
            <span class="text-sm whitespace-nowrap overflow-hidden sidebar-link-text">Главная</span>
          </router-link>
          <router-link
            to="/catalog"
            active-class="!bg-slate-500 !text-white"
            class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-400 hover:text-slate-900 nav-link"
          >
            <CubeIcon class="h-5 w-5 shrink-0" />
            <span class="text-sm whitespace-nowrap overflow-hidden sidebar-link-text">Каталог</span>
          </router-link>
          <router-link
            to="/contacts"
            exact-active-class="!bg-slate-500 !text-white"
            class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-400 hover:text-slate-900 nav-link"
          >
            <EnvelopeIcon class="h-5 w-5 shrink-0" />
            <span class="text-sm whitespace-nowrap overflow-hidden sidebar-link-text">О компании</span>
          </router-link>
        </div>
      </nav>
      <!-- Панель 2: Контакты -->
      <div
        class="flex flex-col bg-white/50 backdrop-blur-2xl backdrop-saturate-150 text-slate-700 border border-white/20 rounded-xl overflow-hidden shadow-2xl min-w-0 shrink-0"
        :class="isSidebarOpen ? 'w-[15.4rem] sidebar-expanded' : 'w-16'"
      >
        <div class="flex flex-col gap-2 p-2" :class="sidebarExpanded ? '' : 'items-center'">
          <a :href="`tel:${telHref}`" class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-400 hover:text-slate-900 nav-link no-underline" aria-label="Позвонить">
            <PhoneIcon class="h-5 w-5 shrink-0" />
            <span class="text-sm whitespace-nowrap overflow-hidden sidebar-link-text">{{ contactPhone }}</span>
            <span v-if="sidebarExpanded" role="button" tabindex="0" class="shrink-0 w-8 h-8 flex items-center justify-center rounded-lg hover:bg-slate-200 cursor-pointer" aria-label="Копировать" @click.stop.prevent="copyToClipboard(contactPhone)">
              <DocumentDuplicateIcon class="h-4 w-4" />
            </span>
          </a>
          <a :href="`mailto:${contactEmail}`" class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-400 hover:text-slate-900 nav-link no-underline" aria-label="Написать">
            <EnvelopeIcon class="h-5 w-5 shrink-0" />
            <span class="text-sm whitespace-nowrap overflow-hidden sidebar-link-text">{{ contactEmail }}</span>
            <span v-if="sidebarExpanded" role="button" tabindex="0" class="shrink-0 w-8 h-8 flex items-center justify-center rounded-lg hover:bg-slate-200 cursor-pointer" aria-label="Копировать" @click.stop.prevent="copyToClipboard(contactEmail)">
              <DocumentDuplicateIcon class="h-4 w-4" />
            </span>
          </a>
        </div>
      </div>
    <button
      type="button"
      class="absolute top-[1.6rem] -right-2 w-4 h-9 flex items-center justify-center rounded-full bg-white/50 backdrop-blur-2xl backdrop-saturate-150 text-slate-600 hover:text-slate-800 z-10 shadow-sm border border-white/20"
      :aria-label="isSidebarOpen ? 'Свернуть меню' : 'Развернуть меню'"
      @click="toggleSidebar"
    >
      <ChevronLeftIcon v-if="isSidebarOpen" class="h-3 w-3" />
      <ChevronRightIcon v-else class="h-3 w-3" />
    </button>
  </div>

    <!-- Mobile: bottom bar - two panels -->
    <div class="lg:hidden fixed bottom-4 left-4 right-4 flex gap-2 z-50">
      <!-- Панель 1: Главная, Каталог, О компании -->
      <nav class="flex-1 flex bg-white/50 backdrop-blur-2xl backdrop-saturate-150 text-slate-700 border border-white/20 rounded-xl overflow-hidden shadow-2xl">
      <router-link
        to="/"
        exact-active-class="!bg-slate-500 !text-white"
        class="flex flex-col items-center justify-center flex-1 py-3 text-slate-600 hover:bg-slate-400 hover:text-slate-900"
      >
          <HomeIcon class="h-6 w-6" />
        <span class="text-xs mt-0.5">Главная</span>
      </router-link>
      <router-link
        to="/catalog"
        active-class="!bg-slate-500 !text-white"
        class="flex flex-col items-center justify-center flex-1 py-3 text-slate-600 hover:bg-slate-400 hover:text-slate-900"
      >
          <CubeIcon class="h-6 w-6" />
        <span class="text-xs mt-0.5">Каталог</span>
      </router-link>
      <router-link
        to="/contacts"
        exact-active-class="!bg-slate-500 !text-white"
        class="flex flex-col items-center justify-center flex-1 py-3 text-slate-600 hover:bg-slate-400 hover:text-slate-900"
      >
          <EnvelopeIcon class="h-6 w-6" />
        <span class="text-xs mt-0.5">О компании</span>
      </router-link>
      </nav>
      <!-- Панель 2: Связаться -->
      <button
        type="button"
        class="flex flex-col items-center justify-center px-5 py-3 bg-white/50 backdrop-blur-2xl backdrop-saturate-150 text-slate-600 border border-white/20 rounded-xl shadow-2xl hover:bg-white/60 hover:text-slate-900"
        @click="contactModalOpen = true"
      >
        <ChatBubbleLeftRightIcon class="h-6 w-6" />
        <span class="text-xs mt-0.5">Связаться</span>
      </button>
    </div>

    <main :class="['flex-1 pb-36 lg:pt-0 lg:pb-0', isHome ? 'lg:ml-0 pt-0' : (sidebarExpanded ? 'lg:ml-[17rem] pt-20' : 'lg:ml-20 pt-20')]">
      <router-view :key="$route.fullPath" />
    </main>
    <footer class="bg-white border-t border-slate-200 py-6 mt-auto">
      <div class="container mx-auto px-4 text-center">
        <p class="text-sm text-slate-600">© {{ currentYear }} Урал Металл</p>
      </div>
    </footer>

    <ContactModal v-model="contactModalOpen" />

    <!-- Search button + panel - top right over video (only on home) -->
    <div v-if="route.path === '/'" class="fixed top-4 right-4 z-[60] flex justify-end gap-2">
      <form class="group flex flex-row-reverse items-center gap-2" @submit.prevent="doSearch">
        <button
          type="button"
          class="shrink-0 flex items-center justify-center w-12 h-12 rounded-xl bg-white/50 backdrop-blur-2xl backdrop-saturate-150 text-slate-700 hover:bg-white/60 hover:text-slate-900 border border-white/20 shadow-2xl"
          aria-label="Поиск"
          @click.prevent="toggleSearch"
        >
          <MagnifyingGlassIcon class="h-5 w-5" />
        </button>
        <div
          class="overflow-hidden rounded-xl bg-white/50 backdrop-blur-2xl backdrop-saturate-150 border border-white/20 shadow-2xl w-0 opacity-0 group-hover:w-[16.8rem] group-hover:opacity-100"
          :class="{ '!w-[16.8rem] !opacity-100': searchExpanded }"
        >
          <input
            ref="searchInput"
            v-model="searchQuery"
            type="search"
            placeholder="Поиск по названию, артикулу..."
            autocomplete="off"
            class="w-[16.8rem] px-3 py-2.5 text-slate-800 placeholder-slate-500 bg-transparent focus:outline-none border-0 rounded-xl"
            @keydown.esc="searchExpanded = false"
            @keydown.enter.prevent="doSearch"
          />
        </div>
      </form>
    </div>

    <!-- Floating "Связаться с нами" button - desktop only (mobile: in bottom menu) -->
    <button
      type="button"
      class="hidden lg:flex fixed z-40 items-center justify-center w-12 h-12 rounded-xl border border-white/20 bg-white/50 backdrop-blur-2xl backdrop-saturate-150 text-slate-700 shadow-2xl hover:bg-white/60 hover:text-slate-900 right-6 bottom-6"
      aria-label="Связаться с нами"
      @click="contactModalOpen = true"
    >
      <ChatBubbleLeftRightIcon class="h-5 w-5" />
    </button>
  </div>
</template>

<script setup>
import { ref, computed, provide, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  HomeIcon,
  CubeIcon,
  EnvelopeIcon,
  PhoneIcon,
  DocumentDuplicateIcon,
  MagnifyingGlassIcon,
  ChatBubbleLeftRightIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
} from '@heroicons/vue/24/outline'
import ContactModal from './components/ContactModal.vue'
import Logo from './components/Logo.vue'

const route = useRoute()
const router = useRouter()
const isHome = computed(() => route.path === '/')
const isSidebarOpen = ref(route.path === '/')
const sidebarExpanded = computed(() => isSidebarOpen.value)

function toggleSidebar() {
  isSidebarOpen.value = !isSidebarOpen.value
}
const contactModalOpen = ref(false)
const searchQuery = ref('')
const searchExpanded = ref(false)
const searchInput = ref(null)

function toggleSearch() {
  searchExpanded.value = !searchExpanded.value
  if (searchExpanded.value) {
    nextTick(() => searchInput.value?.focus())
  }
}

function doSearch() {
  const q = searchQuery.value?.trim()
  if (q) {
    router.push({ path: '/catalog', query: { q } })
    searchExpanded.value = false
    searchQuery.value = ''
  }
}
const currentYear = new Date().getFullYear()
const contactPhone = ref(import.meta.env.VITE_CONTACT_PHONE || '+7 (XXX) XXX-XX-XX')
const contactEmail = ref(import.meta.env.VITE_CONTACT_EMAIL || 'info@uralmetall.ru')
const telHref = computed(() => contactPhone.value.replace(/\D/g, ''))

provide('openContactModal', () => { contactModalOpen.value = true })
provide('contactPhone', contactPhone)
provide('contactEmail', contactEmail)
provide('telHref', telHref)

async function copyToClipboard(text) {
  try {
    await navigator.clipboard.writeText(text)
  } catch {
    // fallback for older browsers
    const ta = document.createElement('textarea')
    ta.value = text
    ta.style.position = 'fixed'
    ta.style.opacity = '0'
    document.body.appendChild(ta)
    ta.select()
    document.execCommand('copy')
    document.body.removeChild(ta)
  }
}
</script>

<style scoped>
.sidebar-logo-img img {
  height: 2.5rem;
  width: auto;
  max-height: 2.5rem;
}
.sidebar-logo-full {
  display: none;
}
nav.sidebar-expanded .sidebar-logo-compact {
  display: none;
}
nav.sidebar-expanded .sidebar-logo-full {
  display: flex;
}
/* Collapsed: center icon only */
.nav-link {
  justify-content: center;
  padding-left: 0.5rem;
  padding-right: 0.5rem;
  gap: 0;
}
.sidebar-link-text {
  opacity: 0;
  width: 0;
  overflow: hidden;
}
/* Expanded */
nav.sidebar-expanded .nav-link,
.sidebar-expanded .nav-link {
  justify-content: flex-start;
  gap: 0.5rem;
  padding-left: 0.75rem;
  padding-right: 0.75rem;
}
nav.sidebar-expanded .sidebar-link-text,
.sidebar-expanded .sidebar-link-text {
  opacity: 1;
  width: auto;
}
</style>
