

<script setup lang="ts">
  import {ref} from 'vue'
  interface profile {
    traPID:string
    name : string
  }
  const profiles = ref<profile[]>
  ([{traPID:"riku_6174",name:"riku_6174"},{traPID:"masky5859",name:"masky"},
  {traPID:"Futadaruma",name:"フタだるま"},{traPID:"toki",name:"とき"},
  {traPID:"Series_205",name:"Series_205"},{traPID:"riku_6174",name:"riku_6174"},{traPID:"masky5859",name:"masky"},
  {traPID:"Futadaruma",name:"フタだるま"},{traPID:"toki",name:"とき"},
  {traPID:"Series_205",name:"Series_205"},{traPID:"riku_6174",name:"riku_6174"},{traPID:"masky5859",name:"masky"},
  {traPID:"Futadaruma",name:"フタだるま"},{traPID:"toki",name:"とき"},
  {traPID:"Series_205",name:"Series_205"},{traPID:"riku_6174",name:"riku_6174"},{traPID:"masky5859",name:"masky"},
  {traPID:"Futadaruma",name:"フタだるま"},{traPID:"toki",name:"とき"},
  {traPID:"Series_205",name:"Series_205"},{traPID:"riku_6174",name:"riku_6174"},{traPID:"masky5859",name:"masky"},
  {traPID:"Futadaruma",name:"フタだるま"},{traPID:"toki",name:"とき"},
  {traPID:"Series_205",name:"Series_205"}])
</script>

<template>
  <div>友達の友達リスト</div>
  <div class="grid grid-rows-5 grid-flow-col">
    <div v-for="profile in profiles" :key="profile.traPID">
        <div class="flex bg-blue-50">
          <img :src= "'https://q.trap.jp/api/v3/public/icon/'+profile.traPID" alt="アイコン" class="w-6 rounded-full">
          <span>{{ profile.name }}</span>
        </div>
      </div>
  </div>
</template>
<style scoped></style>
