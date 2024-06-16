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
        <img
          v-if="image.url != ''"
          @click="showAroundPersonWeb(image.id)"
          :src="image.url"
          alt="around person"
        />
      </div>
    </div>
  </div>

  <div
    v-if="isModalOpen"
    class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full"
    @click.self="isModalOpen = false"
  >
    <div class="relative top-20 mx-auto p-5 border shadow-lg rounded-md bg-white w-3/4">
      <div class="mt-3 text-center">
        <h3 class="text-lg leading-6 font-medium text-gray-900">{{ targetId }}さんの情報</h3>
        <div class="max-w-4xl mx-auto my-4 overflow-hidden shadow-md sm:rounded-lg">
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
            @click="isModalOpen = false"
          >
            閉じる
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { targetId, imageURLs } from '../store'

const isModalOpen = ref(false)

const getConnections = async (id: string) => {
  try {
    const response = await fetch(`api/users/${id}/connections`)
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
    const connections = await response.json()
    let sortedConnections = connections
      .sort((a, b) => -a.strength + b.strength)
      .map((item) => ({
        id: item.id,
        url: `https://q.trap.jp/api/v3/public/icon/${item.id}`
      }))
      .slice(0, 36)
    const requiredConnections = 36
    const currentLength = sortedConnections.length
    if (currentLength < requiredConnections) {
      const emptyConnections = Array.from({ length: requiredConnections - currentLength }, () => ({
        id: '',
        url: ''
      }))
      sortedConnections = sortedConnections.concat(emptyConnections)
    }

    imageURLs.value = sortedConnections
    targetId.value = id
  } catch (error) {
    console.error('There was a problem with the fetch operation:', error)
  }
}
const showCenterPersonInfo = async () => {
  isModalOpen.value = true
  await getUserGroups(targetId.value)
}
const showAroundPersonWeb = async (id: string) => {
  await getConnections(id)
}
const userGroups = ref<User[]>([])
const getUserGroups = async (id: string) => {
  try {
    const response = await fetch(`/api/users/${id}`)
    if (!response.ok) {
      alert('適切なデータが見つかりません')
      throw new Error('Network response was not ok')
    }
    const users = await response.json()
    userGroups.value = users.groups
  } catch (error) {
    console.error(error)
  }
}
onMounted(async () => {
  await getConnections(targetId.value)
  await getUserGroups(targetId.value)
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
