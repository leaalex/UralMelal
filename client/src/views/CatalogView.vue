<template>
  <div class="container mx-auto px-4 py-6 lg:py-8 overflow-x-hidden">
    <div class="flex flex-col lg:flex-row gap-6 lg:items-start">
      <CategorySidebar
        v-model="categoryId"
        v-model:open="sidebarOpen"
        :categories="categories"
      />
      <div class="flex-1 min-w-0">
        <div class="flex flex-wrap items-center justify-between gap-4 mb-6">
          <h1 class="text-2xl font-bold text-slate-900">Каталог</h1>
          <UiButton
            variant="secondary"
            class="lg:hidden"
            @click="sidebarOpen = true"
          >
            <template #icon>
              <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
              </svg>
            </template>
            Категории
          </UiButton>
        </div>

        <div class="flex flex-col sm:flex-row gap-3 mb-4">
          <div class="relative flex-1">
            <input
              v-model="q"
              type="search"
              placeholder="Поиск по названию, артикулу..."
              class="w-full rounded-lg border border-slate-300 px-4 py-2.5 pl-10 text-slate-800 placeholder-slate-400 focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 focus:outline-none"
              @input="debouncedLoad"
            />
            <svg
              class="absolute left-3 top-1/2 -translate-y-1/2 h-5 w-5 text-slate-400"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
          <div class="flex items-center gap-2">
            <select
              v-model="sort"
              class="rounded-lg border border-slate-300 px-3 py-2.5 text-sm text-slate-700 bg-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 focus:outline-none"
            >
              <option value="">По умолчанию</option>
              <option value="name_asc">По названию (А–Я)</option>
              <option value="name_desc">По названию (Я–А)</option>
              <option value="price_asc">Сначала дешевле</option>
              <option value="price_desc">Сначала дороже</option>
            </select>
            <div class="flex rounded-lg border border-slate-300 overflow-hidden">
              <button
                type="button"
                :class="viewMode === 'grid' ? 'bg-slate-200 text-slate-900' : 'bg-white text-slate-600 hover:bg-slate-50'"
                class="p-2.5"
                title="Сетка"
                @click="viewMode = 'grid'"
              >
                <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
                </svg>
              </button>
              <button
                type="button"
                :class="viewMode === 'list' ? 'bg-slate-200 text-slate-900' : 'bg-white text-slate-600 hover:bg-slate-50'"
                class="p-2.5"
                title="Список"
                @click="viewMode = 'list'"
              >
                <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16" />
                </svg>
              </button>
            </div>
          </div>
        </div>

        <div v-if="activeFilters.length" class="flex flex-wrap items-center gap-2 mb-4">
          <span class="text-sm text-slate-500">Фильтры:</span>
          <button
            v-for="f in activeFilters"
            :key="f.key"
            type="button"
            class="inline-flex items-center gap-1 rounded-full bg-slate-100 px-3 py-1 text-sm text-slate-700 hover:bg-slate-200"
            @click="clearFilter(f.key)"
          >
            {{ f.label }}
            <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <div v-if="apiError" class="mb-4 p-4 rounded-lg bg-amber-50 border border-amber-200 text-amber-800">
          {{ apiError }}
          <UiButton variant="secondary" size="sm" class="mt-2" @click="apiError = null; load(); loadCategories()">
            Повторить
          </UiButton>
        </div>
        <p v-if="!loading && !apiError" class="text-sm text-slate-500 mb-4">
          Найдено: {{ total }}
        </p>

        <div v-if="loading" class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-4">
          <div v-for="i in 6" :key="i" class="rounded-xl border border-slate-200 bg-white overflow-hidden animate-pulse">
            <div class="h-44 bg-slate-200" />
            <div class="p-4 space-y-2">
              <div class="h-4 bg-slate-200 rounded w-3/4" />
              <div class="h-3 bg-slate-200 rounded w-1/2" />
              <div class="h-3 bg-slate-200 rounded w-1/3" />
            </div>
          </div>
        </div>

        <div
          v-else-if="products.length"
          :class="viewMode === 'list' ? 'space-y-3' : 'grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-4'"
        >
          <ProductCard
            v-for="p in products"
            :key="p.id"
            :product="p"
            :view="viewMode"
          />
        </div>

        <div v-else class="py-16 text-center">
          <p class="text-slate-500 mb-4">Товары не найдены</p>
          <UiButton variant="secondary" @click="clearAllFilters">
            Сбросить фильтры
          </UiButton>
        </div>

        <div v-if="!loading && products.length" class="mt-8 flex flex-col sm:flex-row items-center justify-between gap-4">
          <p class="text-sm text-slate-500">
            Страница {{ page }} из {{ totalPages }}
          </p>
          <UiPagination v-model:page="page" :total="total" :per-page="perPage" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import client from '../api/client'
