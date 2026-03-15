<template>
  <router-link
    to="/"
    class="flex items-center shrink-0 hover:opacity-90 no-underline"
    :class="compact ? 'justify-center w-full' : 'justify-start'"
  >
    <img
      :src="compact ? '/logo.png' : '/logo-full.png'"
      alt="Урал Металл"
      class="shrink-0 object-contain"
      :class="imgSizeClass"
    />
  </router-link>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  /** Узкое меню: logo.png (иконка) */
  compact: { type: Boolean, default: false },
  /** Размер: sm, md, lg */
  size: { type: String, default: 'md' },
})

const imgSizeClass = computed(() => ({
  'h-8 w-auto max-h-8': props.size === 'sm' || props.compact,
  'h-9 w-auto max-h-9': props.size === 'md' && !props.compact,
  'h-10 w-auto max-h-10': props.size === 'lg',
}))
</script>
