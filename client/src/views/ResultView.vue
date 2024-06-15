<template>
  <div>This is an About page.</div>
  <button
    class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-3 px-5 rounded fucus:outline-none focus:shadow-outline mt-3"
    @click="showArray"
  >
    見る
  </button>
</template>

<script setup lang="ts">
import { userInfo } from '.././store.ts'
const showArray = () => {
  console.log(userInfo.value)
}
</script>

<style scoped></style>
