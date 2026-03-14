<template>
  <div class="flex flex-col flex-1 min-h-0">
    <div class="flex gap-4 min-h-[calc(100vh-2rem)]">
      <!-- Left: tree (высота как меню: top-4 bottom-4) -->
      <aside class="w-72 shrink-0 flex flex-col border border-slate-200 rounded-xl bg-white overflow-hidden h-[calc(100vh-2rem)]">
        <div class="shrink-0 p-2 border-b border-slate-200 flex items-center justify-between gap-2">
          <span class="text-sm font-medium text-slate-600">Список категорий</span>
          <UiButton size="md" variant="secondary" @click="openCreate" :aria-label="'Добавить категорию'">
          <PlusIcon class="h-5 w-5" />
        </UiButton>
        </div>
        <div class="flex-1 overflow-y-auto p-2 min-h-0">
          <CategoryTreeNode
            v-for="c in categories"
            :key="c.id"
            :node="c"
            :selected-id="selectedId"
            :depth="0"
            @select="selectedId = $event"
          />
          <UiEmpty v-if="!categories.length" message="Нет категорий" class="py-8" />
        </div>
      </aside>

      <!-- Right: form / details (высота как меню, контент прокручивается внутри) -->
      <section class="flex-1 min-w-0 flex flex-col border border-slate-200 rounded-xl bg-white overflow-hidden h-[calc(100vh-2rem)]">
        <div class="flex flex-col flex-1 min-h-0 p-6">
          <div v-if="selectedId === null" class="flex flex-col items-center justify-center py-16 text-slate-500 flex-1">
            <p>Выберите категорию или создайте новую</p>
          </div>

          <template v-else>
            <!-- Breadcrumbs + Tabs (вкладки справа в виде иконок) -->
            <div class="flex items-center justify-between gap-4 mb-4 shrink-0">
              <!-- Breadcrumbs слева -->
              <nav v-if="breadcrumbs.length" class="flex items-center gap-2 text-sm text-slate-500 min-w-0">
                <span
                  v-for="(crumb, i) in breadcrumbs"
                  :key="crumb.id"
                  class="flex items-center gap-2 shrink-0"
                >
                  <button
                    v-if="i > 0"
                    type="button"
                    class="hover:text-slate-700"
                    @click="selectedId = crumb.id"
                  >
                    {{ crumb.name }}
                  </button>
                  <span v-else>{{ crumb.name }}</span>
                  <span v-if="i < breadcrumbs.length - 1" class="text-slate-300">/</span>
                </span>
              </nav>
              <h3 v-else-if="selectedId === 'new'" class="text-lg font-semibold text-slate-800">Новая категория</h3>

              <!-- Кнопки-вкладки справа (объединённая группа с иконками и подписями) -->
              <div v-if="selectedId !== 'new'" class="flex rounded-lg border border-slate-200 overflow-hidden shrink-0">
                <button
                  type="button"
                  class="flex items-center gap-2 px-3 py-2 text-sm transition-colors"
                  :class="activeTab === 'description' ? 'bg-slate-200 text-slate-900' : 'bg-white text-slate-500 hover:bg-slate-100 hover:text-slate-700'"
                  @click="activeTab = 'description'"
                >
                  <DocumentTextIcon class="h-4 w-4 shrink-0" />
                  <span>Описание</span>
                </button>
                <button
                  type="button"
                  class="flex items-center gap-2 px-3 py-2 text-sm border-l border-slate-200 transition-colors"
                  :class="activeTab === 'subcategories' ? 'bg-slate-200 text-slate-900' : 'bg-white text-slate-500 hover:bg-slate-100 hover:text-slate-700'"
                  @click="activeTab = 'subcategories'"
                >
                  <FolderIcon class="h-4 w-4 shrink-0" />
                  <span>Подкатегории</span>
                </button>
                <button
                  type="button"
                  class="flex items-center gap-2 px-3 py-2 text-sm border-l border-slate-200 transition-colors"
                  :class="activeTab === 'products' ? 'bg-slate-200 text-slate-900' : 'bg-white text-slate-500 hover:bg-slate-100 hover:text-slate-700'"
                  @click="activeTab = 'products'"
                >
                  <CubeIcon class="h-4 w-4 shrink-0" />
                  <span>Товары</span>
                </button>
              </div>
            </div>

            <!-- Tab content -->
            <div class="flex-1 overflow-y-auto min-h-0">
              <!-- Описание -->
              <div v-show="activeTab === 'description'" class="space-y-4">
                <UiInput v-model="form.name" label="Название *" />
                <UiSelect v-model="form.parent_id" label="Родитель">
                  <option value="">Корень (без родителя)</option>
                  <option
                    v-for="c in flatCategories"
                    :key="c.id"
                    :value="c.id"
                    :disabled="selectedCategory?.id === c.id || isDescendantOfSelected(c)"
                  >
                    {{ c.name }}
                  </option>
                </UiSelect>
                <UiInput v-model.number="form.sort_order" type="number" label="Порядок сортировки" />

                <div class="flex justify-between pt-2">
                  <div>
                    <UiButton v-if="selectedId !== 'new'" variant="ghost-danger" @click="del">Удалить</UiButton>
                  </div>
                  <UiButton @click="save">Сохранить</UiButton>
                </div>
              </div>

              <!-- Подкатегории -->
              <div v-show="activeTab === 'subcategories' && selectedId !== 'new'" class="space-y-4">
                <div class="flex items-center justify-between mb-2">
                  <span class="text-sm text-slate-600">Подкатегории</span>
                  <UiButton size="md" variant="secondary" @click="openAddChild">
                    <span class="inline-flex items-center gap-1">
                      <PlusIcon class="h-5 w-5 shrink-0" />
                      <span>Подкатегория</span>
                    </span>
                  </UiButton>
                </div>
                <div v-if="selectedCategory?.children?.length" class="flex flex-wrap gap-2">
                  <button
                    v-for="ch in selectedCategory.children"
                    :key="ch.id"
                    type="button"
                    class="px-3 py-2 rounded-lg border border-slate-200 hover:bg-slate-50 text-sm"
                    @click="selectedId = ch.id"
                  >
                    {{ ch.name }}
                    <span v-if="ch.product_count > 0" class="text-slate-400 ml-1">({{ ch.product_count }})</span>
                  </button>
                </div>
                <UiEmpty v-else message="Нет подкатегорий" class="py-8" />
              </div>

              <!-- Товары -->
              <div v-show="activeTab === 'products' && selectedId !== 'new'" class="space-y-4">
                <div class="flex items-center justify-between">
                  <span class="text-sm text-slate-600">Товары в категории</span>
                  <UiButton size="md" variant="secondary" @click="openAddProduct">
                    <span class="inline-flex items-center gap-1">
                      <PlusIcon class="h-5 w-5 shrink-0" />
                      <span>Товар</span>
                    </span>
                  </UiButton>
                </div>
                <UiTable
                  v-if="categoryProducts.length"
                  :columns="productTableColumns"
                  :data="categoryProducts"
                  row-key="id"
                >
                  <template #actions="{ row }">
                    <UiButton variant="ghost-primary" size="sm" class="mr-2" @click="openEditProduct(row)">Изменить</UiButton>
                    <UiButton variant="ghost-danger" size="sm" @click="delProduct(row)">Удалить</UiButton>
                  </template>
                </UiTable>
                <UiEmpty v-else-if="!productsLoading" message="Нет товаров" class="py-8" />
              </div>
            </div>
          </template>
        </div>
      </section>
    </div>

    <!-- Add product modal: existing or new -->
    <UiModal v-model="addProductModal" size="md">
      <template #title>
        <h2 class="text-xl font-bold">Добавить товар</h2>
      </template>
      <div class="space-y-4">
        <div class="flex gap-2 border-b border-slate-200">
          <button
            type="button"
            class="px-3 py-2 text-sm font-medium -mb-px border-b-2 transition-colors"
            :class="addProductMode === 'existing' ? 'border-slate-800 text-slate-800' : 'border-transparent text-slate-500 hover:text-slate-700'"
            @click="addProductMode = 'existing'"
          >
            Из каталога
          </button>
          <button
            type="button"
            class="px-3 py-2 text-sm font-medium -mb-px border-b-2 transition-colors"
            :class="addProductMode === 'new' ? 'border-slate-800 text-slate-800' : 'border-transparent text-slate-500 hover:text-slate-700'"
            @click="addProductMode = 'new'"
          >
            Новый товар
          </button>
        </div>

        <!-- Existing: search + list -->
        <div v-if="addProductMode === 'existing'">
          <UiInput v-model="productSearchQ" placeholder="Поиск по названию или артикулу..." @keyup.enter="searchExistingProducts" />
          <UiButton class="mt-2" size="sm" @click="searchExistingProducts">Найти</UiButton>
          <div class="mt-4 max-h-48 overflow-y-auto space-y-2">
            <button
              v-for="p in existingProducts"
              :key="p.id"
              type="button"
              class="w-full text-left px-3 py-2 rounded-lg border border-slate-200 hover:bg-slate-50 text-sm"
              @click="assignProductToCategory(p)"
            >
              {{ p.name }}
              <span v-if="p.sku" class="text-slate-400 ml-1">({{ p.sku }})</span>
            </button>
            <p v-if="existingProducts.length === 0 && existingProductsSearched" class="text-sm text-slate-500 py-2">Ничего не найдено</p>
          </div>
        </div>

        <!-- New: form -->
        <div v-else class="space-y-4 max-h-[50vh] overflow-y-auto">
          <div class="space-y-3">
            <UiInput v-model="newProductForm.name" label="Название *" />
            <UiInput v-model="newProductForm.sku" label="Артикул" />
            <UiInput v-model="newProductForm.description" label="Описание (Markdown)" tag="textarea" :rows="3" />
            <div>
              <label class="block text-sm font-medium text-slate-700 mb-1.5">Изображение</label>
              <div class="flex gap-2">
                <UiInput v-model="newProductForm.image_path" placeholder="URL или загрузите" class="flex-1" />
                <input type="file" accept="image/*" @change="uploadNewProductFile" class="rounded-lg border border-slate-300 px-3 py-2.5 text-sm" />
              </div>
            </div>
          </div>
          <div class="space-y-3">
            <h3 class="text-sm font-medium text-slate-600 border-b border-slate-200 pb-1">Характеристики</h3>
            <div class="grid grid-cols-2 gap-3">
              <UiInput v-model="newProductForm.size" label="Размер" />
              <UiInput v-model="newProductForm.mark" label="Марка" />
              <UiInput v-model="newProductForm.length" label="Длина" />
              <UiInput v-model="newProductForm.city" label="Город" />
            </div>
          </div>
          <div class="space-y-3">
            <h3 class="text-sm font-medium text-slate-600 border-b border-slate-200 pb-1">Цены и наличие</h3>
            <div class="grid grid-cols-2 gap-3">
              <UiInput v-model.number="newProductForm.stock" type="number" label="В наличии" />
              <UiInput v-model.number="newProductForm.price_1t" type="number" step="0.01" label="Цена за 1 т" />
              <UiInput v-model.number="newProductForm.price_5t" type="number" step="0.01" label="Цена за 5 т" />
              <UiInput v-model.number="newProductForm.price_10t" type="number" step="0.01" label="Цена за 10 т" />
            </div>
          </div>
          <div class="space-y-3">
            <UiInput v-model="newProductForm.source_url" label="Ссылка на источник" />
            <UiInput v-model="newProductForm.characteristics" tag="textarea" :rows="2" placeholder="Дн: 50; Вид: Зонт; Вес, кг: 0,09" label="Характеристики" />
          </div>
        </div>
      </div>
      <div class="mt-6 flex gap-2">
        <UiButton v-if="addProductMode === 'new'" @click="createNewProduct">Сохранить</UiButton>
        <UiButton variant="secondary" @click="addProductModal = false">Отмена</UiButton>
      </div>
    </UiModal>

    <!-- Edit product modal -->
    <UiModal v-model="editProductModal" size="md">
      <template #title>
        <h2 class="text-xl font-bold">Редактировать товар</h2>
      </template>
      <div class="space-y-4 max-h-[70vh] overflow-y-auto pr-2">
        <div class="space-y-3">
          <h3 class="text-sm font-medium text-slate-600 border-b border-slate-200 pb-1">Основные</h3>
          <UiInput v-model="editProductForm.name" label="Название *" />
          <UiInput v-model="editProductForm.sku" label="Артикул" />
          <UiSelect v-model="editProductForm.category_id" label="Категория">
            <option value="">—</option>
            <option v-for="c in flatCategories" :key="c.id" :value="c.id">{{ c.name }}</option>
          </UiSelect>
          <UiInput v-model="editProductForm.description" label="Описание (Markdown)" tag="textarea" :rows="4" />
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-1.5">Изображение</label>
            <div class="flex gap-2">
              <UiInput v-model="editProductForm.image_path" placeholder="URL или загрузите" class="flex-1" />
              <input type="file" accept="image/*" @change="uploadEditProductFile" class="rounded-lg border border-slate-300 px-3 py-2.5 text-sm" />
            </div>
          </div>
        </div>
        <div class="space-y-3">
          <h3 class="text-sm font-medium text-slate-600 border-b border-slate-200 pb-1">Характеристики</h3>
          <div class="grid grid-cols-2 gap-3">
            <UiInput v-model="editProductForm.size" label="Размер" />
            <UiInput v-model="editProductForm.mark" label="Марка" />
            <UiInput v-model="editProductForm.length" label="Длина" />
            <UiInput v-model="editProductForm.city" label="Город" />
          </div>
        </div>
        <div class="space-y-3">
          <h3 class="text-sm font-medium text-slate-600 border-b border-slate-200 pb-1">Цены и наличие</h3>
          <div class="grid grid-cols-2 gap-3">
            <UiInput v-model.number="editProductForm.stock" type="number" label="В наличии" />
            <UiInput v-model.number="editProductForm.price_1t" type="number" step="0.01" label="Цена за 1 т" />
            <UiInput v-model.number="editProductForm.price_5t" type="number" step="0.01" label="Цена за 5 т" />
            <UiInput v-model.number="editProductForm.price_10t" type="number" step="0.01" label="Цена за 10 т" />
          </div>
        </div>
        <div class="space-y-3">
          <h3 class="text-sm font-medium text-slate-600 border-b border-slate-200 pb-1">Ссылка</h3>
          <UiInput v-model="editProductForm.source_url" label="Ссылка на источник" />
        </div>
        <div class="space-y-3">
          <h3 class="text-sm font-medium text-slate-600 border-b border-slate-200 pb-1">Атрибуты</h3>
          <UiInput v-model="editProductForm.characteristics" tag="textarea" :rows="3" placeholder="Дн: 50; Вид: Зонт; Вес, кг: 0,09" label="Характеристики" />
        </div>
      </div>
      <div class="mt-6 flex gap-2">
        <UiButton @click="saveEditProduct">Сохранить</UiButton>
        <UiButton variant="secondary" @click="editProductModal = false">Отмена</UiButton>
      </div>
    </UiModal>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { PlusIcon, DocumentTextIcon, FolderIcon, CubeIcon } from '@heroicons/vue/24/outline'
