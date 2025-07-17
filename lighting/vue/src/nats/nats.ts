// src/services/natsService.ts
import { ref, reactive } from 'vue'
import { connect, StringCodec, type NatsConnection, type Subscription } from 'nats.ws'

export type MsgHandler = (msg: string) => void

const handlersBySubject = new Map<string, Set<MsgHandler>>()

const subscriptions = new Map<string, Subscription>()

let nc: NatsConnection | null = null
const sc = StringCodec()

// Connect to NATS - call this once when your app starts
export async function connectNats() {
  if (nc) return // Already connected
  try {
    nc = await connect({ servers: `ws://${window.location.hostname}:8080` })
    console.log(`Connected to NATS server: ${nc.getServer()}`)
  } catch (err) {
    console.error('Error connecting to NATS:', err)
  }
}

// Subscribe to a subject
export function subscribeTo(subject: string, handler: (msg: string) => void) {
  if (!nc) return

  if (!handlersBySubject.has(subject)) {
    handlersBySubject.set(subject, new Set<MsgHandler>())
  }

  const handlers = handlersBySubject.get(subject)
  handlers?.add(handler)

  let sub
  if (!subscriptions.has(subject)) {
    sub = nc.subscribe(subject)
    subscriptions.set(subject, sub)
  } else {
    sub = subscriptions.get(subject)
  }

  ;(async () => {
    for await (const msg of sub!) {
      const data = sc.decode(msg.data)
      console.log('data', data)
      const handlers = handlersBySubject.get(subject)
      if (handlers) {
        for (const handler of handlers) {
          handler(data)
        }
      }
    }
  })()
}

// Unsubscribe from a subject (important for component cleanup)
export function unsubscribeFrom(subject: string, handler: (msg: string) => void) {
  const sub = subscriptions.get(subject)
  if (sub) {
    const handlers = handlersBySubject.get(subject) // Optional: clear messages on unsubscribe
    if (handlers) {
      handlers.delete(handler)
      if (handlers.size == 0) {
        sub.unsubscribe()
        subscriptions.delete(subject)
        handlersBySubject.delete(subject)
      }
    }
  }
}
