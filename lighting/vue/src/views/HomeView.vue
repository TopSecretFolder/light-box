<script setup lang="ts">
import { subscribeTo, unsubscribeFrom } from '@/nats/nats'
import { computed, onMounted, onUnmounted, reactive, ref } from 'vue'
import LuaEditor from '@/components/LuaEditor.vue'

const about = ref('')

fetch('/about').then((d) => {
  d.text().then((b) => {
    about.value = b
  })
})

type DebugMsg = {
  index: number
  hex: string
  brightness: number
}

const leds = reactive(new Map<number, string>())

const sorted = computed(() => {
  const result = Array(leds.size).fill('')
  for (const [key, value] of leds) {
    result[key] = value
  }
  return result
})

const ledDebugHandler = (msg: string) => {
  const parsed = JSON.parse(msg) as DebugMsg
  leds.set(parsed.index, parsed.hex)
}

onMounted(() => {
  subscribeTo('led-debug', ledDebugHandler)
})

onUnmounted(() => {
  unsubscribeFrom('led-debug', ledDebugHandler)
})
</script>

<template>
  <main>
    <section>
      {{ about }}
    </section>
    <section class="flex flex-row gap-4">
      <span v-for="(hex, index) in sorted" :key="index" :style="{ color: `#${hex}` }">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          height="24px"
          viewBox="0 -960 960 960"
          width="24px"
          :fill="`#${hex}`"
        >
          <path
            d="M480-80q-33 0-56.5-23.5T400-160h160q0 33-23.5 56.5T480-80ZM320-200v-80h320v80H320Zm10-120q-69-41-109.5-110T180-580q0-125 87.5-212.5T480-880q125 0 212.5 87.5T780-580q0 81-40.5 150T630-320H330Zm24-80h252q45-32 69.5-79T700-580q0-92-64-156t-156-64q-92 0-156 64t-64 156q0 54 24.5 101t69.5 79Zm126 0Z"
          />
        </svg>
      </span>
    </section>
    <section>
      <LuaEditor />
    </section>
  </main>
</template>
