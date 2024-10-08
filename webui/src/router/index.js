import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import LoginView from '../views/LoginView.vue';
import ProfileView from '../views/ProfileView.vue';

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
    },
    { path: '/profile/:id', name: 'profile', component: ProfileView }
  ]
});

// eslint-disable-next-line no-unused-vars
router.beforeEach(async (to, from) => {
  // if the user is not logged in and tries to access a page other than login, redirect to login
  let token = await localStorage.getItem('token_wasa_1982801');
  if (token === null && to.name !== 'login') {
    return { name: 'login' };
  }

  // if the user is logged in and tries to access the login page, don't allow it
  if (token !== null && to.name === 'login') {
    return { name: from.name };
  }
});

export default router;
