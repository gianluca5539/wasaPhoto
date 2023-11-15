import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    }
  ]
})

// eslint-disable-next-line no-unused-vars
router.beforeEach(async (to, from) => {
  // if the user is not logged in and tries to access a page other than login, redirect to login
  let token = await localStorage.getItem('token')
  if (token === null && to.name !== 'login') {
    return { name: 'login' }
  }

  // if the user is logged in and tries to access the login page, don't allow it
  if (token !== null && to.name === 'login') {
    return { name: from.name }
  }
})

export default router
