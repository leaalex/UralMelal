/**
 * Format price for display (Russian locale)
 * @param {number} value - Price value
 * @returns {string} Formatted string, e.g. "12 500 р/т"
 */
export function formatPrice(value) {
  if (value == null || value <= 0 || Number.isNaN(value)) return ''
  return new Intl.NumberFormat('ru-RU', {
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  }).format(value) + ' р/т'
}

/**
 * Get lowest non-zero price from product
 * @param {Object} product - Product with price_1t, price_5t, price_10t
 * @returns {number|null} Lowest price or null
 */
export function getLowestPrice(product) {
  const prices = [product?.price_1t, product?.price_5t, product?.price_10t].filter(
    (p) => p != null && p > 0 && !Number.isNaN(p)
  )
  return prices.length ? Math.min(...prices) : null
}

/**
 * Parse characteristics string into key-value array
 * Format: "Дн: 50; Вид: Зонт; Вес, кг: 0,09"
 * @param {string} str
 * @returns {Array<{key: string, value: string}>}
 */
export function parseSpecs(str) {
  if (!str || typeof str !== 'string') return []
  return str
    .split(';')
    .map((s) => s.trim())
    .filter(Boolean)
    .map((s) => {
      const colon = s.indexOf(':')
      if (colon < 0) return { key: s, value: '' }
      return { key: s.slice(0, colon).trim(), value: s.slice(colon + 1).trim() }
    })
}
