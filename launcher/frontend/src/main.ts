import { createApp } from "vue";
import { createPinia } from "pinia";
import PrimeVue from 'primevue/config';
import Aura from '@primevue/themes/aura';

import App from "./App.vue";
import router from "./router";
import i18n from "./i18n";

import "./style.scss";

// import "./assets/main.css";

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(i18n);
app.use(PrimeVue, { theme: { preset: Aura }});

app.mount("#app");
