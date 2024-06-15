

<script setup lang="ts">
  import {ref} from 'vue'
  interface profile {
    traPID:string
    name : string
  }
  const profilesA = ref<profile[]>([{traPID:"riku_6174",name:"riku_6174"},{traPID:"riku_6174",name:"riku_6174"},{traPID:"riku_6174",name:"riku_6174"},{traPID:"riku_6174",name:"riku_6174"},{traPID:"riku_6174",name:"riku_6174"}])
  const profilesB = ref<profile[]>([{traPID:"masky5859",name:"masky"},{traPID:"masky5859",name:"masky"},{traPID:"masky5859",name:"masky"},{traPID:"masky5859",name:"masky"},{traPID:"masky5859",name:"masky"}])
  const profilesC = ref<profile[]>([{traPID:"Futadaruma",name:"フタだるま"},{traPID:"Futadaruma",name:"フタだるま"},{traPID:"Futadaruma",name:"フタだるま"},{traPID:"Futadaruma",name:"フタだるま"},{traPID:"Futadaruma",name:"フタだるま"}])
  const profilesD = ref<profile[]>([{traPID:"toki",name:"とき"},{traPID:"toki",name:"とき"},{traPID:"toki",name:"とき"},{traPID:"toki",name:"とき"},{traPID:"toki",name:"とき"}])
  const profilesE = ref<profile[]>([{traPID:"Series_205",name:"Series_205"},{traPID:"Series_205",name:"Series_205"},{traPID:"Series_205",name:"Series_205"},{traPID:"Series_205",name:"Series_205"},{traPID:"Series_205",name:"Series_205"}])
</script>

<template>
  <div>友達の友達リスト</div>
  <div class="grid grid-rows-1 grid-flow-col">
    <div>
      <div v-for="profile in profilesA" :key="profile.traPID" >
        <div class="flex bg-blue-50 justify-center">
          <img :src= "'https://q.trap.jp/api/v3/public/icon/'+profile.traPID" alt="アイコン" class="w-6 rounded-full">
          <span>{{ profile.name }}</span>
        </div>
      </div>
    </div>
    <div>
      <div v-for="profile in profilesB" :key="profile.traPID">
        <div class="flex bg-red-50 justify-center">
          <img :src= "'https://q.trap.jp/api/v3/public/icon/'+profile.traPID" alt="アイコン" class="w-6 rounded-full">
          <span>{{ profile.name }}</span>
        </div>
      </div>
    </div>
    <div>
      <div v-for="profile in profilesC" :key="profile.traPID">
        <div class="flex bg-green-50 justify-center">
          <img :src= "'https://q.trap.jp/api/v3/public/icon/'+profile.traPID" alt="アイコン" class="w-6 rounded-full">
          <span>{{ profile.name }}</span>
        </div>
      </div>
    </div>
    <div>
      <div v-for="profile in profilesD" :key="profile.traPID">
        <div class="flex bg-yellow-50 justify-center">
          <img :src= "'https://q.trap.jp/api/v3/public/icon/'+profile.traPID" alt="アイコン" class="w-6 rounded-full">
          <span>{{ profile.name }}</span>
        </div>
      </div>
    </div>
    <div>
      <div v-for="profile in profilesE" :key="profile.traPID">
        <div class="flex bg-purple-50 justify-center">
          <img :src= "'https://q.trap.jp/api/v3/public/icon/'+profile.traPID" alt="アイコン" class="w-6 rounded-full">
          <span>{{ profile.name }}</span>
        </div>
      </div>
    </div>
  </div>
    
</template>
<style scoped></style>
