<template>
  <span
    class="inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium transition-colors duration-150"
    :class="variantClasses"
  >
    <slot>{{ label }}</slot>
  </span>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  variant: { type: String, default: 'default' },
  label: String,
})

const variantClasses = computed(() => {
  const map = {
    default: 'bg-slate-100 text-slate-700',
    new: 'bg-emerald-100 text-emerald-800',
    processed: 'bg-slate-200 text-slate-600',
    success: 'bg-emerald-100 text-emerald-800',
    warning: 'bg-amber-100 text-amber-800',
    danger: 'bg-red-100 text-red-800',
  }
  return map[props.variant] || map.default
})
</script>
