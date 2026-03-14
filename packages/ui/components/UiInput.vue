<template>
  <div>
    <label v-if="label" class="block text-sm font-medium text-slate-700 mb-1.5">{{ label }}</label>
    <component
      :is="tag"
      :value="modelValue"
      v-bind="$attrs"
      class="w-full rounded-lg border border-slate-300 bg-white px-3 py-2.5 text-slate-800 shadow-sm transition-all duration-150 focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 focus:outline-none disabled:opacity-50"
      :class="{ 'border-red-500': error }"
      @input="emitValue"
    />
    <p v-if="error" class="mt-1 text-sm text-red-500">{{ error }}</p>
  </div>
</template>

<script setup>
defineOptions({ inheritAttrs: false })

const props = defineProps({
  modelValue: [String, Number],
  label: String,
  error: String,
  tag: { type: String, default: 'input' },
})

const emit = defineEmits(['update:modelValue'])

function emitValue(e) {
  const v = props.tag === 'textarea' ? e.target.value : (e.target.type === 'number' ? Number(e.target.value) : e.target.value)
  emit('update:modelValue', v)
}
</script>
