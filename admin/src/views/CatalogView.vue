<template>
  <div>
    <div class="flex flex-wrap items-center gap-2 mb-4">
      <UiInput v-model="q" placeholder="Поиск..." class="w-48" @keyup.enter="load" />
      <UiSelect v-model="categoryId" @change="load">
        <option value="">Все категории</option>
        <option v-for="c in flatCategories" :key="c.id" :value="c.id">{{ c.name }}</option>
      </UiSelect>
      <UiButton @click="load">Найти</UiButton>
      <UiButton variant="secondary" to="/catalog/import">Импорт</UiButton>
      <UiButton variant="success" @click="openCreate">+ Товар</UiButton>
    </div>
    <UiTable
      v-if="products.length"
      :columns="tableColumns"
      :data="products"
      row-key="id"
    >
      <template #cell-category="{ value }">
        {{ value?.name ?? '-' }}
      </template>
      <template #actions="{ row }">
        <UiButton variant="ghost-primary" size="sm" class="mr-2" @click="openEdit(row)">Изменить</UiButton>
        <UiButton variant="ghost-danger" size="sm" @click="del(row)">Удалить</UiButton>
      </template>
    </UiTable>
    <UiEmpty v-else message="Нет товаров" />

    <div class="mt-4">
      <UiPagination v-model:page="page" :total="total" :per-page="perPage" />
    </div>

    <UiModal v-model="showModal">
      <template #title>
        <h2 class="text-xl font-bold">{{ editing?.id ? 'Редактировать' : 'Новый товар' }}</h2>
      </template>
      <div class="space-y-4 max-h-[70vh] overflow-y-auto pr-2">
        <div class="space-y-3">
          <h3 class="text-sm font-medium text-slate-600 border-b border-slate-200 pb-1">Основные</h3>
          <UiInput v-model="form.name" label="Название *" />
          <UiInput v-model="form.sku" label="Артикул" />
          <UiSelect v-model="form.category_id" label="Категория">
            <option value="">—</option>
            <option v-for="c in flatCategories" :key="c.id" :value="c.id">{{ c.name }}</option>
          </UiSelect>
          <UiInput v-model="form.description" label="Описание (Markdown)" tag="textarea" :rows="4" />
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-1.5">Изображение</label>
            <div class="flex gap-2">
              <UiInput v-model="form.image_path" placeholder="URL или загрузите" class="flex-1" />
              <input type="file" accept="image/*" @change="uploadFile" class="rounded-lg border border-slate-300 px-3 py-2.5 text-sm" />
            </div>
          </div>
        </div>
        <div class="space-y-3">
          <h3 class="text-sm font-medium text-slate-600 border-b border-slate-200 pb-1">Характеристики</h3>
          <div class="grid grid-cols-2 gap-3">
            <UiInput v-model="form.size" label="Размер" />
            <UiInput v-model="form.mark" label="Марка" />
            <UiInput v-model="form.length" label="Длина" />
            <UiInput v-model="form.city" label="Город" />
          </div>
        </div>
        <div class="space-y-3">
          <h3 class="text-sm font-medium text-slate-600 border-b border-slate-200 pb-1">Цены и наличие</h3>
          <div class="grid grid-cols-2 gap-3">
            <UiInput v-model.number="form.stock" type="number" label="В наличии" />
            <UiInput v-model.number="form.price_1t" type="number" step="0.01" label="Цена за 1 т" />
            <UiInput v-model.number="form.price_5t" type="number" step="0.01" label="Цена за 5 т" />
            <UiInput v-model.number="form.price_10t" type="number" step="0.01" label="Цена за 10 т" />
          </div>
        </div>
        <div class="space-y-3">
          <h3 class="text-sm font-medium text-slate-600 border-b border-slate-200 pb-1">Ссылка</h3>
          <UiInput v-model="form.source_url" label="Ссылка на источник" placeholder="https://..." />
        </div>
        <div class="space-y-3">
          <h3 class="text-sm font-medium text-slate-600 border-b border-slate-200 pb-1">Атрибуты</h3>
          <UiInput v-model="form.characteristics" tag="textarea" :rows="3" placeholder="Дн: 50; Вид: Зонт; Вес, кг: 0,09" label="Характеристики (формат: Ключ: значение; другой: другое)" />
        </div>
      </div>
      <div class="mt-6 flex gap-2">
        <UiButton @click="save">Сохранить</UiButton>
        <UiButton variant="secondary" @click="showModal = false">Отмена</UiButton>
      </div>
    </UiModal>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import {
  UiButton,
  UiEmpty,
  UiInput,
  UiModal,
  UiPagination,
  UiSelect,
  UiTable,
} from '@ui'
import client from '../api/client'

