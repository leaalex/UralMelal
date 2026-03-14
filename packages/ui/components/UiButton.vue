<template>
  <component
    :is="to ? 'router-link' : 'button'"
    :to="to"
    :type="to ? undefined : (props.type || 'button')"
    :disabled="disabled || loading"
    class="inline-flex items-center justify-center gap-2 rounded-lg font-medium transition-all duration-150 disabled:opacity-50 disabled:cursor-not-allowed"
    :class="[variantClasses, sizeClasses]"
    @click="!to && !loading && $emit('click')"
  >
    <span v-if="loading" class="inline-block h-4 w-4 animate-spin rounded-full border-2 border-current border-t-transparent" />
    <span v-else-if="$slots.icon" class="inline-flex shrink-0"><slot name="icon" /></span>
    <slot />
  </component>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  variant: { type: String, default: 'primary' },
  size: { type: String, default: 'md' },
  type: { type: String, default: 'button' },
  disabled: Boolean,
  loading: Boolean,
  to: { type: [String, Object], default: null },
})

defineEmits(['click'])

const variantClasses = computed(() => {
  const map = {
    primary: 'bg-slate-800 text-white hover:bg-slate-700',
    secondary: 'border border-slate-300 bg-white hover:bg-slate-50 text-slate-700',
    danger: 'bg-red-600 text-white hover:bg-red-500',
    success: 'bg-green-600 text-white hover:bg-green-500',
    outline: 'border-2 border-white/50 bg-transparent hover:bg-white/10 text-white',
    ghost: 'text-slate-500 hover:bg-slate-100 hover:text-slate-600',
    'ghost-primary': 'text-blue-600 hover:bg-blue-50 hover:text-blue-700',
    'ghost-danger': 'text-red-600 hover:bg-red-50 hover:text-red-700',
  }
  return map[props.variant] || map.primary
})

const sizeClasses = computed(() => {
  const map = {
    sm: 'px-3 py-1.5 text-sm',
    md: 'px-4 py-2 text-sm',
    lg: 'px-6 py-3 text-base',
    xl: 'px-8 py-4 text-base',
  }
  return map[props.size] || map.md
})
</script>
