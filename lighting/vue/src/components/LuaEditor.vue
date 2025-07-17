<script setup lang="ts">
import { ref } from 'vue'
import { Codemirror } from 'vue-codemirror'

const code = ref('-- Another Lua example\nlocal name = "World"\nprint("Hello, " .. name)')
</script>

<template>
  <Codemirror
    v-model="code"
    :style="{ height: '400px' }"
    :autofocus="true"
    :indent-with-tab="true"
    :tab-size="2"
  />
</template>
