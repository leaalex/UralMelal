<template>
  <UiModal v-model="open" size="lg">
    <template #title>
      <h3 class="text-lg font-semibold text-slate-900">Оставить заявку</h3>
    </template>
    <ContactForm hide-title />
  </UiModal>
</template>

<script setup>
import { computed } from 'vue'
import { UiModal } from '@ui'
import ContactForm from './ContactForm.vue'

const props = defineProps({
  modelValue: { type: Boolean, default: false },
})

const emit = defineEmits(['update:modelValue'])

const open = computed({
  get: () => props.modelValue,
  set: (v) => emit('update:modelValue', v),
})
</script>
