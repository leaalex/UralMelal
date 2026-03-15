<template>
  <tr
    class="group hover:bg-slate-50 cursor-pointer border-b border-slate-100 last:border-0"
    @click="navigate"
  >
    <td class="px-4 py-3 w-20 align-middle">
      <div class="w-16 h-16 rounded-lg bg-slate-100 overflow-hidden shrink-0">
        <img
          v-if="product.image_path"
          :src="product.image_path"
          :alt="product.name"
          class="w-full h-full object-cover"
          loading="lazy"
        />
        <div
          v-else
          class="w-full h-full flex items-center justify-center text-slate-400 text-xs"
        >
          Нет фото
        </div>
      </div>
    </td>
    <td class="px-4 py-3 align-middle">
      <span class="font-medium text-slate-800 line-clamp-2">{{ product.name }}</span>
    </td>
    <td class="px-4 py-3 w-24 align-middle text-sm text-slate-600">
      <span v-if="product.sku">Арт. {{ product.sku }}</span>
      <span v-else class="text-slate-400">—</span>
    </td>
    <td class="px-4 py-3 w-20 align-middle text-sm text-slate-600">
      {{ product.size || '—' }}
    </td>
    <td class="px-4 py-3 w-24 align-middle text-sm text-slate-600">
      {{ product.mark || '—' }}
    </td>
    <td class="px-4 py-3 w-24 align-middle text-sm text-slate-600">
      {{ product.city || '—' }}
    </td>
    <td class="px-4 py-3 w-28 align-middle text-sm">
      <span v-if="lowestPrice" class="font-medium text-slate-800">
        от {{ formatPrice(lowestPrice) }}
      </span>
      <span v-else class="text-slate-400">Цена по запросу</span>
    </td>
    <td class="px-4 py-3 w-24 align-middle text-sm">
      <span
        v-if="hasStock"
        class="inline-flex items-center gap-1 text-emerald-600"
      >
        <span class="h-1.5 w-1.5 rounded-full bg-emerald-500" />
        В наличии
      </span>
      <span v-else class="text-slate-400">—</span>
    </td>
  </tr>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { formatPrice, getLowestPrice } from '../utils/format'

const props = defineProps({
  product: { type: Object, required: true },
})

const router = useRouter()

const lowestPrice = computed(() => getLowestPrice(props.product))

const hasStock = computed(() => {
  const s = props.product?.stock
  return s != null && s > 0
})

function navigate() {
  router.push(`/catalog/${props.product.id}`)
}
</script>
