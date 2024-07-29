import { createI18n } from "vue-i18n";

import ru from "./locales/ru.json";
import en from "./locales/en.json";

const i18n = createI18n({
  locale: "ru",
  fallbackLocale: "ru",
  legacy: false,
  messages: {
    "ru": ru,
    en: en,
  },
});

export default i18n;
