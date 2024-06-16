<template>
  <!-- TODO: convert CSS to TailwindCSS -->
  <div class="flex justify-center items-center">
    <div class="big-circle">
      <div class="center-item">
        <img
          @click="showCenterPersonInfo"
          :src="`https://q.trap.jp/api/v3/public/icon/${targetId}`"
          alt="center icon"
          class="rounded-full w-16 h-16"
        />
      </div>
      <div v-for="(image, index) in imageURLs" :key="index" class="circle-item">
        <img @click="showAroundPersonWeb(image.id)" :src="image.url" alt="around person" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { targetId, imageURLs } from '../store'

const getConnections = async (id: string) => {
  try {
    const response = await fetch(`api/users/${id}/connections`)
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
    const connections = await response.json()
    const sortedConnections = connections
      .sort((a, b) => -a.strength + b.strength)
      .map((item) => ({
        id: item.id,
        url: `https://q.trap.jp/api/v3/public/icon/${item.id}`
      }))
    imageURLs.value = sortedConnections
    targetId.value = id
    console.log(imageURLs)
  } catch (error) {
    console.error('There was a problem with the fetch operation:', error)
  }
}
const showCenterPersonInfo = async () => {
  console.log('center person info')
}
const showAroundPersonWeb = async (id: string) => {
  await getConnections(id)
}
onMounted(async () => {
  await getConnections(targetId.value)
  const items = document.querySelectorAll('.circle-item')
  const initialRadius = 100 // 最初の円の半径
  const container = document.querySelector('.big-circle') as HTMLElement
  const containerCenter = {
    x: container.offsetWidth / 2,
    y: container.offsetHeight / 2
  }
  let currentRadius = initialRadius
  let currentAngle = 0
  const itemsPerCircle = 12
  const angleOffset = (30 * Math.PI) / 180 // 各円ごとの開始角度（ラジアン）
  items.forEach((item, index) => {
    if (index % itemsPerCircle === 0 && index !== 0) {
      currentRadius += 60 // 新しい円の半径を増やす
      currentAngle = angleOffset * (index / itemsPerCircle) // 新しい円の開始角度
    }
    const angle = currentAngle + (index % itemsPerCircle) * ((2 * Math.PI) / itemsPerCircle)
    const x = Math.cos(angle) * currentRadius + containerCenter.x - item.clientWidth / 2
    const y = Math.sin(angle) * currentRadius + containerCenter.y - item.clientHeight / 2
    ;(item as HTMLElement).style.transform = `translate(${x}px, ${y}px)`
  })
})
</script>

<style scoped>
.big-circle {
  position: relative;
  width: 500px; /* コンテナの幅 */
  height: 500px; /* コンテナの高さ */
  border-radius: 50%;
}

.center-item {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.circle-item {
  position: absolute;
  width: 50px; /* 画像アイテムの幅 */
  height: 50px; /* 画像アイテムの高さ */
}
.circle-item img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}
</style>
