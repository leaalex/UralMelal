<template>
  <div class="rounded-xl border border-slate-200 bg-white p-6 shadow-sm ring-1 ring-slate-200/50 transition-all duration-150 hover:shadow-md hover:ring-slate-300/50">
    <p class="text-sm font-medium text-slate-500">{{ label }}</p>
    <p class="mt-2 text-3xl font-bold tracking-tight text-slate-800">{{ value }}</p>
    <div v-if="$slots.icon" class="mt-2">
      <slot name="icon" />
    </div>
  </div>
</template>

<script setup>
defineProps({
  label: { type: String, default: '' },
  value: { type: [String, Number], default: 0 },
})
</script>
