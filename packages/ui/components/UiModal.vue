<template>
  <Teleport to="body">
    <Transition name="modal">
      <div v-if="modelValue" class="fixed inset-0 flex items-center justify-center z-[100] p-4" @click.self="close">
        <div class="absolute inset-0 bg-black/50 transition-opacity" />
        <div
          class="relative bg-white rounded-xl shadow-lg max-w-2xl w-full max-h-[90vh] overflow-y-auto ring-1 ring-slate-200/50 transition-all duration-200"
          :class="[sizeClass, paddingClass]"
        >
          <button
            v-if="showClose"
            type="button"
            class="absolute top-4 right-4 rounded-xl p-2 text-slate-500 hover:bg-slate-100 hover:text-slate-600 transition-colors duration-150"
            aria-label="Закрыть"
            @click="close"
          >
            <XMarkIcon class="h-5 w-5" />
          </button>
          <div v-if="$slots.title" class="mb-3 sm:mb-4 pr-10">
            <slot name="title" />
          </div>
          <slot />
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { computed } from 'vue'
import { XMarkIcon } from '@heroicons/vue/24/outline'

const props = defineProps({
  modelValue: Boolean,
  size: { type: String, default: 'md' },
  showClose: { type: Boolean, default: true },
})

const emit = defineEmits(['update:modelValue'])

const sizeClass = computed(() => {
  const map = { sm: 'max-w-sm', md: 'max-w-2xl', lg: 'max-w-4xl' }
  return map[props.size] || map.md
})

const paddingClass = computed(() => 'p-4 sm:p-6')

function close() {
  emit('update:modelValue', false)
}
</script>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.2s ease;
}
.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
</style>
