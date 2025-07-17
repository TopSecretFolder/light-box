import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import VueKonva from 'vue-konva'
import { connectNats } from './nats/nats'

connectNats()
  .then(() => {
    const app = createApp(App)

    app.use(VueKonva)
    app.use(router)

    app.mount('#app')
  })
  .catch((error) => {
    console.error(error)
  })

// setup nats
