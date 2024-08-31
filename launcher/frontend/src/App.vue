<script setup lang="ts">
import { useI18n } from "vue-i18n";
import background from "@/assets/img/background.jpg";

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
  <div class="relative">
    <!-- Картинка на заднем фоне -->
    <img class="absolute inset-0 w-full h-full object-cover z-[-10]" :src="background" />

    <!-- Header -->
    <div class="w-full h-16 bg-gray-700 bg-opacity-5 backdrop-blur-sm border-b border-gray-700/50">
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
    <div class="view h-[537px] max-h-[537px] relative">
      <router-view />
    </div>
  </div>
</template>