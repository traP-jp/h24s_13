<template>
  <div class="grid grid-cols-1 place-items-center">
    <div><h1 class="text-6xl">Welcome to GameView!</h1></div>
    <div class="text-4xl h-40">hogeさんとの繋がりが近い順に並べてください</div>
    <draggable v-model="userInfo" item-key="id" class="grid grid-cols-5 gap-4" handle=".handle">
      <template #item="{ element }">
        <span class="handle">
          <img
            :src="`https://q.trap.jp/api/v3/public/icon/${element.id}`"
            alt="Profile Icon"
            class="rounded-full w-16 h-16"
          />
        </span>
      </template>
    </draggable>
  </div>
  <br/>
  <div class="flex justify-center w-full mt-3">
      <div class="w-96 bg-gray-300 h-10 bg-gradient-to-r from-red-500 to-blue-500 via-green-400"></div>
    </div>
  <br/>
<div class="flex justify-center">
    <button
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-3 px-5 rounded fucus:outline-none focus:shadow-outline mt-3"
      @click="checkResult"
    >
      結果を見る
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import draggable from 'vuedraggable'
// import { globalState} from '../store/store.ts'
import WebComponent from '../components/WebComponent.vue'
import { userInfo } from '.././store.ts'

const router = useRouter()
const checkResult = () => {
  alert('結果を見るボタンが押されました')
  console.log(userInfo.value)
  router.push('/result')
}
</script>

<style scoped></style>
