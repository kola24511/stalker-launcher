<script setup lang="ts">
import { useI18n } from "vue-i18n";

const { t } = useI18n();
</script>

<template>
  <div class="w-9/12 mx-auto pt-6">
    <div class="bg-gray-700 bg-opacity-10 backdrop-blur-sm border border-gray-700/50">
      <div class="flex flex-col justify-center items-center h-full">
        <div class="text-xl flex">
          <span>
            {{ t("aboutpage.title") }} 
          </span>
          <span class="text-blue-700">
            <a href="https://github.com/kola24511">
              kola2451
            </a>
          </span>
        </div>
      <div>
        {{ t("aboutpage.about-text") }} 
      </div>
    </div>
    </div>
  </div>
</template>
