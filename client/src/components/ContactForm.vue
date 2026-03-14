<template>
  <div :class="hideTitle ? '' : 'mt-6'">
    <h3 v-if="!hideTitle" class="text-lg font-semibold mb-4">Оставить заявку</h3>
    <form @submit.prevent="submit" class="space-y-4 w-full min-w-[320px]">
      <input v-model="form.website_url" type="text" name="website" class="absolute -left-[9999px]" tabindex="-1" autocomplete="off" />
      <div>
        <label class="block text-sm font-medium text-slate-700 mb-1">Имя *</label>
        <input v-model="form.name" type="text" required class="w-full border border-slate-300 rounded-lg px-3 py-2.5 text-slate-800 focus:ring-2 focus:ring-slate-300 focus:border-slate-400 outline-none" />
      </div>
      <div>
        <label class="block text-sm font-medium text-slate-700 mb-1">Телефон *</label>
        <input v-model="form.phone" v-maska="'+7 (###) ###-##-##'" type="tel" required class="w-full border border-slate-300 rounded-lg px-3 py-2.5 text-slate-800 focus:ring-2 focus:ring-slate-300 focus:border-slate-400 outline-none" />
      </div>
      <div>
        <label class="block text-sm font-medium text-slate-700 mb-1">Тема</label>
        <input v-model="form.subject" type="text" class="w-full border border-slate-300 rounded-lg px-3 py-2.5 text-slate-800 focus:ring-2 focus:ring-slate-300 focus:border-slate-400 outline-none" />
      </div>
      <label class="flex items-start gap-2 cursor-pointer">
        <input v-model="form.consent" type="checkbox" required class="mt-1" />
        <span class="text-sm text-slate-700">Согласен на <a href="/privacy" class="underline">обработку персональных данных</a></span>
      </label>
      <p v-if="message" :class="success ? 'text-emerald-600' : 'text-red-600'" class="text-sm">{{ message }}</p>
      <div class="flex justify-end pt-1">
        <UiButton type="submit">Отправить</UiButton>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import { UiButton } from '@ui'
defineProps({ hideTitle: { type: Boolean, default: false } })
import { useRoute } from 'vue-router'
import client from '../api/client'

const route = useRoute()

const form = reactive({
  name: '',
  phone: '',
  subject: '',
  website_url: '',
  consent: false,
})
const message = ref('')
const success = ref(false)

watch(() => route.query.contact_subject, (subject) => {
  if (subject) form.subject = subject
}, { immediate: true })

async function submit() {
  message.value = ''
  if (!form.consent) {
    message.value = 'Необходимо согласие на обработку ПД'
    return
  }
  try {
    await client.post('/leads', {
      name: form.name,
      phone: form.phone.replace(/\D/g, ''),
      subject: form.subject,
      website_url: form.website_url,
      consent: form.consent,
    })
    message.value = 'Заявка отправлена!'
    success.value = true
    form.name = ''
    form.phone = ''
    form.subject = ''
  } catch {
    message.value = 'Ошибка отправки'
    success.value = false
  }
}
</script>
