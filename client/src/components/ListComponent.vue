<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { targetId } from '.././store'

interface profile {
  traPID: string
  name: string
}

const frendsOfFrends = ref<profile[]>([])
const profilesA = ref<profile[]>([])
const profilesB = ref<profile[]>([])
const profilesC = ref<profile[]>([])
const profilesD = ref<profile[]>([])
const profilesE = ref<profile[]>([])

async function fetchListData() {
  const response = await fetch(`/api/users/${targetId.value}/random?count=25`)
  const data = await response.json()
  frendsOfFrends.value = data.map((id: string) => ({
    traPID: id,
    name: id
  }))
  profilesA.value = frendsOfFrends.value.slice(0, 5)
  profilesB.value = frendsOfFrends.value.slice(5, 10)
  profilesC.value = frendsOfFrends.value.slice(10, 15)
  profilesD.value = frendsOfFrends.value.slice(15, 20)
  profilesE.value = frendsOfFrends.value.slice(20, 25)
}

const click = async (id: string) => {
  targetId.value = id
  await fetchListData()
}

onMounted(fetchListData)

watch(targetId, fetchListData)
</script>

<template>
  <div class="w-3/5 mx-auto">
    <div class="font-bold text-2xl">{{ targetId }}さんの友達の友達リスト</div>
    <div class="grid grid-rows-1 grid-flow-col">
      <div class="grid bg-blue-50">
        <div v-for="profile in profilesA" :key="profile.traPID">
          <div class="flex justify-center mt-1 mb-1 cursor-pointer" @click="click(profile.traPID)">
            <img
              :src="'https://q.trap.jp/api/v3/public/icon/' + profile.traPID"
              alt="アイコン"
              class="w-6 rounded-full"
            />
            <span>{{ profile.name }}</span>
          </div>
        </div>
      </div>
      <div class="grid bg-red-50">
        <div v-for="profile in profilesB" :key="profile.traPID">
          <div class="flex justify-center mt-1 mb-1 cursor-pointer" @click="click(profile.traPID)">
            <img
              :src="'https://q.trap.jp/api/v3/public/icon/' + profile.traPID"
              alt="アイコン"
              class="w-6 rounded-full"
            />
            <span>{{ profile.name }}</span>
          </div>
        </div>
      </div>
      <div class="grid bg-green-50">
        <div v-for="profile in profilesC" :key="profile.traPID">
          <div class="flex justify-center mt-1 mb-1 cursor-pointer" @click="click(profile.traPID)">
            <img
              :src="'https://q.trap.jp/api/v3/public/icon/' + profile.traPID"
              alt="アイコン"
              class="w-6 rounded-full"
            />
            <span>{{ profile.name }}</span>
          </div>
        </div>
      </div>
      <div class="grid bg-yellow-50">
        <div v-for="profile in profilesD" :key="profile.traPID">
          <div class="flex justify-center mt-1 mb-1 cursor-pointer" @click="click(profile.traPID)">
            <img
              :src="'https://q.trap.jp/api/v3/public/icon/' + profile.traPID"
              alt="アイコン"
              class="w-6 rounded-full"
            />
            <span>{{ profile.name }}</span>
          </div>
        </div>
      </div>
      <div class="grid bg-purple-50">
        <div v-for="profile in profilesE" :key="profile.traPID">
          <div class="flex justify-center mt-1 mb-1 cursor-pointer" @click="click(profile.traPID)">
            <img
              :src="'https://q.trap.jp/api/v3/public/icon/' + profile.traPID"
              alt="アイコン"
              class="w-6 rounded-full"
            />
            <span>{{ profile.name }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped></style>
