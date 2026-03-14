<template>
  <div>
    <label v-if="label" class="block text-sm font-medium text-slate-700 mb-1.5">{{ label }}</label>
    <select
      :value="modelValue"
      class="w-full rounded-lg border border-slate-300 bg-white px-3 py-2.5 text-slate-800 shadow-sm transition-all duration-150 focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 focus:outline-none"
      :class="{ 'border-red-500': error }"
      v-bind="$attrs"
      @change="$emit('update:modelValue', $event.target.value)"
    >
      <slot />
    </select>
    <p v-if="error" class="mt-1 text-sm text-red-500">{{ error }}</p>
  </div>
</template>

<script setup>
defineProps({
  modelValue: [String, Number, null],
  label: String,
  error: String,
})

defineEmits(['update:modelValue'])
</script>
