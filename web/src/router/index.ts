import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/login.vue'

const routes = [
  {
    path: '/',
    name: 'Login',
    component: Login
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
