import { createApp } from 'vue'
import { vMaska } from 'maska'
import App from './App.vue'
import router from './router'
import './style.css'

const app = createApp(App)
app.use(router)
app.directive('maska', vMaska)
app.mount('#app')
