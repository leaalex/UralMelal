<template>
  <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between mb-6">
    <h1 class="text-2xl font-bold tracking-tight text-slate-800">{{ title }}</h1>
    <div v-if="$slots.actions" class="flex flex-wrap gap-2">
      <slot name="actions" />
    </div>
  </div>
</template>

<script setup>
defineProps({ title: { type: String, default: '' } })
</script>
