<template>
  <div class="grid grid-cols-1 place-items-center">
    <div class="text-4xl h-40">{{ targetId }}さんとの繋がりが近い順に並べてください</div>
    <draggable
      v-model="changeableUsers"
      item-key="id"
      class="grid grid-cols-5 gap-4"
      handle=".handle"
    >
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

  <br />
  <div class="flex justify-center w-full mt-3 relative"></div>
  <div class="flex justify-center w-full mt-3">
    <div class="triangle-left"></div>
    <div class="w-96 bg-gray-300 h-10 bg-gradient-to-r from-red-500 to-blue-300 via-red-300"></div>
    <div class="triangle-right"></div>
  </div>
  <br />
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
import { changeableUsers, targetId } from '.././store'

const router = useRouter()
const checkResult = () => {
  router.push('/result')
}
</script>

<style scoped>
.triangle-left {
  width: 0;
  height: 0;
  border-top: 20px solid transparent;
  border-bottom: 20px solid transparent;
  border-right: 20px solid #ed2b2be6;
}
.triangle-right {
  width: 0;
  height: 0;
  border-top: 20px solid transparent;
  border-bottom: 20px solid transparent;
  border-left: 20px solid #60a5faa3;
}
</style>
