<template>
  <div class="container mx-auto app">
    <div class="grid grid-cols-1 place-items-center mb-0">
      <div class="app text-4xl h-32 font-bold my-0">{{ targetId }}さんとの繋がりが強い順に並べてください</div>
      <draggable
        v-model="changeableUsers"
        item-key="id"
        class="grid grid-cols-5 gap-4 h-16"
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

  <div class="flex justify-center w-full mt-3 relative"></div>
  <div class="flex justify-center w-full mt-3">
      <div class="triangle-left"></div>
      <div class="w-96 bg-gray-300 h-10 bg-gradient-to-r from-red-500 to-blue-300 via-red-300"></div>
      <div class="triangle-right"></div>
    </div>
    <div class="flex justify-center items-center gap-4">
      <div class="text-1xl h-10 font-bold mx-8">強い</div>
      <div class="text-1xl h-10 font-bold mx-24">繋がり</div>
      <div class="text-1xl h-10 font-bold mx-8">弱い </div>
    </div>

  <div class="flex justify-center">
    <button
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-3 px-5 rounded fucus:outline-none focus:shadow-outline mt-3"
      @click="checkResult"
    >
      結果を見る
    </button>
  </div>
  <div class="flex justify-center">
    <svg-icon type="mdi" :path="path" class="w-5"></svg-icon>
  </div>

  <button @click="showHintModal">showHINT</button>
  <div
    v-if="isModalOpen"
    class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full"
    @click.self="isModalOpen = false"
  >
    <div class="relative top-20 mx-auto p-5 border shadow-lg rounded-md bg-white w-3/4">
      <div class="mt-3 text-center">
        <h3 class="text-lg leading-6 font-medium text-gray-900">ヒント</h3>
        <div class="mt-2 px-7 py-3">
          <p class="text-sm text-gray-500">ヒントを見たい人をクリック</p>
        </div>
        <div class="grid grid-cols-5 gap-4 place-items-center my-4">
          <div v-for="user in changeableUsers" :key="user.id" class="flex justify-center">
            <img
              :src="user.url"
              alt="Profile Icon"
              class="rounded-full w-16 h-16"
              @click="showHint(user.id)"
            />
          </div>
        </div>
        <div
          v-if="isShowHint"
          class="max-w-4xl mx-auto my-4 overflow-hidden shadow-md sm:rounded-lg"
        >
          <div class="grid grid-cols-5 divide-x divide-gray-200">
            <div class="col-span-5 bg-gray-800 p-3">
              <p class="text-center text-lg font-semibold text-white">ユーザーグループ</p>
            </div>
            <template v-for="group in userGroups" :key="group">
              <div class="p-4 bg-white hover:bg-gray-50">
                <p class="text-center text-sm text-gray-900 break-words whitespace-normal">
                  {{ group }}
                </p>
              </div>
            </template>
          </div>
        </div>
        <div class="items-center px-4 py-3">
          <button
            class="px-4 py-2 bg-blue-500 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-300"
            @click="closeHintModal"
          >
            閉じる
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import draggable from 'vuedraggable'
// import { globalState} from '../store/store.ts'
import WebComponent from '../components/WebComponent.vue'
import { changeableUsers, targetId } from '.././store'
import SvgIcon from '@jamescoyle/vue-icon'
import { mdiLightbulbOnOutline } from '@mdi/js'

const path = mdiLightbulbOnOutline

const router = useRouter()
const checkResult = () => {
  router.push('/result')
}
const isModalOpen = ref(false)
const isShowHint = ref(false)
const showHintModal = () => {
  isModalOpen.value = true
}
const closeHintModal = () => {
  isModalOpen.value = false
  isShowHint.value = false
}

const userGroups = ref<User[]>([])
const showHint = async (id: string) => {
  try {
    const response = await fetch(`/api/users/${id}`)
    if (!response.ok) {
      alert('適切なデータが見つかりません')
      throw new Error('Network response was not ok')
    }
    const users = await response.json()
    userGroups.value = users.groups
    isShowHint.value = true
  } catch (error) {
    console.error('There was a problem with the fetch operation:', error)
  }
}
</script>

<style scoped>
.setWidth {
  width: 75%;
}
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
svg {
  height: 40px;
  width: 40px;
  position: relative;
  left: 240px;
  bottom: 215px;
}
</style>
