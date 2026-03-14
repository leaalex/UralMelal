<template>
  <div class="container mx-auto px-4 py-6 lg:py-8">
    <div v-if="loading" class="max-w-4xl mx-auto">
      <div class="animate-pulse space-y-6">
        <div class="h-4 bg-slate-200 rounded w-48" />
        <div class="flex flex-col md:flex-row gap-6">
          <div class="md:w-1/2 h-80 bg-slate-200 rounded-xl" />
          <div class="md:w-1/2 space-y-4">
            <div class="h-8 bg-slate-200 rounded w-3/4" />
            <div class="h-4 bg-slate-200 rounded w-1/2" />
            <div class="h-24 bg-slate-200 rounded" />
          </div>
        </div>
      </div>
    </div>
    <div v-else-if="product" class="max-w-5xl mx-auto">
      <nav class="flex items-center gap-2 text-sm text-slate-500 mb-6">
        <router-link to="/catalog" class="hover:text-slate-700">Каталог</router-link>
        <template v-for="(crumb, i) in breadcrumbs" :key="crumb.id">
          <span class="text-slate-400">/</span>
          <router-link
            v-if="i < breadcrumbs.length - 1"
            :to="{ path: '/catalog', query: { cat: crumb.id } }"
            class="hover:text-slate-700"
          >
            {{ crumb.name }}
          </router-link>
          <span v-else class="text-slate-800 font-medium">{{ crumb.name }}</span>
        </template>
      </nav>

      <div class="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
        <div class="flex flex-col md:flex-row">
          <div class="md:w-1/2 h-64 md:h-96">
            <img
              v-if="product.image_path"
              :src="product.image_path"
              :alt="product.name"
              class="w-full h-full object-cover"
              loading="lazy"
            />
            <div
              v-else
              class="w-full h-full bg-slate-100 flex items-center justify-center text-slate-400"
            >
              Нет фото
            </div>
          </div>
          <div class="md:w-1/2 p-6 lg:p-8 flex flex-col">
            <div class="flex flex-wrap gap-2 mb-3">
              <span
                v-if="product.size"
                class="inline-flex items-center rounded-md bg-slate-100 px-2.5 py-0.5 text-sm font-medium text-slate-700"
              >
                {{ product.size }}
              </span>
              <span
                v-if="product.mark"
                class="inline-flex items-center rounded-md bg-slate-100 px-2.5 py-0.5 text-sm font-medium text-slate-700"
              >
                {{ product.mark }}
              </span>
              <span
                v-if="hasStock"
                class="inline-flex items-center gap-1 rounded-md bg-emerald-50 px-2.5 py-0.5 text-sm font-medium text-emerald-700"
              >
                <span class="h-1.5 w-1.5 rounded-full bg-emerald-500" />
                В наличии
              </span>
              <span
                v-else-if="product.stock !== undefined"
                class="inline-flex items-center rounded-md bg-slate-100 px-2.5 py-0.5 text-sm text-slate-500"
              >
                Нет в наличии
              </span>
            </div>

            <h1 class="text-2xl font-bold text-slate-800 mb-2">{{ product.name }}</h1>
            <p v-if="product.sku" class="text-slate-500 mb-4">Артикул: {{ product.sku }}</p>
            <p v-if="product.city" class="text-sm text-slate-500 mb-4">{{ product.city }}</p>

            <div v-if="hasAnyPrice" class="mb-6">
              <h3 class="font-semibold text-slate-800 mb-2">Цены</h3>
              <div class="overflow-hidden rounded-lg border border-slate-200">
                <table class="w-full text-sm">
                  <thead class="bg-slate-50">
                    <tr>
                      <th class="px-4 py-2.5 text-left font-medium text-slate-700">Св. 1 т</th>
                      <th class="px-4 py-2.5 text-left font-medium text-slate-700">Св. 5 т</th>
                      <th class="px-4 py-2.5 text-left font-medium text-slate-700">Св. 10 т</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr class="divide-x divide-slate-200">
                      <td class="px-4 py-2.5 text-slate-800">{{ formatPrice(product.price_1t) || '—' }}</td>
                      <td class="px-4 py-2.5 text-slate-800">{{ formatPrice(product.price_5t) || '—' }}</td>
                      <td class="px-4 py-2.5 text-slate-800">{{ formatPrice(product.price_10t) || '—' }}</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>

            <div v-if="specs.length" class="mb-6">
              <h3 class="font-semibold text-slate-800 mb-2">Характеристики</h3>
              <dl class="space-y-1.5 text-sm">
                <div
                  v-for="s in specs"
                  :key="s.key"
                  class="flex justify-between gap-4 py-1.5 border-b border-slate-100 last:border-0"
                >
                  <dt class="text-slate-600">{{ s.key }}</dt>
                  <dd class="text-slate-800 text-right">{{ s.value }}</dd>
                </div>
              </dl>
            </div>

            <div v-if="product.description" class="mb-6 prose prose-slate prose-sm max-w-none">
              <MarkdownContent :content="product.description" />
            </div>

            <div class="mt-auto">
              <UiButton size="lg" class="w-full sm:w-auto" @click="requestPrice">
                Запросить цену
              </UiButton>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-else class="text-center py-16">
      <p class="text-slate-500 mb-4">Товар не найден</p>
      <router-link to="/catalog" class="text-blue-600 hover:underline">Вернуться в каталог</router-link>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import client from '../api/client'
import { MarkdownContent, UiButton } from '@ui'
import { formatPrice, parseSpecs } from '../utils/format'

const route = useRoute()
const router = useRouter()

const product = ref(null)
const loading = ref(true)

const breadcrumbs = computed(() => {
  const crumbs = []
  const cat = product.value?.category
  if (cat) crumbs.push({ id: cat.id, name: cat.name })
  if (product.value?.name) crumbs.push({ id: product.value.id, name: product.value.name })
  return crumbs
})

const specs = computed(() => {
  const attr = product.value?.attributes?.find((a) => a.key === 'characteristics')
  return attr?.value ? parseSpecs(attr.value) : []
})

const hasStock = computed(() => {
  const s = product.value?.stock
  return s != null && s > 0
})

const hasAnyPrice = computed(() => {
  const p = product.value
  return p && (p.price_1t > 0 || p.price_5t > 0 || p.price_10t > 0)
})

function requestPrice() {
  router.push({
    path: route.path,
    query: { ...route.query, contact_subject: product.value?.name ?? '' },
    hash: 'contact',
  })
  nextTick(() => {
    document.getElementById('contact')?.scrollIntoView({ behavior: 'smooth' })
  })
}

onMounted(async () => {
  try {
    const { data } = await client.get(`/products/${route.params.id}`)
    product.value = data
  } finally {
    loading.value = false
  }
})
</script>
