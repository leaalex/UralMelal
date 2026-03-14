<template>
  <div class="mt-1">
    <div class="flex items-center gap-1 min-w-0">
      <button
        v-if="hasChildren"
        type="button"
        class="shrink-0 p-1 rounded-xl text-slate-500 hover:text-slate-600 hover:bg-slate-100"
        :aria-label="expanded ? 'Свернуть' : 'Развернуть'"
        @click="expanded = !expanded"
      >
        <ChevronRightIcon
          class="h-4 w-4 transition-transform"
          :class="{ 'rotate-90': expanded }"
        />
      </button>
      <button
        type="button"
        class="flex-1 min-w-0 text-left px-3 py-2 rounded-xl text-sm font-medium transition-colors truncate"
        :class="isSelected ? 'bg-slate-200 text-slate-900 font-semibold' : 'text-slate-700 hover:bg-slate-100'"
        @click="emit('select', category.id)"
      >
        {{ category.name }}
      </button>
    </div>
    <div v-if="hasChildren && expanded" class="ml-4 mt-0.5 border-l border-slate-200 pl-2">
      <CategoryTree
        v-for="child in category.children"
        :key="child.id"
        :category="child"
        :selected-id="selectedId"
        @select="emit('select', $event)"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { ChevronRightIcon } from '@heroicons/vue/24/outline'

const props = defineProps({
  category: { type: Object, required: true },
  selectedId: { type: [Number, String, null], default: null },
})

const emit = defineEmits(['select'])

const expanded = ref(false)

const hasChildren = computed(() => {
  const children = props.category?.children
  return Array.isArray(children) && children.length > 0
})

const isSelected = computed(() => props.selectedId === props.category.id)

watch(
  () => props.selectedId,
  (id) => {
    if (id && hasChildren.value) {
      const isDescendant = findInTree(props.category.children, id)
      if (isDescendant) expanded.value = true
    }
  },
  { immediate: true }
)

function findInTree(nodes, targetId) {
  if (!nodes?.length) return false
  for (const n of nodes) {
    if (n.id === targetId) return true
    if (findInTree(n.children, targetId)) return true
  }
  return false
}
</script>
