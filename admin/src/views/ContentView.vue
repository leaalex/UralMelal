<template>
  <div>
    <UiPageHeader title="Контент" />
    <div class="space-y-6">
      <UiCard v-for="block in blocks" :key="block.id" :class="block.hidden ? 'opacity-75' : ''">
        <div class="flex items-center justify-between mb-2">
          <h2 class="text-lg font-semibold text-slate-800">{{ block.title }} ({{ block.page }})</h2>
          <label class="flex items-center gap-2 cursor-pointer text-sm text-slate-600">
            <input v-model="block.hidden" type="checkbox" class="rounded border-slate-300" @change="saveBlock(block)" />
            Скрыть на сайте
          </label>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <UiInput v-model="block.body" label="Markdown" tag="textarea" :rows="8" class="[&_textarea]:font-mono [&_textarea]:text-sm" />
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-1.5">Предпросмотр</label>
            <div class="rounded-lg border border-slate-200 p-4 min-h-[200px] bg-slate-50 prose prose-sm max-w-none">
              <MarkdownContent :content="block.body" />
            </div>
          </div>
        </div>
        <UiButton class="mt-4" @click="saveBlock(block)">Сохранить</UiButton>
      </UiCard>
    </div>
    <UiEmpty v-if="blocks.length === 0" message="Нет контент-блоков" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { MarkdownContent, UiButton, UiCard, UiEmpty, UiInput, UiPageHeader } from '@ui'
import client from '../api/client'

const blocks = ref([])

async function load() {
  const [home, contacts] = await Promise.all([
    client.get('/admin/content?page=home'),
    client.get('/admin/content?page=contacts'),
  ])
  blocks.value = [...(home.data || []), ...(contacts.data || [])]
}
async function saveBlock(block) {
  await client.put(`/admin/content/${block.id}`, {
    title: block.title,
    body: block.body,
    sort_order: block.sort_order,
    hidden: block.hidden,
  })
}

onMounted(load)
</script>
