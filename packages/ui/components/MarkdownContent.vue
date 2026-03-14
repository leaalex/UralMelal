<template>
  <div class="prose prose-slate max-w-none" v-html="rendered"></div>
</template>

<script setup>
import { computed } from 'vue'
import MarkdownIt from 'markdown-it'

const props = defineProps({ content: { type: String, default: '' } })
const md = new MarkdownIt()
const rendered = computed(() => md.render(props.content || ''))
</script>
