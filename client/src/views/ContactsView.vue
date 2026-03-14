<template>
  <div class="container mx-auto px-4 py-8">
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-8">
      <h1 class="text-3xl font-bold text-slate-900">Контакты</h1>
      <UiButton size="lg" class="shrink-0" @click="openContactModal?.()">
        Связаться со мной
      </UiButton>
    </div>
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-12">
      <div>
        <div v-for="block in contactBlocks" :key="block.id" class="mb-6 text-slate-800 prose prose-slate max-w-none">
          <h2 v-if="block.title" class="text-xl font-semibold mb-2 text-slate-900">{{ block.title }}</h2>
          <MarkdownContent :content="block.body" />
        </div>
      </div>
      <div>
        <div id="map" class="w-full h-80 rounded-lg bg-slate-200"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, inject, onMounted } from 'vue'
import { UiButton } from '@ui'
const openContactModal = inject('openContactModal')
import client from '../api/client'
import { MarkdownContent } from '@ui'

const contactBlocks = ref([])

onMounted(async () => {
  const { data } = await client.get('/content/contacts')
  contactBlocks.value = data || []
  if (typeof ymaps !== 'undefined') {
    ymaps.ready(() => {
      new ymaps.Map('map', {
        center: [56.8389, 60.6057],
        zoom: 12,
      })
    })
  }
})
</script>
