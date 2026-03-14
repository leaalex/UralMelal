<template>
  <div class="overflow-x-auto rounded-xl border border-slate-200 bg-white shadow-sm ring-1 ring-slate-200/50">
    <table class="w-full">
      <thead class="bg-slate-50/80">
        <tr>
          <th
            v-for="col in columns"
            :key="col.key"
            class="px-4 py-3 text-left text-xs font-semibold uppercase tracking-wider text-slate-600"
            :class="col.align === 'right' ? 'text-right' : 'text-left'"
          >
            {{ col.label }}
          </th>
          <th v-if="$slots.actions" class="px-4 py-3 text-right text-xs font-semibold uppercase tracking-wider text-slate-600"></th>
        </tr>
      </thead>
      <tbody class="divide-y divide-slate-100">
        <tr
          v-for="(row, i) in data"
          :key="rowKey ? row[rowKey] : i"
          class="transition-colors duration-150 even:bg-slate-50/50 hover:bg-slate-100/80"
        >
          <td
            v-for="col in columns"
            :key="col.key"
            class="px-4 py-3 text-sm text-slate-700"
            :class="col.align === 'right' ? 'text-right' : 'text-left'"
          >
            <slot :name="`cell-${col.key}`" :row="row" :column="col" :value="row[col.key]">
              {{ row[col.key] ?? '-' }}
            </slot>
          </td>
          <td v-if="$slots.actions" class="px-4 py-3 text-right">
            <slot name="actions" :row="row" />
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
defineProps({
  columns: { type: Array, required: true },
  data: { type: Array, default: () => [] },
  rowKey: { type: String, default: 'id' },
})
</script>
