<template>
  <body>
    <div class="app mt-40">
      <div class="text-8xl place-items-center">
        <h1>タイトルの名前</h1>
      </div>

      <div class="place-items-center mt-10">
        <button
          @click="home"
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-4 px-12 rounded fucus:outline-none focus:shadow-outline mt-3"
        >
          ゲームをプレイ
        </button>
      </div>
    </div>
  </body>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

function home() {
  router.push('/prepare')
}
</script>

<style>
.app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
body {
  background-color: rgb(217, 244, 244);
}
</style>
