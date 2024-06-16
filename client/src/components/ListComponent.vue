<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { targetId } from '.././store'

interface profile {
  traPID: string
  name: string
}
// const profilesA = ref<profile[]>([
//   { name: 'riku_6174' },
//   { traPID: 'riku_6174', name: 'riku_6174' },
//   { traPID: 'riku_6174', name: 'riku_6174' },
//   { traPID: 'riku_6174', name: 'riku_6174' },
//   { traPID: 'riku_6174', name: 'riku_6174' }
// ])
// const profilesB = ref<profile[]>([
//   { traPID: 'masky5859', name: 'masky' },
//   { traPID: 'masky5859', name: 'masky' },
//   { traPID: 'masky5859', name: 'masky' },
//   { traPID: 'masky5859', name: 'masky' },
//   { traPID: 'masky5859', name: 'masky' }
// ])
// const profilesC = ref<profile[]>([
//   { traPID: 'Futadaruma', name: 'フタだるま' },
//   { traPID: 'Futadaruma', name: 'フタだるま' },
//   { traPID: 'Futadaruma', name: 'フタだるま' },
//   { traPID: 'Futadaruma', name: 'フタだるま' },
//   { traPID: 'Futadaruma', name: 'フタだるま' }
// ])
// const profilesD = ref<profile[]>([
//   { traPID: 'toki', name: 'とき' },
//   { traPID: 'toki', name: 'とき' },
//   { traPID: 'toki', name: 'とき' },
//   { traPID: 'toki', name: 'とき' },
//   { traPID: 'toki', name: 'とき' }
// ])
// const profilesE = ref<profile[]>([
//   { traPID: 'Series_205', name: 'Series_205' },
//   { traPID: 'Series_205', name: 'Series_205' },
//   { traPID: 'Series_205', name: 'Series_205' },
//   { traPID: 'Series_205', name: 'Series_205' },
//   { traPID: 'Series_205', name: 'Series_205' }
// ])
const frendsOfFrends = ref<profile[]>([])
const profilesA = ref<profile[]>([])
const profilesB = ref<profile[]>([])
const profilesC = ref<profile[]>([])
const profilesD = ref<profile[]>([])
const profilesE = ref<profile[]>([])
onMounted(async () => {
  console.log(targetId.value)
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
})
</script>

<template>
  <div class="font-bold text-2xl">友達の友達リスト</div>
  <div class="grid grid-rows-1 grid-flow-col">
    <div class="grid bg-blue-50">
      <div v-for="profile in profilesA" :key="profile.traPID">
        <div class="flex justify-center mt-1 mb-1">
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
        <div class="flex justify-center mt-1 mb-1">
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
        <div class="flex justify-center mt-1 mb-1">
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
        <div class="flex justify-center mt-1 mb-1">
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
        <div class="flex justify-center mt-1 mb-1">
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
</template>
<style scoped></style>
