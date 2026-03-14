<template>
  <div class="min-h-screen flex flex-col bg-slate-50 text-slate-800 overflow-x-hidden">
    <!-- Mobile: logo in top-left over video -->
    <header class="lg:hidden fixed top-4 left-4 z-[60] drop-shadow-lg">
      <router-link to="/" class="block">
        <Logo size="lg" />
      </router-link>
    </header>

    <!-- Desktop: vertical sidebar (admin style), expanded on home, collapsed + hover on other pages -->
    <nav
      class="hidden lg:flex group fixed top-4 left-4 bottom-4 z-50 flex-col bg-white/50 backdrop-blur-2xl backdrop-saturate-150 text-slate-700 border border-white/20 rounded-xl overflow-hidden shadow-2xl transition-[width] duration-200 ease-out min-w-0"
      :class="isHome ? 'w-[15.4rem] sidebar-expanded' : 'w-16 hover:w-[15.4rem]'"
      @mouseenter="sidebarHovered = true"
      @mouseleave="sidebarHovered = false"
    >
      <div class="shrink-0 pt-6 pb-5 border-b border-white/20 flex items-center justify-center min-h-[4.5rem] min-w-0 px-2">
        <div class="sidebar-logo-compact flex justify-center items-center w-12 min-w-12 h-10">
          <Logo compact size="sm" class="sidebar-logo-img" />
        </div>
        <div class="sidebar-logo-full w-full flex justify-start items-center h-10">
          <Logo size="sm" class="sidebar-logo-img" />
        </div>
      </div>
      <div class="flex-1 flex flex-col gap-2 pt-4 overflow-y-auto min-h-0 min-w-0" :class="isHome ? 'px-2' : 'px-2 items-center'">
        <router-link
          to="/"
          exact-active-class="!bg-slate-500 !text-white"
          class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-400 hover:text-slate-900 transition-colors duration-200 nav-link"
        >
          <HomeIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap overflow-hidden sidebar-link-text">Главная</span>
        </router-link>
        <router-link
          to="/catalog"
          active-class="!bg-slate-500 !text-white"
          class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-400 hover:text-slate-900 transition-colors duration-200 nav-link"
        >
          <CubeIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap overflow-hidden sidebar-link-text">Каталог</span>
        </router-link>
        <router-link
          to="/contacts"
          exact-active-class="!bg-slate-500 !text-white"
          class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-400 hover:text-slate-900 transition-colors duration-200 nav-link"
        >
          <EnvelopeIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap overflow-hidden sidebar-link-text">О компании</span>
        </router-link>
      </div>
      <div class="flex flex-col gap-2 py-4 shrink-0 min-w-0 border-t border-white/20" :class="isHome ? 'px-2' : 'px-2 items-center'">
        <a :href="`tel:${telHref}`" class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-400 hover:text-slate-900 transition-colors duration-200 nav-link no-underline" aria-label="Позвонить">
          <PhoneIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap overflow-hidden sidebar-link-text">{{ contactPhone }}</span>
          <span v-if="sidebarExpanded" role="button" tabindex="0" class="shrink-0 w-8 h-8 flex items-center justify-center rounded-lg hover:bg-slate-200 cursor-pointer" aria-label="Копировать" @click.stop.prevent="copyToClipboard(contactPhone)">
            <DocumentDuplicateIcon class="h-4 w-4" />
          </span>
        </a>
        <a :href="`mailto:${contactEmail}`" class="flex items-center h-12 w-full min-w-12 rounded-lg text-slate-600 hover:bg-slate-400 hover:text-slate-900 transition-colors duration-200 nav-link no-underline" aria-label="Написать">
          <EnvelopeIcon class="h-5 w-5 shrink-0" />
          <span class="text-sm whitespace-nowrap overflow-hidden sidebar-link-text">{{ contactEmail }}</span>
          <span v-if="sidebarExpanded" role="button" tabindex="0" class="shrink-0 w-8 h-8 flex items-center justify-center rounded-lg hover:bg-slate-200 cursor-pointer" aria-label="Копировать" @click.stop.prevent="copyToClipboard(contactEmail)">
            <DocumentDuplicateIcon class="h-4 w-4" />
          </span>
        </a>
      </div>
    </nav>

    <!-- Mobile: bottom bar - two panels -->
    <div class="lg:hidden fixed bottom-4 left-4 right-4 flex gap-2 z-50">
      <!-- Панель 1: Главная, Каталог, О компании -->
      <nav class="flex-1 flex bg-white/50 backdrop-blur-2xl backdrop-saturate-150 text-slate-700 border border-white/20 rounded-xl overflow-hidden shadow-2xl">
      <router-link
        to="/"
        exact-active-class="!bg-slate-500 !text-white"
        class="flex flex-col items-center justify-center flex-1 py-3 text-slate-600 hover:bg-slate-400 hover:text-slate-900 transition-colors"
      >
          <HomeIcon class="h-6 w-6" />
        <span class="text-xs mt-0.5">Главная</span>
      </router-link>
      <router-link
        to="/catalog"
        active-class="!bg-slate-500 !text-white"
        class="flex flex-col items-center justify-center flex-1 py-3 text-slate-600 hover:bg-slate-400 hover:text-slate-900 transition-colors"
      >
          <CubeIcon class="h-6 w-6" />
        <span class="text-xs mt-0.5">Каталог</span>
      </router-link>
      <router-link
        to="/contacts"
        exact-active-class="!bg-slate-500 !text-white"
        class="flex flex-col items-center justify-center flex-1 py-3 text-slate-600 hover:bg-slate-400 hover:text-slate-900 transition-colors"
      >
          <EnvelopeIcon class="h-6 w-6" />
        <span class="text-xs mt-0.5">О компании</span>
      </router-link>
      </nav>
      <!-- Панель 2: Связаться -->
      <button
        type="button"
        class="flex flex-col items-center justify-center px-5 py-3 bg-white/50 backdrop-blur-2xl backdrop-saturate-150 text-slate-600 border border-white/20 rounded-xl shadow-2xl hover:bg-slate-100 hover:text-slate-900 transition-colors"
        @click="contactModalOpen = true"
      >
        <ChatBubbleLeftRightIcon class="h-6 w-6" />
        <span class="text-xs mt-0.5">Связаться</span>
      </button>
    </div>

    <main :class="['flex-1 pb-36 lg:pt-0 lg:pb-0 transition-[margin-left] duration-200 ease-out', isHome ? 'lg:ml-0 pt-0' : (sidebarExpanded ? 'lg:ml-[17rem] pt-20' : 'lg:ml-20 pt-20')]">
      <router-view :key="$route.fullPath" />
    </main>
    <footer class="bg-white border-t border-slate-200 py-6 mt-auto">
      <div class="container mx-auto px-4 text-center">
        <p class="text-sm text-slate-600">© {{ currentYear }} Урал Металл</p>
      </div>
    </footer>

    <ContactModal v-model="contactModalOpen" />

    <!-- Search button + panel - top right over video (only on home) -->
    <div v-if="route.path === '/'" class="fixed top-4 right-4 z-[60] flex justify-end">
      <form
        class="group flex flex-row-reverse items-center rounded-xl border border-slate-200 bg-white/95 backdrop-blur-sm shadow-lg overflow-hidden hover:ring-2 hover:ring-slate-200 transition-shadow"
        @submit.prevent="doSearch"
      >
        <button
          type="button"
          class="shrink-0 flex items-center justify-center w-12 h-12 text-slate-700 hover:bg-slate-50 hover:text-slate-900 transition-colors"
          aria-label="Поиск"
          @click.prevent="toggleSearch"
        >
          <MagnifyingGlassIcon class="h-5 w-5" />
        </button>
        <div
          class="overflow-hidden transition-[width] duration-200 ease-out w-0 group-hover:w-56"
          :class="{ '!w-56': searchExpanded }"
        >
          <input
            ref="searchInput"
            v-model="searchQuery"
            type="search"
            placeholder="Поиск по названию, артикулу..."
            autocomplete="off"
            class="w-56 px-3 py-2.5 text-slate-800 placeholder-slate-400 bg-transparent focus:outline-none border-0"
            @keydown.esc="searchExpanded = false"
            @keydown.enter.prevent="doSearch"
          />
        </div>
      </form>
    </div>

    <!-- Floating "Связаться с нами" button - desktop only (mobile: in bottom menu) -->
    <button
      type="button"
      class="hidden lg:flex fixed z-40 items-center justify-center w-12 h-12 rounded-xl border border-slate-200 bg-white text-slate-700 shadow-lg hover:bg-slate-50 hover:text-slate-900 transition-colors right-6 bottom-6"
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
} from '@heroicons/vue/24/outline'
import ContactModal from './components/ContactModal.vue'
import Logo from './components/Logo.vue'

const route = useRoute()
const router = useRouter()
const isHome = computed(() => route.path === '/')
const sidebarHovered = ref(false)
const sidebarExpanded = computed(() => isHome.value || sidebarHovered.value)
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
nav.group:hover .sidebar-logo-compact {
  display: none;
}
nav.group:hover .sidebar-logo-full {
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
/* Expanded: home or hover */
nav.sidebar-expanded .nav-link {
  justify-content: flex-start;
  gap: 0.5rem;
  padding-left: 0.75rem;
  padding-right: 0.75rem;
}
nav.sidebar-expanded .sidebar-link-text {
  opacity: 1;
  width: auto;
}
nav.group:hover .nav-link {
  justify-content: flex-start;
  gap: 0.5rem;
  padding-left: 0.75rem;
  padding-right: 0.75rem;
}
nav.group:hover .sidebar-link-text {
  opacity: 1;
  width: auto;
}
</style>
