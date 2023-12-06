import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import axios from './services/axios.js';

import './assets/dashboard.css';
import './assets/main.css';

const app = createApp(App);
app.config.globalProperties.$axios = axios;
app.use(router);
app.mount('#app');
