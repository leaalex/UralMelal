<template>
  <router-link
    :to="`/catalog/${product.id}`"
    class="block bg-white rounded-xl border border-slate-200 shadow-sm hover:shadow-md hover:border-slate-300 transition-all duration-200 overflow-hidden"
    :class="view === 'list' ? 'flex flex-row' : ''"
  >
    <div
      class="shrink-0 bg-slate-100"
      :class="view === 'list' ? 'w-32 h-32 md:w-40 md:h-40' : 'w-full h-44'"
    >
      <img
        v-if="product.image_path"
        :src="product.image_path"
        :alt="product.name"
        class="w-full h-full object-cover"
        loading="lazy"
      />
      <div
        v-else
        class="w-full h-full flex items-center justify-center text-slate-400 text-sm"
      >
        Нет фото
      </div>
    </div>
    <div
      class="flex flex-col p-4 min-w-0"
      :class="view === 'list' ? 'flex-1 justify-center' : ''"
    >
      <div class="flex flex-wrap gap-1.5 mb-2">
        <span
          v-if="product.size"
          class="inline-flex items-center rounded-lg bg-slate-100 px-2 py-0.5 text-xs font-medium text-slate-700"
        >
          {{ product.size }}
        </span>
        <span
          v-if="product.mark"
          class="inline-flex items-center rounded-lg bg-slate-100 px-2 py-0.5 text-xs font-medium text-slate-700"
        >
          {{ product.mark }}
        </span>
      </div>
      <h3 class="font-semibold text-slate-800 line-clamp-2 mb-1">
        {{ product.name }}
      </h3>
      <div class="flex flex-wrap items-center gap-x-3 gap-y-1 text-sm text-slate-500 mt-auto">
        <span v-if="product.city">{{ product.city }}</span>
        <span
          v-if="product.sku"
          class="text-slate-400"
        >
          Арт. {{ product.sku }}
        </span>
      </div>
      <div class="flex items-center justify-between gap-2 mt-2">
        <span
          v-if="lowestPrice"
          class="text-sm font-semibold text-slate-800"
        >
          от {{ formatPrice(lowestPrice) }}
        </span>
        <span
          v-else
          class="text-sm text-slate-400"
        >
          Цена по запросу
        </span>
        <span
          v-if="hasStock"
          class="inline-flex items-center gap-1 text-xs text-emerald-600"
        >
          <span class="h-1.5 w-1.5 rounded-full bg-emerald-500" />
          В наличии
        </span>
        <span
          v-else
          class="text-xs text-slate-400"
        >
          —
        </span>
      </div>
    </div>
  </router-link>
</template>

<script setup>
import { computed } from 'vue'
import { formatPrice, getLowestPrice } from '../utils/format'

const props = defineProps({
  product: { type: Object, required: true },
  view: { type: String, default: 'grid' },
})

const lowestPrice = computed(() => getLowestPrice(props.product))

const hasStock = computed(() => {
  const s = props.product?.stock
  return s != null && s > 0
})
</script>
