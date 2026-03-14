<template>
  <div
    class="rounded-xl border border-slate-200 bg-white shadow-sm ring-1 ring-slate-200/50 transition-shadow duration-150 hover:shadow-md"
    :class="paddingClass"
  >
    <slot />
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  padding: { type: String, default: 'md' },
})

const paddingClass = computed(() => {
  const map = { none: 'p-0', sm: 'p-4', md: 'p-6', lg: 'p-8' }
  return map[props.padding] || map.md
})
</script>
