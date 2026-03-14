<template>
  <div class="flex flex-col items-center justify-center py-12 text-center">
    <div class="mb-4 rounded-full bg-slate-100 p-4">
      <InboxIcon class="h-10 w-10 text-slate-400" />
    </div>
    <p class="text-slate-500">{{ message }}</p>
    <div v-if="$slots.action" class="mt-4">
      <slot name="action" />
    </div>
  </div>
</template>

<script setup>
import { InboxIcon } from '@heroicons/vue/24/outline'

defineProps({ message: { type: String, default: 'Нет данных' } })
</script>
