import { createApp } from 'vue'
import App from './App.vue'
import router from './routers'
import antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css'

const app = createApp(App)
app.use(router).use(antd).mount('#app')
