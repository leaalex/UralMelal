<template>
  <div>
    <UiPageHeader title="Категории">
      <template #actions>
        <UiButton @click="openCreate">+ Категория</UiButton>
      </template>
    </UiPageHeader>

    <UiCard v-if="categories.length">
      <div v-for="c in categories" :key="c.id" class="flex items-center justify-between py-4 border-b border-slate-100 last:border-0">
        <span class="text-slate-800">{{ c.name }}</span>
        <div>
          <UiButton variant="ghost-primary" size="sm" class="mr-2" @click="openEdit(c)">Изменить</UiButton>
          <UiButton variant="ghost-danger" size="sm" @click="del(c)">Удалить</UiButton>
        </div>
      </div>
    </UiCard>
    <UiEmpty v-else message="Нет категорий" />

    <UiModal v-model="showModal" size="sm">
      <template #title>
        <h2 class="text-xl font-bold">{{ editing?.id ? 'Редактировать' : 'Новая категория' }}</h2>
      </template>
      <div class="space-y-3">
        <UiInput v-model="form.name" label="Название *" />
        <UiSelect v-model="form.parent_id" label="Родитель">
          <option value="">Нет</option>
          <option v-for="c in categories" :key="c.id" :value="c.id" :disabled="editing && c.id === editing.id">{{ c.name }}</option>
        </UiSelect>
      </div>
      <div class="mt-6 flex gap-2">
        <UiButton @click="save">Сохранить</UiButton>
        <UiButton variant="secondary" @click="showModal = false">Отмена</UiButton>
      </div>
    </UiModal>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { UiButton, UiCard, UiEmpty, UiInput, UiModal, UiPageHeader, UiSelect } from '@ui'
import client from '../api/client'

const categories = ref([])
const showModal = ref(false)
const editing = ref(null)
const form = ref({ name: '', parent_id: '' })

async function load() {
  const { data } = await client.get('/categories')
  categories.value = flatten(data)
}
function flatten(list, out = []) {
  for (const c of list) {
    out.push(c)
    if (c.children?.length) flatten(c.children, out)
  }
  return out
}
function openCreate() {
  editing.value = null
  form.value = { name: '', parent_id: '' }
  showModal.value = true
}
function openEdit(c) {
  editing.value = c
  form.value = { name: c.name, parent_id: c.parent_id ?? '' }
  showModal.value = true
}
async function save() {
  if (!form.value.name) return
  const payload = { name: form.value.name, parent_id: form.value.parent_id || null }
  if (editing.value) {
    await client.put(`/admin/categories/${editing.value.id}`, payload)
  } else {
    await client.post('/admin/categories', payload)
  }
  showModal.value = false
  load()
}
async function del(c) {
  if (!confirm('Удалить категорию?')) return
  await client.delete(`/admin/categories/${c.id}`)
  load()
}

onMounted(load)
</script>
