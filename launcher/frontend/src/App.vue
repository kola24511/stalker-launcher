<script setup lang="ts">
import { useI18n } from "vue-i18n";

const { t, availableLocales: languages, locale } = useI18n();

const onclickLanguageHandle = (item: string) => {
  item !== locale.value ? (locale.value = item) : false;
};

const onclickMinimise = () => {};

const onclickQuit = () => {};

document.body.addEventListener("click", function (event) {
  event.preventDefault();
});
</script>

<template>
  <!-- Header -->
  <div class="w-full h-16 bg-stone-900">
    <div class="flex items-center justify-between h-full px-4 text-white">
      <div class="nav space-x-6">
        <router-link to="/">{{ t("nav.home") }}</router-link>
        <router-link to="/about">{{ t("nav.about") }}</router-link>
      </div>
      <!-- Menu -->
    <div class="menu">
      <div class="language flex space-x-4 hover:cursor-pointer">
        <div
          v-for="item in languages"
          :key="item"
          :class="{ active: item === locale }"
          @click="onclickLanguageHandle(item)"
          class="lang-item"
        >
          {{ t("languages." + item) }}
        </div>
      </div>
    </div>
    </div>
  </div>
  <!-- Page -->
  <div class="view h-[537px] max-h-[537px]">
    <router-view />
  </div>
</template>