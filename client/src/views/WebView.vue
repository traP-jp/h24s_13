<template>
  <WebComponent />
  <ListComponent />
  <div class="mx-auto w-1/4">
    <button
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-3 px-5 rounded focus:outline-none focus:shadow-outline my-4 w-full"
      @click="router.push('/prepare')"
    >
      ゲームに戻る
    </button>
  </div>
</template>

<script setup lang="ts">
import ListComponent from '../components/ListComponent.vue'
import WebComponent from '../components/WebComponent.vue'
import { useRouter } from 'vue-router'

const router = useRouter()
</script>

<style scoped></style>