const products = ref([])
const categories = ref([])
const q = ref('')
const categoryId = ref('')
const page = ref(1)
const total = ref(0)
const perPage = 20
const showModal = ref(false)
const editing = ref(null)
const emptyForm = () => ({
  name: '',
  sku: '',
  category_id: '',
  description: '',
  image_path: '',
  size: '',
  mark: '',
  length: '',
  city: '',
  stock: 0,
  price_1t: 0,
  price_5t: 0,
  price_10t: 0,
  source_url: '',
  characteristics: '',
})
const form = ref(emptyForm())

const tableColumns = [
  { key: 'name', label: 'Название' },
  { key: 'sku', label: 'Артикул' },
  { key: 'category', label: 'Категория' },
]

const flatCategories = computed(() => {
  const f = (list, prefix = '') => {
    let out = []
    for (const c of list) {
      out.push({ ...c, name: prefix + c.name })
      if (c.children?.length) out = out.concat(f(c.children, prefix + '— '))
    }
    return out
  }
  return f(categories.value)
})

async function load() {
  try {
    const params = { page: page.value, per_page: perPage }
    if (q.value) params.q = q.value
    if (categoryId.value) params.category_id = categoryId.value
    const { data } = await client.get('/products', { params })
    products.value = data?.products ?? []
    total.value = data?.total ?? 0
  } catch (e) {
    products.value = []
    total.value = 0
    console.error('Ошибка загрузки товаров:', e)
  }
}
async function loadCategories() {
  const { data } = await client.get('/categories')
  categories.value = data
}

async function uploadFile(e) {
  const file = e.target.files?.[0]
  if (!file) return
  const fd = new FormData()
  fd.append('file', file)
  const { data } = await client.post('/admin/upload', fd, { headers: { 'Content-Type': 'multipart/form-data' } })
  form.value.image_path = data.url
}

function openCreate() {
  editing.value = null
  form.value = emptyForm()
  showModal.value = true
}
function openEdit(p) {
  editing.value = p
  const chars = p.attributes?.find((a) => a.key === 'characteristics')
  form.value = {
    name: p.name,
    sku: p.sku || '',
    category_id: p.category_id ?? '',
    description: p.description || '',
    image_path: p.image_path || '',
    size: p.size || '',
    mark: p.mark || '',
    length: p.length || '',
    city: p.city || '',
    stock: p.stock ?? 0,
    price_1t: p.price_1t ?? 0,
    price_5t: p.price_5t ?? 0,
    price_10t: p.price_10t ?? 0,
    source_url: p.source_url || '',
    characteristics: chars?.value || '',
  }
  showModal.value = true
}
function buildPayload() {
  const f = form.value
  const attrs = []
  if (f.characteristics?.trim()) {
    attrs.push({ key: 'characteristics', value: f.characteristics.trim() })
  }
  return {
    name: f.name,
    sku: f.sku || '',
    category_id: f.category_id || null,
    description: f.description || '',
    image_path: f.image_path || '',
    size: f.size || '',
    mark: f.mark || '',
    length: f.length || '',
    city: f.city || '',
    stock: f.stock ?? 0,
    price_1t: f.price_1t ?? 0,
    price_5t: f.price_5t ?? 0,
    price_10t: f.price_10t ?? 0,
    source_url: f.source_url || '',
    attributes: attrs,
  }
}
async function save() {
  if (!form.value.name) return
  const payload = buildPayload()
  if (editing.value) {
    await client.put(`/admin/products/${editing.value.id}`, payload)
  } else {
    await client.post('/admin/products', payload)
  }
  showModal.value = false
  load()
}
async function del(p) {
  if (!confirm('Удалить товар?')) return
  await client.delete(`/admin/products/${p.id}`)
  load()
}

watch(page, load)

onMounted(async () => {
  await loadCategories()
  load()
})
</script>
