<template>
  <div class="hello">
    <draggable v-model="data" item-key="id" handle=".handle">
      <template #item="{ element }">
        <div class="drag-item">
          <span class="handle">ここを押せば動かせます。</span>
          {{ element.id }}
        </div>
      </template>
    </draggable>
  </div>
</template>

<script>
import draggable from 'vuedraggable'

export default {
  components: {
    draggable
  },
  data() {
    return {
      data: [
        {
          id: 1,
          content: 'テスト1'
        },
        {
          id: 2,
          content: 'テスト2'
        },
        {
          id: 3,
          content: 'テスト3'
        }
      ]
    }
  }
}
</script>

<style scoped>
.drag-item {
  background: rgb(233, 249, 255);
  margin: 10px 0;
}
</style>
