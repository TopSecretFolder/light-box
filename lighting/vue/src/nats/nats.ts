// src/services/natsService.ts
import { ref, reactive } from 'vue'
import { connect, StringCodec, type NatsConnection, type Subscription } from 'nats.ws'

// Reactive state shared globally
export const isConnected = ref(false)
const messagesBySubject = reactive<Map<string, string[]>>(new Map())
const subscriptions = new Map<string, Subscription>()

let nc: NatsConnection | null = null
const sc = StringCodec()

// Connect to NATS - call this once when your app starts
export async function connectNats() {
  if (nc) return // Already connected
  try {
    nc = await connect({ servers: `ws://${window.location.host}` })
    isConnected.value = true
    console.log(`Connected to NATS server: ${nc.getServer()}`)
  } catch (err) {
    console.error('Error connecting to NATS:', err)
  }
}

// Subscribe to a subject
export function subscribeTo(subject: string) {
  if (!nc || subscriptions.has(subject)) return

  // Ensure an array exists for this subject's messages
  if (!messagesBySubject.has(subject)) {
    messagesBySubject.set(subject, [])
  }

  const sub = nc.subscribe(subject)
  subscriptions.set(subject, sub)
  ;(async () => {
    for await (const msg of sub) {
      const data = sc.decode(msg.data)
      messagesBySubject.get(subject)?.push(data)
    }
  })()
}

// Unsubscribe from a subject (important for component cleanup)
export function unsubscribeFrom(subject: string) {
  const sub = subscriptions.get(subject)
  if (sub) {
    sub.unsubscribe()
    subscriptions.delete(subject)
    messagesBySubject.delete(subject) // Optional: clear messages on unsubscribe
  }
}

// Expose the messages map for components to use
export { messagesBySubject }
