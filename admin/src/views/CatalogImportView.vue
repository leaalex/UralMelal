<template>
  <div>
    <UiPageHeader title="Импорт каталога" />
    <UiCard class="max-w-lg">
      <UiSelect v-model="format" label="Формат">
        <option value="json">JSON</option>
        <option value="csv">CSV</option>
        <option value="mc">MC (металлсервис)</option>
      </UiSelect>
      <div class="mt-4">
        <label class="block text-sm font-medium text-slate-700 mb-1.5">Файл</label>
        <input
          type="file"
          @change="onFile"
          :accept="format === 'json' ? '.json' : '.csv'"
          class="w-full rounded-lg border border-slate-300 bg-white px-3 py-2.5 text-sm file:mr-4 file:rounded-lg file:border-0 file:bg-slate-100 file:px-4 file:py-2 file:text-sm file:font-medium"
        />
      </div>
      <p v-if="result" class="mt-4 text-emerald-600 font-medium">
        Создано: {{ result.created ?? result }}
        <span v-if="result.skipped != null">, пропущено: {{ result.skipped }}</span>
      </p>
      <p v-if="error" class="mt-4 text-red-500">{{ error }}</p>
      <UiButton class="mt-4" :disabled="!file" @click="upload">Загрузить</UiButton>
    </UiCard>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { UiButton, UiCard, UiPageHeader, UiSelect } from '@ui'
import client from '../api/client'

const format = ref('json')
const file = ref(null)
const result = ref(null)
const error = ref('')

function onFile(e) {
  file.value = e.target.files?.[0]
  result.value = null
  error.value = ''
}
async function upload() {
  if (!file.value) return
  result.value = null
  error.value = ''
  try {
    const fd = new FormData()
    fd.append('file', file.value)
    const { data } = await client.post(`/admin/products/import?format=${format.value}`, fd)
    result.value = data
    file.value = null
  } catch (e) {
    error.value = e.response?.data?.error || 'Ошибка'
  }
}
</script>
