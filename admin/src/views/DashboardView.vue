<template>
  <div>
    <UiPageHeader title="Дашборд" />
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <UiStatCard label="Товаров в каталоге" :value="stats.products" />
      <UiStatCard v-if="auth.isManager" label="Новых заявок" :value="stats.leads" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { UiPageHeader, UiStatCard } from '@ui'
import client from '../api/client'
import { useAuthStore } from '../stores/auth'

const auth = useAuthStore()
const stats = ref({ products: 0, leads: 0 })

onMounted(async () => {
  try {
    const [p, l] = await Promise.all([
      client.get('/products?per_page=1'),
      auth.isManager ? client.get('/admin/leads?status=new') : Promise.resolve({ data: { total: 0 } }),
    ])
    stats.value.products = p.data.total ?? 0
    stats.value.leads = l.data?.total ?? 0
  } catch {}
})
</script>
