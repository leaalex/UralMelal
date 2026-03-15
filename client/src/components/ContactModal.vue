<template>
  <UiModal v-model="open" size="lg">
    <template #title>
      <div class="text-center py-0 sm:py-0">
        <h3 class="text-lg sm:text-xl font-bold text-slate-900">Связаться с нами</h3>
        <p class="mt-0.5 sm:mt-1 text-xs sm:text-sm text-slate-500">Ответим в ближайшее время</p>
      </div>
    </template>

    <div class="space-y-3 sm:space-y-6 -mx-1 sm:mx-0">
      <!-- Часть 1: Телефон и почта -->
      <div class="p-0">
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-2 sm:gap-4">
          <div class="flex items-center gap-2 sm:gap-4 p-2.5 sm:p-4 rounded-lg sm:rounded-xl bg-slate-50 hover:bg-slate-100">
            <a :href="`tel:${telHref}`" class="group flex items-center gap-2 sm:gap-3 min-w-0 flex-1">
              <div class="flex-shrink-0 w-9 h-9 sm:w-10 sm:h-10 rounded-lg bg-slate-200 flex items-center justify-center">
                <PhoneIcon class="h-5 w-5 sm:h-5 sm:w-5 text-slate-600" />
              </div>
              <div class="min-w-0">
                <span class="block text-[10px] sm:text-xs font-medium text-slate-500">Позвонить</span>
                <span class="block text-sm sm:text-base font-semibold text-slate-800 truncate">{{ contactPhone }}</span>
              </div>
            </a>
            <button
              type="button"
              class="flex-shrink-0 w-6 h-6 rounded flex items-center justify-center text-slate-400 hover:bg-slate-200 hover:text-slate-600"
              aria-label="Скопировать телефон"
              @click.stop="copyToClipboard(contactPhone)"
            >
              <DocumentDuplicateIcon class="h-4 w-4" />
            </button>
          </div>
          <div class="flex items-center gap-2 sm:gap-4 p-2.5 sm:p-4 rounded-lg sm:rounded-xl bg-slate-50 hover:bg-slate-100">
            <a :href="`mailto:${contactEmail}`" class="group flex items-center gap-2 sm:gap-3 min-w-0 flex-1">
              <div class="flex-shrink-0 w-9 h-9 sm:w-10 sm:h-10 rounded-lg bg-slate-200 flex items-center justify-center">
                <EnvelopeIcon class="h-5 w-5 sm:h-5 sm:w-5 text-slate-600" />
              </div>
              <div class="min-w-0">
                <span class="block text-[10px] sm:text-xs font-medium text-slate-500">Написать</span>
                <span class="block text-sm sm:text-base font-semibold text-slate-800 truncate">{{ contactEmail }}</span>
              </div>
            </a>
            <button
              type="button"
              class="flex-shrink-0 w-6 h-6 rounded flex items-center justify-center text-slate-400 hover:bg-slate-200 hover:text-slate-600"
              aria-label="Скопировать почту"
              @click.stop="copyToClipboard(contactEmail)"
            >
              <DocumentDuplicateIcon class="h-4 w-4" />
            </button>
          </div>
        </div>
      </div>

      <!-- Разделитель -->
      <div class="flex items-center gap-2 sm:gap-4">
        <div class="flex-1 h-px bg-slate-200" />
        <span class="text-[10px] sm:text-xs font-medium text-slate-400">или</span>
        <div class="flex-1 h-px bg-slate-200" />
      </div>

      <!-- Часть 2: Оставить заявку -->
      <div class="p-0">
        <p class="text-slate-600 text-sm mb-3 sm:mb-5">Оставьте заявку — мы свяжемся с вами</p>
        <ContactForm hide-title compact />
      </div>
    </div>
  </UiModal>
</template>

<script setup>
import { computed, inject } from 'vue'
import { PhoneIcon, EnvelopeIcon, DocumentDuplicateIcon } from '@heroicons/vue/24/outline'
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

const contactPhone = inject('contactPhone')
const contactEmail = inject('contactEmail')
const telHref = inject('telHref')

async function copyToClipboard(text) {
  const value = text?.value ?? text ?? ''
  if (!value) return
  try {
    await navigator.clipboard.writeText(value)
  } catch {
    const ta = document.createElement('textarea')
    ta.value = value
    ta.style.position = 'fixed'
    ta.style.opacity = '0'
    document.body.appendChild(ta)
    ta.select()
    document.execCommand('copy')
    document.body.removeChild(ta)
  }
}
</script>
