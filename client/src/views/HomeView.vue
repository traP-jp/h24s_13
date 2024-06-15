<template>
  <div class="about">
    <h1>タイトルの名前</h1>
  </div>
  <button v-on:click="home">ゲームをプレイ</button>
</template>

<script>
export default {
  methods: {
    home(){
      this.$router.push('/router')
    }
  }
}
</script>
<style>
</style>
