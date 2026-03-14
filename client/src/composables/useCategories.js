import { ref, computed } from 'vue'
import client from '../api/client'

const categoriesCache = ref([])

export function useCategories() {
  const categories = ref([])

  async function loadCategories() {
    if (categoriesCache.value.length) {
      categories.value = categoriesCache.value
      return true
    }
    try {
      const { data } = await client.get('/categories')
      categories.value = data ?? []
      categoriesCache.value = categories.value
      return true
    } catch (e) {
      return false
    }
  }

  const rootCategories = computed(() =>
    (categories.value ?? []).filter((c) => !c.parent_id)
  )

  function findCategoryById(list, id) {
    if (!list?.length || (id !== 0 && !id)) return null
    const numId = Number(id)
    for (const c of list) {
      if (Number(c.id) === numId) return c
      const found = findCategoryById(c.children || [], id)
      if (found) return found
    }
    return null
  }

  return {
    categories,
    rootCategories,
    loadCategories,
    findCategoryById,
  }
}
