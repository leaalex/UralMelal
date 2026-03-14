<template>
  <aside
    class="fixed inset-y-0 left-0 z-40 w-72 max-w-[85vw] bg-white border-r border-slate-200 shadow-lg transform transition-transform duration-200 ease-out lg:relative lg:z-0 lg:left-auto lg:inset-auto lg:w-64 lg:max-w-none lg:shrink-0 lg:shadow-none lg:translate-x-0"
    :class="open ? 'translate-x-0' : 'max-lg:-translate-x-full'"
  >
    <div class="flex items-center justify-between p-4 border-b border-slate-200 lg:hidden">
      <span class="font-semibold text-slate-800">Категории</span>
      <button
        type="button"
        class="p-2 rounded-xl text-slate-500 hover:bg-slate-100 hover:text-slate-600 transition-colors"
        aria-label="Закрыть"
        @click="emit('update:open', false)"
      >
        <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>
    <nav class="p-4 overflow-y-auto overflow-x-hidden max-h-[calc(100vh-4rem)] lg:max-h-none min-w-0">
      <button
        type="button"
        class="w-full text-left px-3 py-2.5 rounded-xl text-sm font-medium transition-colors"
        :class="!modelValue ? 'bg-slate-200 text-slate-900 font-semibold' : 'text-slate-700 hover:bg-slate-100'"
        @click="emit('update:modelValue', null)"
      >
        Все категории
      </button>
      <CategoryTree
        v-for="c in rootCategories"
        :key="c.id"
        :category="c"
        :selected-id="modelValue"
        @select="emit('update:modelValue', $event)"
      />
    </nav>
  </aside>
  <div
    v-if="open"
    class="fixed inset-0 z-30 bg-black/30 lg:hidden"
    aria-hidden="true"
    @click="emit('update:open', false)"
  />
</template>

<script setup>
import { computed } from 'vue'
import CategoryTree from './CategoryTree.vue'

const props = defineProps({
  categories: { type: Array, default: () => [] },
  modelValue: { type: [Number, String, null], default: null },
  open: { type: Boolean, default: false },
})

const emit = defineEmits(['update:modelValue', 'update:open'])

const rootCategories = computed(() => {
  return (props.categories ?? []).filter((c) => !c.parent_id)
})
</script>
