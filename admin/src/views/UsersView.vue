<template>
  <div>
    <div class="flex justify-end mb-4">
      <UiButton @click="openCreate">+ Пользователь</UiButton>
    </div>
    <UiTable
      v-if="users.length"
      :columns="tableColumns"
      :data="users"
      row-key="id"
    >
      <template #actions="{ row }">
        <UiButton variant="ghost-primary" size="sm" class="mr-2" @click="openEdit(row)">Изменить</UiButton>
        <UiButton variant="ghost-danger" size="sm" @click="del(row)">Удалить</UiButton>
      </template>
    </UiTable>
    <UiEmpty v-else message="Нет пользователей" />

    <UiModal v-model="showModal" size="sm">
      <template #title>
        <h2 class="text-xl font-bold">{{ editing?.id ? 'Редактировать' : 'Новый пользователь' }}</h2>
      </template>
      <div class="space-y-3">
        <UiInput v-model="form.username" label="Логин *" :readonly="!!editing?.id" />
        <UiInput v-model="form.password" type="password" :label="editing?.id ? 'Пароль (оставить пустым)' : 'Пароль *'" :placeholder="editing?.id ? 'Не менять' : undefined" />
        <UiSelect v-model="form.role" label="Роль">
          <option value="admin">Администратор</option>
          <option value="manager">Менеджер</option>
          <option value="editor">Редактор</option>
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
import { UiButton, UiEmpty, UiInput, UiModal, UiSelect, UiTable } from '@ui'
import client from '../api/client'

const users = ref([])
const showModal = ref(false)
const editing = ref(null)
const form = ref({ username: '', password: '', role: 'editor' })

const tableColumns = [
  { key: 'username', label: 'Логин' },
  { key: 'role', label: 'Роль' },
]

async function load() {
  const { data } = await client.get('/admin/users')
  users.value = data
}
function openCreate() {
  editing.value = null
  form.value = { username: '', password: '', role: 'editor' }
  showModal.value = true
}
function openEdit(u) {
  editing.value = u
  form.value = { username: u.username, password: '', role: u.role }
  showModal.value = true
}
async function save() {
  if (!form.value.username) return
  const payload = { username: form.value.username, role: form.value.role }
  if (form.value.password) payload.password = form.value.password
  if (editing.value) {
    await client.put(`/admin/users/${editing.value.id}`, payload)
  } else {
    if (!form.value.password) return
    await client.post('/admin/users', { ...payload, password: form.value.password })
  }
  showModal.value = false
  load()
}
async function del(u) {
  if (!confirm('Удалить пользователя?')) return
  await client.delete(`/admin/users/${u.id}`)
  load()
}

onMounted(load)
</script>
