<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { changeableUsers, targetId } from '@/store'

const router = useRouter()

const traQID = ref<string>('')

const startWithRandom = async () => {
  const me = await (await fetch(`/api/users/me`)).json()
  const myID = me.id
  const randomFriends = await (await fetch(`/api/users/${myID}/random?count=1`)).json()
  await start(randomFriends[0])
}

const startWithTraqID = async (id: string) => {
  if (id === '') {
    alert('traQ IDを入力してください')
    return
  }
  await start(id)
}

const start = async (targetID: string) => {
  targetId.value = targetID
  try {
    const response = await fetch(`api/quiz/new?id=${targetID}`)
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
  <div class="container mx-auto app flex flex-col gap-8">
    <h1 class="text-4xl font-bold">あそびかた</h1>
    <div class="mx-auto flex flex-col gap-4 font-bold text-left">
      <p>
        次の画面で、選ばれた人と「繋がり」がある5人が表示されます。<br>
        選ばれた人と「繋がりの強い順」、つまり「交流が多い順」に並び変えてください。
      </p>
      <p>
        「AさんとBさん繋がりの強さ」は <br>
        「AさんがBさんのtimesで発言した数」 + 「BさんがAさんのtimesで発言した数」 <br>
        と定義します。
      </p>
    </div>
    <div class="flex justify-center">
      <div class="grid grid-rows-2 grid-cols-2 gap-x-8 gap-y-2">
        <div></div>
        <div class="flex flex-col gap-2 m-auto">
          <label for="traQid" class="font-bold">traQ ID を入力</label>
          <input
            id="traQid"
            type="text"
            v-model="traQID"
            class="bg-white px-3 text-gray-700 border border-gray-700 rounded"
          />
        </div>
        <button
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-3 px-5 rounded fucus:outline-none focus:shadow-outline mt-3"
          @click="startWithRandom()"
        >
          ランダムな友達で遊ぶ
        </button>
        <button
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-3 px-5 rounded fucus:outline-none focus:shadow-outline mt-3"
          @click="startWithTraqID(traQID)"
        >
          ID を指定して遊ぶ
        </button>
      </div>
    </div>
  </div>
</template>

<style></style>
