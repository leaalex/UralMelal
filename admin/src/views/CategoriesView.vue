<template>
  <div class="flex flex-col h-full">
    <UiPageHeader title="Категории">
      <template #actions>
        <UiButton @click="openCreate">+ Категория</UiButton>
      </template>
    </UiPageHeader>

    <div class="flex flex-1 gap-4 min-h-0 mt-4">
      <!-- Left: tree -->
      <aside class="w-72 shrink-0 overflow-y-auto border border-slate-200 rounded-xl bg-white">
        <div class="p-2">
          <CategoryTreeNode
            v-for="c in categories"
            :key="c.id"
            :node="c"
            :selected-id="selectedId"
            :depth="0"
            @select="selectedId = $event"
          />
        </div>
        <UiEmpty v-if="!categories.length" message="Нет категорий" class="py-8" />
      </aside>

      <!-- Right: form / details -->
      <section class="flex-1 min-w-0 overflow-y-auto border border-slate-200 rounded-xl bg-white p-6">
        <div v-if="selectedId === null" class="flex flex-col items-center justify-center py-16 text-slate-500">
          <p>Выберите категорию или создайте новую</p>
        </div>

        <div v-else class="space-y-4">
          <h3 v-if="selectedId === 'new'" class="text-lg font-semibold text-slate-800">Новая категория</h3>
          <!-- Breadcrumbs -->
          <nav v-else-if="breadcrumbs.length" class="flex items-center gap-2 text-sm text-slate-500">
            <span
              v-for="(crumb, i) in breadcrumbs"
              :key="crumb.id"
              class="flex items-center gap-2"
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

          <div class="flex gap-2 pt-2">
            <UiButton @click="save">Сохранить</UiButton>
            <UiButton v-if="selectedId !== 'new'" variant="secondary" @click="openAddChild">+ Подкатегория</UiButton>
            <UiButton v-if="selectedId !== 'new'" variant="ghost-danger" @click="del">Удалить</UiButton>
          </div>

          <!-- Children list -->
          <div v-if="selectedCategory?.children?.length" class="pt-4 border-t border-slate-200">
            <h4 class="text-sm font-medium text-slate-600 mb-2">Подкатегории</h4>
            <div class="flex flex-wrap gap-2">
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
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { UiButton, UiEmpty, UiInput, UiPageHeader, UiSelect } from '@ui'
import CategoryTreeNode from '../components/CategoryTreeNode.vue'
import client from '../api/client'

const categories = ref([])
const selectedId = ref(null)
const form = ref({ name: '', parent_id: '', sort_order: 0 })

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

onMounted(load)
</script>
