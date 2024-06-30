<template>
  <div class="p-8">
    <div class="flex justify-center">
      <h1 class="text-3xl">URL short</h1>
    </div>
    <div class="flex justify-center">
      <div class="grid grid-cols-1">
        <input type="text" placeholder="URL to shorten here!"
          class="outline-none border border-gray-950 p-2 m-3 rounded-lg" v-model="longurl">
        <button class="bg-blue-500 p-2 rounded-lg text-white" @click="shortme">Shorten the link!</button>
        <span class="bg-red-200 border border-red-950 text-red-950 p-2 m-4 rounded-lg" v-show="error != ''">{{ error
          }}</span>
        <span class="bg-blue-200 border border-blue-950 text-blue-950 p-2 m-4 rounded-lg overflow-visible text-wrap" v-show="shorted != ''">Here is the short link: <a :href="shorted" target="_blank">{{ shorted }}</a></span>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { isEmpty } from 'lodash';
import { ref } from 'vue';
import { MakeShort } from './assets/ts/api';

const longurl = ref("");
const error = ref("");
const shorted = ref("");

async function shortme() {
  if (isEmpty(longurl.value.trim())) {
    error.value = "You must enter a url"
    return
  }
  error.value = ""

  const res = await MakeShort(longurl.value)
  shorted.value = `http://${window.location.host}${res.short}`
}


</script>
