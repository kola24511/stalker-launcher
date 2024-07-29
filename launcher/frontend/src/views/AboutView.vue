<script setup lang="ts">
import { useI18n } from "vue-i18n";

const { t } = useI18n();
</script>

<template>
  <div class="about">
    <div class="flex justify-center items-center h-full">
      <div class="text-xl flex">
        {{ t("aboutpage.title") }} 
         <span class="text-blue-800">
          <a href="https://github.com/kola24511"> kola2451</a>
         </span>
      </div>
    </div>
  </div>
</template>
