<template>
  <div class="w-full px-4 py-6 lg:py-8 overflow-x-hidden min-w-0">
    <!-- Хлебные крошки -->
    <nav class="flex items-center gap-2 text-sm text-slate-500 mb-6 flex-wrap">
      <router-link to="/catalog" class="hover:text-slate-700">Каталог</router-link>
      <template v-if="searchQuery">
        <span class="text-slate-400">/</span>
        <span class="text-slate-800 font-medium">Поиск: «{{ searchQuery }}»</span>
      </template>
      <template v-else v-for="(crumb, i) in breadcrumbs" :key="crumb.id">
        <span class="text-slate-400">/</span>
        <router-link
          v-if="i < breadcrumbs.length - 1"
          :to="crumb.link"
          class="hover:text-slate-700"
        >
          {{ crumb.name }}
        </router-link>
        <span v-else class="text-slate-800 font-medium">{{ crumb.name }}</span>
      </template>
    </nav>

    <!-- Уровень 0: Разделы -->
    <div v-if="catalogLevel === 'sections'" class="space-y-6">
      <h1 class="text-2xl font-bold text-slate-900">Разделы каталога</h1>
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
        <router-link
          v-for="cat in rootCategories"
          :key="cat.id"
          :to="{ path: '/catalog', query: { cat: cat.id } }"
          class="block p-6 rounded-xl border border-slate-200 bg-white hover:border-slate-300 hover:shadow-md text-slate-800"
        >
          <div class="flex items-center gap-3">
            <div class="w-12 h-12 rounded-lg bg-slate-100 flex items-center justify-center shrink-0">
              <FolderIcon class="h-6 w-6 text-slate-500" />
            </div>
            <div class="min-w-0">
              <span class="font-semibold">{{ cat.name }}</span>
              <p v-if="cat.children?.length" class="text-xs text-slate-500 mt-0.5">
                {{ cat.children.length }} подраздел{{ cat.children.length === 1 ? '' : cat.children.length < 5 ? 'а' : 'ов' }}
              </p>
            </div>
            <ChevronRightIcon class="h-5 w-5 text-slate-400 shrink-0 ml-auto" />
          </div>
        </router-link>
      </div>
      <p v-if="!loading && rootCategories.length === 0" class="text-slate-500 py-8 text-center">Разделы не найдены</p>
    </div>

    <!-- Уровень 1: Подразделы -->
    <div v-else-if="catalogLevel === 'subsections'" class="space-y-6">
      <h1 class="text-2xl font-bold text-slate-900">{{ currentCategory?.name }}</h1>
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
        <router-link
          v-for="cat in currentCategory?.children ?? []"
          :key="cat.id"
          :to="{ path: '/catalog', query: { cat: cat.id } }"
          class="block p-6 rounded-xl border border-slate-200 bg-white hover:border-slate-300 hover:shadow-md text-slate-800"
        >
          <div class="flex items-center gap-3">
            <div class="w-12 h-12 rounded-lg bg-slate-100 flex items-center justify-center shrink-0">
              <FolderIcon class="h-6 w-6 text-slate-500" />
            </div>
            <span class="font-semibold">{{ cat.name }}</span>
            <ChevronRightIcon class="h-5 w-5 text-slate-400 shrink-0 ml-auto" />
          </div>
        </router-link>
      </div>
    </div>

    <!-- Уровень 2: Товары -->
    <div v-else-if="catalogLevel === 'products'" class="space-y-6">
      <div class="flex flex-wrap items-center justify-between gap-4">
        <h1 class="text-2xl font-bold text-slate-900">{{ currentCategory?.name }}</h1>
        <div class="flex items-center gap-2">
          <select
            v-model="sort"
            class="rounded-lg border border-slate-300 px-3 py-2.5 text-sm text-slate-700 bg-white focus:ring-2 focus:ring-slate-200 focus:outline-none"
            @change="page = 1; syncToUrl(); load()"
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
              <Squares2X2Icon class="h-5 w-5" />
            </button>
            <button
              type="button"
              :class="viewMode === 'list' ? 'bg-slate-200 text-slate-900' : 'bg-white text-slate-600 hover:bg-slate-50'"
              class="p-2.5"
              title="Список"
              @click="viewMode = 'list'"
            >
              <ListBulletIcon class="h-5 w-5" />
            </button>
          </div>
        </div>
      </div>

      <div v-if="searchQuery" class="flex items-center gap-2 mb-4">
        <span class="text-sm text-slate-500">Поиск:</span>
        <span class="text-sm font-medium text-slate-700">«{{ searchQuery }}»</span>
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

      <div v-if="loading" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
        <div v-for="i in 6" :key="i" class="rounded-xl border border-slate-200 bg-white overflow-hidden">
          <div class="h-44 bg-slate-200" />
          <div class="p-4 space-y-2">
            <div class="h-4 bg-slate-200 rounded w-3/4" />
            <div class="h-3 bg-slate-200 rounded w-1/2" />
          </div>
        </div>
      </div>

      <div
        v-else-if="products.length && viewMode === 'grid'"
        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4"
      >
        <ProductCard
          v-for="p in products"
          :key="p.id"
          :product="p"
          view="grid"
        />
      </div>

      <div
        v-else-if="products.length && viewMode === 'list'"
        class="overflow-x-auto rounded-xl border border-slate-200 bg-white"
      >
        <table class="w-full min-w-[800px]">
          <thead class="bg-slate-50 border-b border-slate-200">
            <tr class="text-left text-xs font-medium text-slate-600 uppercase tracking-wider">
              <th class="px-4 py-3 w-20">Фото</th>
              <th class="px-4 py-3">Название</th>
              <th class="px-4 py-3 w-24">Артикул</th>
              <th class="px-4 py-3 w-20">Размер</th>
              <th class="px-4 py-3 w-24">Марка</th>
              <th class="px-4 py-3 w-24">Город</th>
              <th class="px-4 py-3 w-28">Цена</th>
              <th class="px-4 py-3 w-24">Наличие</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-200">
            <ProductRow
              v-for="p in products"
              :key="p.id"
              :product="p"
            />
          </tbody>
        </table>
      </div>

      <div v-else class="py-16 text-center">
        <p class="text-slate-500 mb-4">Товары не найдены</p>
        <UiButton variant="secondary" @click="router.replace({ path: '/catalog', query: {} })">
          В каталог
        </UiButton>
      </div>

      <div v-if="!loading && products.length" class="mt-8 flex flex-col sm:flex-row items-center justify-between gap-4">
        <p class="text-sm text-slate-500">
          Страница {{ page }} из {{ totalPages }}
        </p>
        <UiPagination v-model:page="page" :total="total" :per-page="perPage" />
      </div>
    </div>

    <!-- Режим поиска (из глобального меню): сразу товары -->
    <div v-else-if="catalogLevel === 'search'" class="space-y-6">
      <h1 class="text-2xl font-bold text-slate-900">Результаты поиска</h1>
      <p v-if="searchQuery" class="text-slate-600">По запросу «{{ searchQuery }}»</p>

      <div class="flex items-center gap-2 mb-4">
        <div class="flex rounded-lg border border-slate-300 overflow-hidden">
          <button
            type="button"
            :class="viewMode === 'grid' ? 'bg-slate-200 text-slate-900' : 'bg-white text-slate-600 hover:bg-slate-50'"
            class="p-2.5"
            title="Сетка"
            @click="viewMode = 'grid'"
          >
            <Squares2X2Icon class="h-5 w-5" />
          </button>
          <button
            type="button"
            :class="viewMode === 'list' ? 'bg-slate-200 text-slate-900' : 'bg-white text-slate-600 hover:bg-slate-50'"
            class="p-2.5"
            title="Список"
            @click="viewMode = 'list'"
          >
            <ListBulletIcon class="h-5 w-5" />
          </button>
        </div>
      </div>

      <div v-if="apiError" class="mb-4 p-4 rounded-lg bg-amber-50 border border-amber-200 text-amber-800">
        {{ apiError }}
        <UiButton variant="secondary" size="sm" class="mt-2" @click="apiError = null; load()">
          Повторить
        </UiButton>
      </div>

      <p v-if="!loading && !apiError" class="text-sm text-slate-500 mb-4">Найдено: {{ total }}</p>

      <div v-if="loading" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
        <div v-for="i in 6" :key="i" class="rounded-xl border border-slate-200 bg-white overflow-hidden">
          <div class="h-44 bg-slate-200" />
          <div class="p-4 space-y-2">
            <div class="h-4 bg-slate-200 rounded w-3/4" />
            <div class="h-3 bg-slate-200 rounded w-1/2" />
          </div>
        </div>
      </div>

      <div
        v-else-if="products.length && viewMode === 'grid'"
        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4"
      >
        <ProductCard v-for="p in products" :key="p.id" :product="p" view="grid" />
      </div>

      <div
        v-else-if="products.length && viewMode === 'list'"
        class="overflow-x-auto rounded-xl border border-slate-200 bg-white"
      >
        <table class="w-full min-w-[800px]">
          <thead class="bg-slate-50 border-b border-slate-200">
            <tr class="text-left text-xs font-medium text-slate-600 uppercase tracking-wider">
              <th class="px-4 py-3 w-20">Фото</th>
              <th class="px-4 py-3">Название</th>
              <th class="px-4 py-3 w-24">Артикул</th>
              <th class="px-4 py-3 w-20">Размер</th>
              <th class="px-4 py-3 w-24">Марка</th>
              <th class="px-4 py-3 w-24">Город</th>
              <th class="px-4 py-3 w-28">Цена</th>
              <th class="px-4 py-3 w-24">Наличие</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-200">
            <ProductRow v-for="p in products" :key="p.id" :product="p" />
          </tbody>
        </table>
      </div>

      <div v-else class="py-16 text-center">
        <p class="text-slate-500 mb-4">По вашему запросу ничего не найдено</p>
        <UiButton variant="secondary" @click="router.replace({ path: '/catalog', query: {} })">В каталог</UiButton>
      </div>

      <div v-if="!loading && products.length" class="mt-8 flex flex-col sm:flex-row items-center justify-between gap-4">
        <p class="text-sm text-slate-500">Страница {{ page }} из {{ totalPages }}</p>
        <UiPagination v-model:page="page" :total="total" :per-page="perPage" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { FolderIcon, ChevronRightIcon, Squares2X2Icon, ListBulletIcon } from '@heroicons/vue/24/outline'
