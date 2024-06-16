<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { API_URL, changeableUsers, targetId } from '.././store'
const router = useRouter()
const traQID = ref<string>('')
const getFivePeople = async (id: string) => {
  if (id === '') {
    alert('traQIDを入力してください')
    return
  }
  targetId.value = id
  try {
    const response = await fetch(`${API_URL}/quiz/new?id=${id}`)
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
    const userIds = await response.json()
    changeableUsers.value = userIds.map((userId: string) => ({ id: userId }))
    router.push('/game')
  } catch (error) {
    // TODO:error handling
    console.error('There was a problem with the fetch operation:', error)
  }
}
</script>

<template>
  <br /><br /><br />
  <h1 class="text-center text-2xl font-bold">ゲームの説明</h1>
  <br /><br />
  <p class="max-w-2xl mx-auto">
    traQidを入力してスタートボタンを押すとゲームが始まります。（入力がない場合はランダムで選ばれます）ゲーム画面では選んだ人（以下hogeさんと呼ぶ）とランダムで5人が表示されるので、5人をhogeさんと繋がりの強い順に並び変えてください。
  </p>
  <p class="max-w-2xl mx-auto">
    ただし“A さんと B さん繋がりの強さ” := “A さんが B さんの times で発言した数” + “B さんが A
    さんの times で発言した数”
  </p>
  <br /><br />
  <div class="flex justify-center">
    <label for="traQid" class="">traQIDを入力してください:</label>
    <input
      id="traQid"
      type="text"
      v-model="traQID"
      class="bg-white px-3 text-gray-700 border border-gray-700 rounded"
    />
  </div>

  <br />
  <div class="flex justify-center">
    <button
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-3 px-5 rounded fucus:outline-none focus:shadow-outline mt-3"
      @click="getFivePeople(traQID)"
    >
      ゲームスタート
    </button>
  </div>
</template>

<style></style>
