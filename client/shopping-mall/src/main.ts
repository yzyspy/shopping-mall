import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from "./router";
import pinia from "./store";
import service from "./api/request.ts";

const app = createApp(App)
app.use(router)
app.use(pinia)
app.use(service)
app.mount('#app')
