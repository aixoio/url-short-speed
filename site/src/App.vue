<template>
  <div class="p-8">
    <div class="flex justify-center">
      <h1 class="text-3xl">URL short</h1>
    </div>
    <div class="flex justify-center">
      <div class="grid grid-cols-1">
        <input type="text" placeholder="URL to shorten here!"
          class="outline-none border border-gray-950 p-2 m-3 rounded-lg" v-model="longurl">
        <button class="bg-blue-500 p-2 rounded-lg text-white">Shorten the link!</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

const longurl = ref("");


</script>
