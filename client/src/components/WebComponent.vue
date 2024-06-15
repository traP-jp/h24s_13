<template>
  <div class="big-circle">
    <div v-for="(image, index) in images" :key="index" class="circle-item">
      <img :src="image" alt="circle image" />
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'

export default {
  name: 'CircleImages',
  setup() {
    const images = ref([
      `https://q.trap.jp/api/v3/public/icon/masky5859`,
      `https://q.trap.jp/api/v3/public/icon/masky5859`,
      `https://q.trap.jp/api/v3/public/icon/masky5859`,
      `https://q.trap.jp/api/v3/public/icon/masky5859`,
      `https://q.trap.jp/api/v3/public/icon/masky5859`,
      `https://q.trap.jp/api/v3/public/icon/masky5859`
    ])

    onMounted(() => {
      const items = document.querySelectorAll('.circle-item')
      const radius = 100 // 中心からの距離
      const container = document.querySelector('.big-circle')
      const containerCenter = {
        x: container.offsetWidth / 2,
        y: container.offsetHeight / 2
      }

      items.forEach((item, index) => {
        const angle = (index / items.length) * (2 * Math.PI) // 画像ごとの角度
        const x = Math.cos(angle) * radius + containerCenter.x - item.offsetWidth / 2
        const y = Math.sin(angle) * radius + containerCenter.y - item.offsetHeight / 2
        item.style.transform = `translate(${x}px, ${y}px)`
      })
    })

    return { images }
  }
}
</script>

<style scoped>
.big-circle {
  position: relative;
  width: 300px; /* コンテナの幅 */
  height: 300px; /* コンテナの高さ */
  border: 1px solid #ccc; /* 視覚的な補助のためのボーダー */
  border-radius: 50%;
}

.circle-item {
  position: absolute;
  width: 50px; /* 画像アイテムの幅 */
  height: 50px; /* 画像アイテムの高さ */
  transition: transform 0.3s;
}

.circle-item img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}
</style>