import client from '../api/client'
import ProductCard from '../components/ProductCard.vue'
import ProductRow from '../components/ProductRow.vue'
import { UiButton, UiPagination } from '@ui'
import { useCategories } from '../composables/useCategories'

const route = useRoute()
const router = useRouter()
const { categories, rootCategories, loadCategories, findCategoryById } = useCategories()

const products = ref([])
const categoryId = ref(null)
const searchQuery = ref('')
const page = ref(1)
const total = ref(0)
const perPage = 20
const sort = ref('')
const viewMode = ref('grid')
const loading = ref(false)
const apiError = ref(null)

const totalPages = computed(() => Math.ceil(total.value / perPage) || 1)

const currentCategory = computed(() => {
  if (!categoryId.value) return null
  return findCategoryById(categories.value, categoryId.value)
})

const hasChildren = computed(() => {
  const c = currentCategory.value
  return c?.children && Array.isArray(c.children) && c.children.length > 0
})

const catalogLevel = computed(() => {
  if (searchQuery.value) return 'search'
  if (!categoryId.value) return 'sections'
  if (hasChildren.value) return 'subsections'
  return 'products'
})

const breadcrumbs = computed(() => {
  const result = []
  if (!currentCategory.value || !categoryId.value) return result
  const cats = categories.value
  if (!cats?.length) return result
  const cat = currentCategory.value
  const path = []
  function collect(c, list) {
    if (!c || !c.id) return
    list.unshift(c)
    const parentId = c.parent_id ?? c.parentId
    if (parentId) {
      const parent = findCategoryById(cats, parentId)
      if (parent) collect(parent, list)
    }
  }
  collect(cat, path)
  for (const c of path) {
    if (c && c.id != null && c.name) {
      result.push({
        id: c.id,
        name: c.name,
        link: { path: '/catalog', query: { cat: c.id } },
      })
    }
  }
  return result
})

function syncFromUrl() {
  searchQuery.value = route.query.q ?? ''
  categoryId.value = route.query.cat ? Number(route.query.cat) : null
  page.value = route.query.page ? Math.max(1, Number(route.query.page)) : 1
  sort.value = route.query.sort ?? ''
}

let skipRouteWatch = false
function syncToUrl() {
  skipRouteWatch = true
  const query = {}
  if (searchQuery.value) query.q = searchQuery.value
  if (categoryId.value) query.cat = categoryId.value
  if (page.value > 1) query.page = page.value
  if (sort.value) query.sort = sort.value
  router.replace({ path: '/catalog', query })
}

async function load() {
  if (catalogLevel.value !== 'products' && catalogLevel.value !== 'search') return
  loading.value = true
  apiError.value = null
  try {
    const params = { page: page.value, per_page: perPage }
    if (searchQuery.value) params.q = searchQuery.value
    if (categoryId.value && catalogLevel.value === 'products') params.category_id = categoryId.value
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

watch([categoryId, () => catalogLevel.value], () => {
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