import { UiButton, UiEmpty, UiInput, UiModal, UiSelect, UiTable } from '@ui'
import CategoryTreeNode from '../components/CategoryTreeNode.vue'
import client from '../api/client'

const categories = ref([])
const selectedId = ref(null)
const form = ref({ name: '', parent_id: '', sort_order: 0 })
const categoryProducts = ref([])
const productsLoading = ref(false)
const addProductModal = ref(false)
const addProductMode = ref('existing')
const productSearchQ = ref('')
const existingProducts = ref([])
const existingProductsSearched = ref(false)
const emptyNewProductForm = () => ({
  name: '',
  sku: '',
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
const newProductForm = ref(emptyNewProductForm())
const editProductModal = ref(false)
const emptyEditProductForm = () => ({
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
const editProductForm = ref(emptyEditProductForm())
const editingProduct = ref(null)
const activeTab = ref('description')

const productTableColumns = [
  { key: 'name', label: 'Название' },
  { key: 'sku', label: 'Артикул' },
]

const selectedCategory = computed(() => {
  if (!selectedId.value) return null
  return findInTree(categories.value, selectedId.value)
})

const flatCategories = computed(() => {
  const out = []
  function walk(list, prefix = '') {
    for (const c of list) {
      out.push({ ...c, name: prefix + c.name })
      if (c.children?.length) walk(c.children, prefix + '— ')
    }
  }
  walk(categories.value)
  return out
})

const breadcrumbs = computed(() => {
  const cat = selectedCategory.value
  if (!cat) return []
  const path = []
  function collect(c, list) {
    if (!c?.id) return
    list.unshift({ id: c.id, name: c.name })
    const pid = c.parent_id
    if (pid) {
      const parent = findInTree(categories.value, pid)
      if (parent) collect(parent, list)
    }
  }
  collect(cat, path)
  return path
})

watch(selectedId, (id) => {
  if (id && selectedCategory.value) {
    form.value = {
      name: selectedCategory.value.name,
      parent_id: selectedCategory.value.parent_id ?? '',
      sort_order: selectedCategory.value.sort_order ?? 0,
    }
  }
  if (id === 'new') activeTab.value = 'description'
  loadProducts()
})

function findInTree(list, id) {
  if (!list?.length) return null
  const numId = Number(id)
  for (const c of list) {
    if (Number(c.id) === numId) return c
    const found = findInTree(c.children, id)
    if (found) return found
  }
  return null
}

function isDescendantOfSelected(c) {
  const sel = selectedCategory.value
  if (!sel || sel.id === 'new') return false
  if (c.id === sel.id) return true
  return isInSubtree(sel.children, c.id)
}
function isInSubtree(nodes, id) {
  if (!nodes?.length) return false
  for (const n of nodes) {
    if (n.id === id) return true
    if (isInSubtree(n.children, id)) return true
  }
  return false
}

async function load() {
  const { data } = await client.get('/categories')
  categories.value = data ?? []
}

function openCreate() {
  form.value = { name: '', parent_id: '', sort_order: 0 }
  selectedId.value = 'new'
}
function openAddChild() {
  if (!selectedCategory.value) return
  form.value = { name: '', parent_id: String(selectedCategory.value.id), sort_order: 0 }
  selectedId.value = 'new'
}

async function save() {
  if (!form.value.name) return
  const payload = {
    name: form.value.name,
    parent_id: form.value.parent_id ? Number(form.value.parent_id) : null,
    sort_order: form.value.sort_order ?? 0,
  }
  if (selectedId.value === 'new') {
    const { data } = await client.post('/admin/categories', payload)
    await load()
    selectedId.value = data?.id
  } else {
    await client.put(`/admin/categories/${selectedId.value}`, payload)
    await load()
  }
}

async function del() {
  if (selectedId.value === 'new' || !selectedCategory.value) return
  if (!confirm('Удалить категорию?')) return
  await client.delete(`/admin/categories/${selectedId.value}`)
  selectedId.value = null
  await load()
}

async function loadProducts() {
  const id = selectedId.value
  if (id === 'new' || id === null || id === undefined) {
    categoryProducts.value = []
    return
  }
  productsLoading.value = true
  try {
    const { data } = await client.get('/products', { params: { category_id: id, per_page: 100 } })
    categoryProducts.value = data?.products ?? []
  } catch {
    categoryProducts.value = []
  } finally {
    productsLoading.value = false
  }
}

function openAddProduct() {
  if (!selectedCategory.value || selectedId.value === 'new') return
  addProductModal.value = true
  addProductMode.value = 'existing'
  productSearchQ.value = ''
  existingProducts.value = []
  existingProductsSearched.value = false
  newProductForm.value = emptyNewProductForm()
  searchExistingProducts()
}

async function searchExistingProducts() {
  existingProductsSearched.value = true
  try {
    const params = { per_page: 20 }
    if (productSearchQ.value) params.q = productSearchQ.value
    const { data } = await client.get('/products', { params })
    existingProducts.value = data?.products ?? []
  } catch {
    existingProducts.value = []
  }
}

async function assignProductToCategory(p) {
  if (!selectedCategory.value) return
  try {
    const attrs = Array.isArray(p.attributes) ? p.attributes : []
    await client.put(`/admin/products/${p.id}`, {
      name: p.name,
      sku: p.sku || '',
      category_id: selectedCategory.value.id,
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
      attributes: attrs,
    })
    addProductModal.value = false
    await loadProducts()
    await load()
  } catch (e) {
    console.error('Ошибка при добавлении товара в категорию:', e)
  }
}

async function uploadNewProductFile(e) {
  const file = e.target?.files?.[0]
  if (!file) return
  const fd = new FormData()
  fd.append('file', file)
  const { data } = await client.post('/admin/upload', fd, { headers: { 'Content-Type': 'multipart/form-data' } })
  newProductForm.value.image_path = data.url
}

function buildNewProductPayload() {
  const f = newProductForm.value
  const attrs = []
  if (f.characteristics?.trim()) {
    attrs.push({ key: 'characteristics', value: f.characteristics.trim() })
  }
  return {
    name: f.name,
    sku: f.sku || '',
    category_id: selectedCategory.value.id,
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
async function createNewProduct() {
  if (!newProductForm.value.name || !selectedCategory.value) return
  try {
    await client.post('/admin/products', buildNewProductPayload())
    addProductModal.value = false
    await loadProducts()
    await load()
  } catch (e) {
    console.error('Ошибка при создании товара:', e)
  }
}

function openEditProduct(p) {
  editingProduct.value = p
  const chars = p.attributes?.find((a) => a.key === 'characteristics')
  editProductForm.value = {
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
  editProductModal.value = true
}

async function uploadEditProductFile(e) {
  const file = e.target?.files?.[0]
  if (!file) return
  const fd = new FormData()
  fd.append('file', file)
  const { data } = await client.post('/admin/upload', fd, { headers: { 'Content-Type': 'multipart/form-data' } })
  editProductForm.value.image_path = data.url
}

function buildEditProductPayload() {
  const f = editProductForm.value
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
async function saveEditProduct() {
  if (!editingProduct.value || !editProductForm.value.name) return
  try {
    await client.put(`/admin/products/${editingProduct.value.id}`, buildEditProductPayload())
    editProductModal.value = false
    await loadProducts()
    await load()
  } catch (e) {
    console.error('Ошибка при сохранении товара:', e)
  }
}

async function delProduct(p) {
  if (!confirm('Удалить товар?')) return
  try {
    await client.delete(`/admin/products/${p.id}`)
    await loadProducts()
    await load()
  } catch (e) {
    console.error('Ошибка при удалении товара:', e)
  }
}

onMounted(load)
</script>
