<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { ref } from "vue";
import {StartGame} from "../../wailsjs/go/main/App"
import {UpdateClient} from "../../wailsjs/go/main/App"

import option from "@/assets/svg/option.svg";

const { t } = useI18n();
const isMenuOpen = ref(false);
const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value;
};


</script>

<template>
  <div class="flex flex-row h-full">
    <div class="w-3/12 relative bg-gray-700 bg-opacity-5 backdrop-blur-sm border-r border-gray-700/50">
      <div class="flex flex-row justify-around absolute inset-x-0 bottom-0 space-x-1 text-center">
        <button class="w-10/12 bg-[#aab15e] p-2.5" @click="StartGame">
          <div class="font-semibold text-xl text-white">
            {{ t("homepage.button-game") }}
          </div>
        </button>
        <button @click="toggleMenu" class="w-2/12 bg-[#aab15e] px-2 py-2">
          <img class="h-8 w-8" :src="option" alt="">
        </button>
        <div v-show="isMenuOpen" class="bg-[#aab15e] rounded absolute right-0 bottom-14">
          <div class="flex flex-col text-white p-1.5">
            <button @click="UpdateClient" class="border-b py-1 text-left">
              {{ t("homepage.button-check-update") }}
            </button>
            <button class="text-left">
              {{ t("homepage.button-force-update") }}
            </button>
          </div>
        </div>
      </div>
    </div>
    <div class="w-9/12 grid justify-items-end">
      <iframe src="https://discord.com/widget?id=959159181053661244&theme=dark" width="350" height="537" allowtransparency="true" frameborder="0" sandbox="allow-popups allow-popups-to-escape-sandbox allow-same-origin allow-scripts"></iframe>
    </div>
  </div>
</template>