import ProductCard from '../components/ProductCard.vue'
import CategorySidebar from '../components/CategorySidebar.vue'
import { UiButton, UiPagination } from '@ui'

const route = useRoute()
const router = useRouter()

const products = ref([])
const categories = ref([])
const q = ref('')
const categoryId = ref(null)
const page = ref(1)
const total = ref(0)
const perPage = 20
const sort = ref('')
const viewMode = ref('grid')
const loading = ref(false)
const sidebarOpen = ref(false)
const apiError = ref(null)

let debounceTimer = null

const totalPages = computed(() => Math.ceil(total.value / perPage) || 1)

const activeFilters = computed(() => {
  const filters = []
  if (q.value) filters.push({ key: 'q', label: `Поиск: "${q.value}"` })
  if (categoryId.value) {
    const cat = findCategoryById(categories.value, categoryId.value)
    filters.push({ key: 'cat', label: cat?.name ?? `Категория #${categoryId.value}` })
  }
  return filters
})

function findCategoryById(list, id) {
  if (!list?.length) return null
  for (const c of list) {
    if (c.id === id) return c
    const found = findCategoryById(c.children, id)
    if (found) return found
  }
  return null
}

function clearFilter(key) {
  if (key === 'q') q.value = ''
  if (key === 'cat') categoryId.value = null
  page.value = 1
  syncToUrl()
  load()
}

function clearAllFilters() {
  q.value = ''
  categoryId.value = null
  page.value = 1
  syncToUrl()
  load()
}

function syncFromUrl() {
  q.value = route.query.q ?? ''
  categoryId.value = route.query.cat ? Number(route.query.cat) : null
  page.value = route.query.page ? Math.max(1, Number(route.query.page)) : 1
  sort.value = route.query.sort ?? ''
}

let skipRouteWatch = false
function syncToUrl() {
  skipRouteWatch = true
  const query = {}
  if (q.value) query.q = q.value
  if (categoryId.value) query.cat = categoryId.value
  if (page.value > 1) query.page = page.value
  if (sort.value) query.sort = sort.value
  router.replace({ path: '/catalog', query })
}

async function loadCategories() {
  try {
    const { data } = await client.get('/categories')
    categories.value = data ?? []
    return true
  } catch (e) {
    apiError.value = 'Не удалось загрузить категории. Проверьте подключение к серверу.'
    return false
  }
}

async function load() {
  loading.value = true
  apiError.value = null
  try {
    const params = { page: page.value, per_page: perPage }
    if (q.value) params.q = q.value
    if (categoryId.value) params.category_id = categoryId.value
    if (sort.value) params.sort = sort.value
    const { data } = await client.get('/products', { params })
    products.value = data?.products ?? []
    total.value = data?.total ?? 0
  } catch (e) {
    products.value = []
    total.value = 0
    apiError.value = 'Не удалось загрузить товары. Проверьте подключение к серверу.'
  } finally {
    loading.value = false
  }
}

function debouncedLoad() {
  clearTimeout(debounceTimer)
  debounceTimer = setTimeout(() => {
    page.value = 1
    syncToUrl()
    load()
  }, 300)
}

watch([categoryId, sort], () => {
  page.value = 1
  syncToUrl()
  load()
})

watch(page, () => {
  syncToUrl()
  load()
})

watch(() => route.query, () => {
  if (skipRouteWatch) {
    skipRouteWatch = false
    return
  }
  syncFromUrl()
  load()
}, { deep: true })

onMounted(async () => {
  syncFromUrl()
  await loadCategories()
  load()
})
</script>
