<template>
  <div>
    <div class="flex flex-wrap items-center gap-2 mb-4">
      <UiSelect v-model="statusFilter" @change="load">
        <option value="">Все</option>
        <option value="new">Новые</option>
        <option value="processed">Обработанные</option>
      </UiSelect>
      <UiButton variant="secondary" @click="load">Обновить</UiButton>
    </div>
    <UiTable
      v-if="leads.length"
      :columns="tableColumns"
      :data="leads"
      row-key="id"
    >
      <template #cell-status="{ value }">
        <UiBadge :variant="value === 'new' ? 'new' : 'processed'" :label="value === 'new' ? 'Новая' : 'Обработана'" />
      </template>
      <template #actions="{ row }">
        <UiButton v-if="row.status === 'new'" size="sm" variant="secondary" @click="markProcessed(row)">Обработать</UiButton>
      </template>
    </UiTable>
    <UiEmpty v-else message="Нет заявок" />

    <div class="mt-4">
      <UiPagination v-model:page="page" :total="total" :per-page="perPage" />
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { UiBadge, UiButton, UiEmpty, UiPagination, UiSelect, UiTable } from '@ui'
import client from '../api/client'

const leads = ref([])
const statusFilter = ref('')
const page = ref(1)
const total = ref(0)
const perPage = 20

const tableColumns = [
  { key: 'name', label: 'Имя' },
  { key: 'phone', label: 'Телефон' },
  { key: 'subject', label: 'Тема' },
  { key: 'created_at', label: 'Дата' },
  { key: 'status', label: 'Статус' },
]

function formatDate(d) {
  if (!d) return '-'
  return new Date(d).toLocaleString('ru')
}

async function load() {
  const params = { page: page.value, per_page: perPage }
  if (statusFilter.value) params.status = statusFilter.value
  const { data } = await client.get('/admin/leads', { params })
  leads.value = data.leads.map((l) => ({ ...l, subject: l.subject || '-', created_at: formatDate(l.created_at) }))
  total.value = data.total
}

async function markProcessed(l) {
  await client.patch(`/admin/leads/${l.id}`, { status: 'processed' })
  load()
}

watch(page, load)

onMounted(load)
</script>
