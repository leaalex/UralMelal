<template>
  <div>
    <UiPageHeader title="Каталог">
      <template #actions>
        <UiInput v-model="q" placeholder="Поиск..." class="w-48" @keyup.enter="load" />
        <UiSelect v-model="categoryId" @change="load">
          <option value="">Все категории</option>
          <option v-for="c in flatCategories" :key="c.id" :value="c.id">{{ c.name }}</option>
        </UiSelect>
        <UiButton @click="load">Найти</UiButton>
        <UiButton variant="secondary" to="/catalog/import">Импорт</UiButton>
        <UiButton variant="success" @click="openCreate">+ Товар</UiButton>
      </template>
    </UiPageHeader>

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
      <div class="space-y-3">
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
  UiPageHeader,
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
const form = ref({ name: '', sku: '', category_id: '', description: '', image_path: '' })

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
  form.value = { name: '', sku: '', category_id: '', description: '', image_path: '' }
  showModal.value = true
}
function openEdit(p) {
  editing.value = p
  form.value = { name: p.name, sku: p.sku || '', category_id: p.category_id ?? '', description: p.description || '', image_path: p.image_path || '' }
  showModal.value = true
}
async function save() {
  if (!form.value.name) return
  const payload = { ...form.value, category_id: form.value.category_id || null }
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
