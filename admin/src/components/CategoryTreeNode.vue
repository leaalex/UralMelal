<template>
  <div class="select-none">
    <div
      class="flex items-center gap-2 py-2 px-2 rounded-lg cursor-pointer hover:bg-slate-100 transition-colors min-w-0"
      :class="[
        { 'bg-slate-200 text-slate-900': isSelected },
        { 'text-slate-700': !isSelected }
      ]"
      :style="{ paddingLeft: `${depth * 16 + 8}px` }"
      @click="emit('select', node.id)"
    >
      <button
        v-if="hasChildren"
        type="button"
        class="shrink-0 p-1 rounded hover:bg-slate-200 -ml-1"
        @click.stop="expanded = !expanded"
      >
        <ChevronRightIcon
          class="h-4 w-4 transition-transform"
          :class="{ 'rotate-90': expanded }"
        />
      </button>
      <span v-else class="w-6 shrink-0" />
      <span class="flex-1 truncate">{{ node.name }}</span>
      <span
        v-if="node.product_count > 0"
        class="shrink-0 text-xs bg-slate-200 text-slate-600 rounded-full px-2 py-0.5"
      >
        {{ node.product_count }}
      </span>
    </div>
    <div v-if="hasChildren && expanded">
      <CategoryTreeNode
        v-for="child in node.children"
        :key="child.id"
        :node="child"
        :selected-id="selectedId"
        :depth="depth + 1"
        @select="emit('select', $event)"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { ChevronRightIcon } from '@heroicons/vue/24/outline'

const props = defineProps({
  node: { type: Object, required: true },
  selectedId: { type: [Number, String, null], default: null },
  depth: { type: Number, default: 0 },
})

const emit = defineEmits(['select'])

const expanded = ref(true)

const hasChildren = computed(() => {
  const ch = props.node?.children
  return Array.isArray(ch) && ch.length > 0
})

const isSelected = computed(() => props.selectedId === props.node.id)

function isDescendant(nodes, targetId) {
  if (!nodes?.length) return false
  for (const n of nodes) {
    if (n.id === targetId) return true
    if (isDescendant(n.children, targetId)) return true
  }
  return false
}

watch(
  () => props.selectedId,
  (id) => {
    if (id && hasChildren.value) {
      if (isDescendant(props.node.children, id)) expanded.value = true
    }
  },
  { immediate: true }
)
</script>
