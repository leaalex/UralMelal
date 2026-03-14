<template>
  <div class="min-h-screen flex flex-col bg-slate-50 text-slate-800">
    <!-- Mobile: logo in top-left over video -->
    <header class="lg:hidden fixed top-4 left-4 z-[60] drop-shadow-lg">
      <router-link to="/" class="block">
        <Logo size="lg" />
      </router-link>
    </header>

    <!-- Desktop: vertical sidebar left, expands fully on hover -->
    <nav class="hidden lg:flex fixed top-4 left-4 bottom-4 w-16 group hover:w-56 flex-col bg-white/95 backdrop-blur-xl text-slate-700 border border-slate-200 z-50 rounded-xl overflow-hidden shadow-xl transition-[width] duration-200 ease-out">
      <div class="shrink-0 pt-6 px-2 pb-5 w-56 border-b border-slate-200 flex items-center min-h-[4.5rem]">
        <!-- Узкое меню: logo.png — по ширине узкой полосы (центрирован) -->
        <div class="sidebar-logo-compact flex justify-center items-center w-12 min-w-12 h-10">
          <Logo compact size="sm" class="sidebar-logo-img" />
        </div>
        <!-- Развёрнутое меню: logo-full.png — выравнивание по левому краю -->
        <div class="sidebar-logo-full w-full flex justify-start items-center h-10">
          <Logo size="sm" class="sidebar-logo-img" />
        </div>
      </div>
      <div class="flex-1 flex flex-col items-start gap-2 pt-4 px-2 w-56 overflow-y-auto">
        <router-link
          to="/"
          exact-active-class="!bg-slate-100 !text-slate-900"
          class="flex items-center h-12 w-full min-w-12 overflow-hidden rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors duration-200"
        >
          <span class="shrink-0 w-12 flex items-center justify-center">
            <HomeIcon class="h-6 w-6" />
          </span>
          <span class="whitespace-nowrap text-sm pl-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200">Главная</span>
        </router-link>
        <router-link
          to="/catalog"
          active-class="!bg-slate-100 !text-slate-900"
          class="flex items-center h-12 w-full min-w-12 overflow-hidden rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors duration-200"
        >
          <span class="shrink-0 w-12 flex items-center justify-center">
            <CubeIcon class="h-6 w-6" />
          </span>
          <span class="whitespace-nowrap text-sm pl-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200">Каталог</span>
        </router-link>
        <router-link
          to="/contacts"
          exact-active-class="!bg-slate-100 !text-slate-900"
          class="flex items-center h-12 w-full min-w-12 overflow-hidden rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors duration-200"
        >
          <span class="shrink-0 w-12 flex items-center justify-center">
            <EnvelopeIcon class="h-6 w-6" />
          </span>
          <span class="whitespace-nowrap text-sm pl-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200">О компании</span>
        </router-link>
      </div>
      <div class="flex flex-col items-start gap-2 py-4 px-2 border-t border-slate-200 shrink-0 w-56">
        <div class="flex items-center h-12 w-full min-w-12 overflow-hidden rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors duration-200">
          <a :href="`tel:${telHref}`" class="shrink-0 w-12 h-12 flex items-center justify-center" aria-label="Позвонить">
            <PhoneIcon class="h-6 w-6" />
          </a>
          <span class="whitespace-nowrap text-sm pl-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200 truncate">{{ contactPhone }}</span>
          <button
            type="button"
            class="shrink-0 ml-1 p-1.5 rounded opacity-0 group-hover:opacity-100 transition-opacity duration-200 hover:bg-slate-200"
            aria-label="Копировать"
            @click.stop="copyToClipboard(contactPhone)"
          >
            <DocumentDuplicateIcon class="h-4 w-4" />
          </button>
        </div>
        <div class="flex items-center h-12 w-full min-w-12 overflow-hidden rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors duration-200">
          <a :href="`mailto:${contactEmail}`" class="shrink-0 w-12 h-12 flex items-center justify-center" aria-label="Написать">
            <EnvelopeIcon class="h-6 w-6" />
          </a>
          <span class="whitespace-nowrap text-sm pl-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200 truncate">{{ contactEmail }}</span>
          <button
            type="button"
            class="shrink-0 ml-1 p-1.5 rounded opacity-0 group-hover:opacity-100 transition-opacity duration-200 hover:bg-slate-200"
            aria-label="Копировать"
            @click.stop="copyToClipboard(contactEmail)"
          >
            <DocumentDuplicateIcon class="h-4 w-4" />
          </button>
        </div>
      </div>
    </nav>

    <!-- Mobile: bottom bar - two panels -->
    <div class="lg:hidden fixed bottom-4 left-4 right-4 flex gap-2 z-50">
      <!-- Панель 1: Главная, Каталог, О компании -->
      <nav class="flex-1 flex bg-white/95 backdrop-blur-xl text-slate-700 border border-slate-200 rounded-xl overflow-hidden shadow-xl">
      <router-link
        to="/"
        exact-active-class="!bg-slate-100 !text-slate-900"
        class="flex flex-col items-center justify-center flex-1 py-3 text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors"
      >
          <HomeIcon class="h-6 w-6" />
        <span class="text-xs mt-0.5">Главная</span>
      </router-link>
      <router-link
        to="/catalog"
        active-class="!bg-slate-100 !text-slate-900"
        class="flex flex-col items-center justify-center flex-1 py-3 text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors"
      >
          <CubeIcon class="h-6 w-6" />
        <span class="text-xs mt-0.5">Каталог</span>
      </router-link>
      <router-link
        to="/contacts"
        exact-active-class="!bg-slate-100 !text-slate-900"
        class="flex flex-col items-center justify-center flex-1 py-3 text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors"
      >
          <EnvelopeIcon class="h-6 w-6" />
        <span class="text-xs mt-0.5">О компании</span>
      </router-link>
      </nav>
      <!-- Панель 2: Связаться -->
      <button
        type="button"
        class="flex flex-col items-center justify-center px-5 py-3 bg-white/95 backdrop-blur-xl text-slate-600 border border-slate-200 rounded-xl shadow-xl hover:bg-slate-100 hover:text-slate-900 transition-colors"
        @click="contactModalOpen = true"
      >
        <ChatBubbleLeftRightIcon class="h-6 w-6" />
        <span class="text-xs mt-0.5">Связаться</span>
      </button>
    </div>

    <main :class="['flex-1 pb-36 lg:pt-0 lg:pb-0 lg:ml-20', route.path === '/' ? 'pt-0' : 'pt-20']">
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
/* Единая высота логотипов в меню — h-10 (40px) */
.sidebar-logo-img img {
  height: 2.5rem;
  width: auto;
  max-height: 2.5rem;
}
/* Узкое меню: показываем logo.png, скрываем logo-full */
.sidebar-logo-full {
  display: none;
}
nav.group:hover .sidebar-logo-compact {
  display: none;
}
nav.group:hover .sidebar-logo-full {
  display: flex;
}
</style>
