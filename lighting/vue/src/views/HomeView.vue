<script setup lang="ts">
import { ref } from 'vue'

const stuff = fetch('/about')
stuff.then((d) => {
  d.text().then((b) => {
    data.value = b
  })
})

const data = ref('')
</script>

<template>
  <main>
    {{ data }}
  </main>
</template>
