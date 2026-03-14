<template>
  <div class="min-h-screen flex flex-col bg-slate-50 text-slate-800">
    <!-- Mobile: top header with logo -->
    <header class="lg:hidden fixed top-4 left-4 right-4 z-50 flex items-center justify-center px-4 py-5 rounded-xl bg-white/95 backdrop-blur-xl border border-slate-200 shadow-lg">
      <Logo size="lg" />
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
            <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
            </svg>
          </span>
          <span class="whitespace-nowrap text-sm pl-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200">Главная</span>
        </router-link>
        <router-link
          to="/catalog"
          active-class="!bg-slate-100 !text-slate-900"
          class="flex items-center h-12 w-full min-w-12 overflow-hidden rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors duration-200"
        >
          <span class="shrink-0 w-12 flex items-center justify-center">
            <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
            </svg>
          </span>
          <span class="whitespace-nowrap text-sm pl-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200">Каталог</span>
        </router-link>
        <router-link
          to="/contacts"
          exact-active-class="!bg-slate-100 !text-slate-900"
          class="flex items-center h-12 w-full min-w-12 overflow-hidden rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors duration-200"
        >
          <span class="shrink-0 w-12 flex items-center justify-center">
            <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
            </svg>
          </span>
          <span class="whitespace-nowrap text-sm pl-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200">Контакты</span>
        </router-link>
      </div>
      <div class="flex flex-col items-start gap-2 py-4 px-2 border-t border-slate-200 shrink-0 w-56">
        <div class="flex items-center h-12 w-full min-w-12 overflow-hidden rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors duration-200">
          <a :href="`tel:${telHref}`" class="shrink-0 w-12 h-12 flex items-center justify-center" aria-label="Позвонить">
            <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.5a2 2 0 012 2v14a2 2 0 01-2 2H5a2 2 0 01-2-2V5zm5.5 2a1.5 1.5 0 100 3h7a1.5 1.5 0 100-3h-7z" />
            </svg>
          </a>
          <span class="whitespace-nowrap text-sm pl-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200 truncate">{{ contactPhone }}</span>
          <button
            type="button"
            class="shrink-0 ml-1 p-1.5 rounded opacity-0 group-hover:opacity-100 transition-opacity duration-200 hover:bg-slate-200"
            aria-label="Копировать"
            @click.stop="copyToClipboard(contactPhone)"
          >
            <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
          </button>
        </div>
        <div class="flex items-center h-12 w-full min-w-12 overflow-hidden rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors duration-200">
          <a :href="`mailto:${contactEmail}`" class="shrink-0 w-12 h-12 flex items-center justify-center" aria-label="Написать">
            <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
            </svg>
          </a>
          <span class="whitespace-nowrap text-sm pl-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200 truncate">{{ contactEmail }}</span>
          <button
            type="button"
            class="shrink-0 ml-1 p-1.5 rounded opacity-0 group-hover:opacity-100 transition-opacity duration-200 hover:bg-slate-200"
            aria-label="Копировать"
            @click.stop="copyToClipboard(contactEmail)"
          >
            <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
          </button>
        </div>
      </div>
    </nav>

    <!-- Mobile: bottom bar -->
    <nav class="lg:hidden fixed bottom-4 left-4 right-4 flex flex-col bg-white/95 backdrop-blur-xl text-slate-700 border border-slate-200 z-50 rounded-xl overflow-hidden shadow-xl min-h-[7rem]">
      <div class="flex flex-shrink-0 py-2">
        <router-link
          to="/"
          exact-active-class="!bg-slate-100 !text-slate-900"
          class="flex flex-col items-center justify-center flex-1 text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors py-2 shadow-lg"
        >
          <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
          </svg>
          <span class="text-xs mt-0.5">Главная</span>
        </router-link>
        <router-link
          to="/catalog"
          active-class="!bg-slate-100 !text-slate-900"
          class="flex flex-col items-center justify-center flex-1 text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors py-2 shadow-lg"
        >
          <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
          </svg>
          <span class="text-xs mt-0.5">Каталог</span>
        </router-link>
        <router-link
          to="/contacts"
          exact-active-class="!bg-slate-100 !text-slate-900"
          class="flex flex-col items-center justify-center flex-1 text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition-colors py-2 shadow-lg"
        >
          <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
          </svg>
          <span class="text-xs mt-0.5">Контакты</span>
        </router-link>
      </div>
      <div class="flex items-center justify-center gap-4 px-4 py-2.5 border-t border-slate-200 shrink-0">
        <div class="flex items-center gap-1.5">
          <a
            :href="`tel:${telHref}`"
            class="flex items-center gap-1.5 text-slate-700 hover:text-slate-900 transition-colors text-sm font-medium"
            aria-label="Позвонить"
          >
            <svg class="h-5 w-5 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.5a2 2 0 012 2v14a2 2 0 01-2 2H5a2 2 0 01-2-2V5zm5.5 2a1.5 1.5 0 100 3h7a1.5 1.5 0 100-3h-7z" />
            </svg>
            <span class="truncate max-w-[100px]">{{ contactPhone }}</span>
          </a>
          <button
            type="button"
            class="p-1.5 rounded text-slate-500 hover:text-slate-800 hover:bg-slate-100 transition-colors"
            aria-label="Копировать телефон"
            @click.stop="copyToClipboard(contactPhone)"
          >
            <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
          </button>
        </div>
        <div class="flex items-center gap-1.5">
          <a
            :href="`mailto:${contactEmail}`"
            class="flex items-center gap-1.5 text-slate-700 hover:text-slate-900 transition-colors text-sm font-medium"
            aria-label="Написать"
          >
            <svg class="h-5 w-5 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
            </svg>
            <span class="truncate max-w-[100px]">{{ contactEmail }}</span>
          </a>
          <button
            type="button"
            class="p-1.5 rounded text-slate-500 hover:text-slate-800 hover:bg-slate-100 transition-colors"
            aria-label="Копировать почту"
            @click.stop="copyToClipboard(contactEmail)"
          >
            <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
          </button>
        </div>
      </div>
    </nav>

    <main class="flex-1 pt-20 pb-36 lg:pt-0 lg:pb-0 lg:ml-20">
      <router-view />
    </main>
    <footer class="bg-white border-t border-slate-200 py-6 mt-auto">
      <div class="container mx-auto px-4 text-center">
        <p class="text-sm text-slate-600">© {{ currentYear }} Урал Металл</p>
      </div>
    </footer>

    <ContactModal v-model="contactModalOpen" />

    <!-- Search button + panel - top right, panel appears on hover -->
    <div class="fixed top-4 right-4 z-50 flex justify-end">
      <form
        class="group flex flex-row-reverse items-center rounded-xl border border-slate-200 bg-white shadow-lg overflow-hidden hover:ring-2 hover:ring-slate-200 transition-shadow"
        @submit.prevent="doSearch"
      >
        <button
          type="button"
          class="shrink-0 flex items-center justify-center w-12 h-12 text-slate-700 hover:bg-slate-50 hover:text-slate-900 transition-colors"
          aria-label="Поиск"
          @click.prevent="searchExpanded = !searchExpanded"
        >
          <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </button>
        <div
          class="overflow-hidden transition-[width] duration-200 ease-out w-0 group-hover:w-56"
          :class="{ '!w-56': searchExpanded }"
        >
          <input
            v-model="searchQuery"
            type="search"
            placeholder="Поиск..."
            class="w-56 px-3 py-2.5 text-slate-800 placeholder-slate-400 bg-transparent focus:outline-none border-0"
            @keydown.esc="searchExpanded = false"
          />
        </div>
      </form>
    </div>

    <!-- Floating "Связаться с нами" button - bottom right, icon only -->
    <button
      type="button"
      class="fixed z-40 flex items-center justify-center w-12 h-12 rounded-xl border border-slate-200 bg-white text-slate-700 shadow-lg hover:bg-slate-50 hover:text-slate-900 transition-colors right-4 bottom-28 lg:right-6 lg:bottom-6"
      aria-label="Связаться с нами"
      @click="contactModalOpen = true"
    >
      <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
      </svg>
    </button>
  </div>
</template>

<script setup>
import { ref, computed, provide } from 'vue'
import { useRouter } from 'vue-router'
import ContactModal from './components/ContactModal.vue'
import Logo from './components/Logo.vue'

const router = useRouter()
const contactModalOpen = ref(false)
const searchQuery = ref('')
const searchExpanded = ref(false)

function doSearch() {
  const q = searchQuery.value?.trim()
  if (q) {
    router.push({ path: '/catalog', query: { q } })
    searchExpanded.value = false
    searchQuery.value = ''
  }
}
const currentYear = new Date().getFullYear()
provide('openContactModal', () => { contactModalOpen.value = true })
const contactPhone = ref(import.meta.env.VITE_CONTACT_PHONE || '+7 (XXX) XXX-XX-XX')
const contactEmail = ref(import.meta.env.VITE_CONTACT_EMAIL || 'info@uralmetall.ru')

const telHref = computed(() => contactPhone.value.replace(/\D/g, ''))

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
