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
    <button @click="checkResult" class="bg-sky-400 rounded-lg text-white font-bold">
      結果を見る
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import draggable from 'vuedraggable'

interface User {
  id: string
}

const userInfo = ref<User[]>([
  { id: 'masky5859' },
  { id: 'pirosiki' },
  { id: 'cp20' },
  { id: 'toki' },
  { id: 'Series_205' }
])

const checkResult = () => {
  alert('結果を見るボタンが押されました')
  console.log(userInfo.value)
}
</script>

<style scoped></style>
