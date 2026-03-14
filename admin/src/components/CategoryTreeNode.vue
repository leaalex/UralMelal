<template>
  <div class="select-none">
    <div
      class="flex items-center gap-2 py-2 px-2 rounded-lg cursor-pointer hover:bg-slate-100 transition-colors min-w-0"
      :class="[
        { 'bg-slate-200 text-slate-900': isSelected },
        { 'text-slate-700': !isSelected }
      ]"
      :style="{ paddingLeft: `${depth * 8 + 4}px` }"
      @click="emit('select', node.id)"
    >
      <button
        v-if="hasChildren"
        type="button"
        class="shrink-0 w-5 h-5 flex items-center justify-center rounded hover:bg-slate-200"
        @click.stop="expanded = !expanded"
        :aria-label="expanded ? 'Свернуть' : 'Развернуть'"
      >
        <FolderIcon v-if="!expanded" class="h-4 w-4 text-slate-500" />
        <FolderOpenIcon v-else class="h-4 w-4 text-slate-500" />
      </button>
      <span v-else class="shrink-0 w-5 h-5 flex items-center justify-center">
        <TagIcon class="h-4 w-4 text-slate-500" />
      </span>
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
import { FolderIcon, FolderOpenIcon, TagIcon } from '@heroicons/vue/24/outline'

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
