<template>
  <div>
    <section class="relative w-full h-screen min-h-[400px] overflow-hidden">
      <video
        autoplay
        muted
        loop
        playsinline
        class="absolute inset-0 w-full h-full object-cover"
      >
        <source src="/video/hero.webm" type="video/webm" />
        <source src="/video/hero.mp4" type="video/mp4" />
      </video>
      <div class="absolute inset-0 bg-white/50" aria-hidden="true" />
      <div class="absolute inset-0 flex flex-col items-center justify-center container mx-auto px-4 text-center">
        <router-link to="/" class="mb-4">
          <img src="/logo-full.png" alt="Урал Металл" class="h-16 md:h-20 lg:h-24 w-auto object-contain" />
        </router-link>
        <p class="text-xl md:text-2xl text-slate-700 mb-8">Надёжный поставщик металлопродукции</p>
        <div class="flex flex-wrap items-center justify-center gap-4">
          <UiButton to="/catalog" variant="secondary" size="xl" class="min-w-[220px] shrink-0 shadow-lg border-slate-300 bg-white/90 hover:bg-white text-slate-800">
            Перейти в каталог
          </UiButton>
          <UiButton variant="secondary" size="xl" class="min-w-[220px] shrink-0 shadow-lg border-slate-300 bg-white/90 hover:bg-white text-slate-800" @click="openContactModal?.()">
            Связаться со мной
          </UiButton>
        </div>
      </div>
    </section>
    <section v-if="homeBlocks.length" class="py-12 bg-white">
      <div class="container mx-auto px-4">
        <div v-for="block in homeBlocks" :key="block.id" class="mb-12 text-slate-800 prose prose-slate max-w-none">
          <h2 v-if="block.title" class="text-2xl font-bold mb-4 text-slate-900">{{ block.title }}</h2>
          <MarkdownContent :content="block.body" />
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, inject, onMounted } from 'vue'
const openContactModal = inject('openContactModal')
import client from '../api/client'
import { MarkdownContent, UiButton } from '@ui'

const homeBlocks = ref([])

onMounted(async () => {
  const { data } = await client.get('/content/home')
  homeBlocks.value = data || []
})
</script>
