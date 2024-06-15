<template>
  <div class="text-8xl aboutw-100 flex justify-center mt-5">
    <h1>タイトルの名前</h1>
  </div>
  <div class="flex justify-center items-center h-screen">
    <button @click="home" class="text-7xl bg-green-400 px-8 py-4">
      ゲームをプレイ
    </button>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

function home() {
  router.push('/router')
}
</script>

<style>
</style>

