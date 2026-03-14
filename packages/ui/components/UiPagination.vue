<template>
  <div class="flex items-center gap-1">
    <button
      :disabled="page <= 1"
      class="rounded-xl px-3 py-2 text-sm font-medium transition-all duration-150 disabled:cursor-not-allowed disabled:opacity-50 border border-slate-300 bg-white hover:bg-slate-50 text-slate-700"
      @click="$emit('update:page', page - 1)"
    >
      Назад
    </button>
    <div class="flex items-center gap-1">
      <template v-for="(item, i) in pageItems" :key="i">
        <button
          v-if="item === '...'"
          disabled
          class="px-2 py-2 text-slate-400 cursor-default"
        >
          …
        </button>
        <button
          v-else
          type="button"
          class="min-w-[2.25rem] rounded-xl px-3 py-2 text-sm font-medium transition-all duration-150 border"
          :class="item === page ? 'border-slate-800 bg-slate-800 text-white' : 'border-slate-300 bg-white hover:bg-slate-50 text-slate-700'"
          @click="$emit('update:page', item)"
        >
          {{ item }}
        </button>
      </template>
    </div>
    <button
      :disabled="page >= totalPages"
      class="rounded-xl px-3 py-2 text-sm font-medium transition-all duration-150 disabled:cursor-not-allowed disabled:opacity-50 border border-slate-300 bg-white hover:bg-slate-50 text-slate-700"
      @click="$emit('update:page', page + 1)"
    >
      Вперёд
    </button>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  page: { type: Number, required: true },
  total: { type: Number, default: 0 },
  perPage: { type: Number, default: 20 },
})

defineEmits(['update:page'])

const totalPages = computed(() => Math.ceil(props.total / props.perPage) || 1)

const pageItems = computed(() => {
  const total = totalPages.value
  const current = props.page
  if (total <= 7) {
    return Array.from({ length: total }, (_, i) => i + 1)
  }
  const items = []
  if (current <= 4) {
    for (let i = 1; i <= 5; i++) items.push(i)
    items.push('...')
    items.push(total)
  } else if (current >= total - 3) {
    items.push(1)
    items.push('...')
    for (let i = total - 4; i <= total; i++) items.push(i)
  } else {
    items.push(1)
    items.push('...')
    for (let i = current - 1; i <= current + 1; i++) items.push(i)
    items.push('...')
    items.push(total)
  }
  return items
})
</script>
