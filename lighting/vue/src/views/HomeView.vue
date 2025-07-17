<script setup lang="ts">
import { subscribeTo, unsubscribeFrom, messagesBySubject } from '@/nats/nats'
import { computed, onMounted, onUnmounted, ref } from 'vue'

const about = ref('')

fetch('/about').then((d) => {
  d.text().then((b) => {
    about.value = b
  })
})

onMounted(() => {
  subscribeTo('led-debug')
})

onUnmounted(() => {
  unsubscribeFrom('led-debug')
})

const ledDebugMsgs = computed(() => {
  return messagesBySubject.get('led-debug') ?? []
})
</script>

<template>
  <main>
    <section>
      {{ about }}
    </section>
    <section>
      {{ ledDebugMsgs }}
    </section>
  </main>
</template>
