<template>
  <div class="grid place-items-center">
    <div class="w-2/5 grid place-items-center">
      <h2 class="font-bold text-4xl my-4">正解</h2>
    </div>
    <div class="w-2/5 grid grid-cols-5 gap-4 place-items-center my-4">
      <div v-for="user in answerUsers" :key="user.id" class="flex justify-center">
        <img
          :src="`https://q.trap.jp/api/v3/public/icon/${user.id}`"
          alt="Profile Icon"
          class="rounded-full w-16 h-16"
        />
      </div>
    </div>
    <div class="w-2/5 grid place-items-center">
      <h2 class="font-bold text-4xl my-4">あなたの答え</h2>
    </div>
    <div class="w-2/5 grid grid-cols-5 gap-4 place-items-center my-4">
      <div v-for="user in changeableUsers" :key="user.id" class="flex justify-center">
        <img
          :src="`https://q.trap.jp/api/v3/public/icon/${user.id}`"
          alt="Profile Icon"
          class="rounded-full w-16 h-16"
        />
      </div>
    </div>
    <button
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-3 px-5 rounded focus:outline-none focus:shadow-outline my-4"
      @click="showConnection"
    >
      {{ targetId }}さんの繋がりを見る
    </button>
  </div>
</template>
<script setup lang="ts">
import { onMounted } from 'vue'
import { changeableUsers, answerUsers, targetId } from '.././store'
import { useRouter } from 'vue-router'

const router = useRouter()

const showConnection = () => {
  router.push('/connection')
}

const getAnswerUsers = async (id: string) => {
  try {
    const answersQuery = changeableUsers.value.map((user) => user.id).join(',')
    const res = await fetch(`/api/quiz/answer?id=${id}&answers=${answersQuery}`)
    const userIds = await res.json()
    answerUsers.value = userIds.map((userId: string) => ({ id: userId }))
  } catch (error) {
    console.error(error)
  }
}

onMounted(async () => {
  await getAnswerUsers(targetId.value)
})
</script>

<style scoped></style>
