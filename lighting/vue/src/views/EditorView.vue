<template>
  <div class="bezier-editor-container">
    <EditorTrack v-model="curves" />
  </div>
</template>
<script setup lang="ts">
import EditorTrack, { type BezierSegment } from '@/components/EditorTrack.vue'
import { ref } from 'vue'

const curves = ref<BezierSegment[]>([
  {
    p0: { x: 0, y: 50 },
    p1: { x: 20, y: 50 },
    p2: { x: 40, y: 50 },
    p3: { x: 60, y: 50 },
  },
])
</script>
