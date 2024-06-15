<script setup lang="ts">

</script>

<template>
  <br><br><br>
  <h1 class="text-center text-2xl font-bold">ゲームの説明</h1>
  <br><br>
  <p class="max-w-2xl mx-auto">
   traQidを入力してスタートボタンを押すとゲームが始まります。（入力がない場合はランダムで選ばれます）ゲーム画面では選んだ人（以下hogeさんと呼ぶ）とランダムで５人が表示されるので、５人をhogeさんと繋がりの強い順に並び変えてください。
  </p>
  <p class="max-w-2xl mx-auto">
    ただし“A さんと B さん繋がりの強さ” := “A さんが B さんの times で発言した数” + “B さんが A さんの times で発言した数”
  </p>
  <br><br>
  <div class="flex justify-center">
    <label for="traQid" class="">traQIDを入力してください:</label>
    <input id="traQid" type="text" class="bg-white px-3 text-gray-700 border border-gray-700 rounded">
  </div>
  
  <br>
  <div class="flex justify-center">
    <button
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-3 px-5 rounded fucus:outline-none focus:shadow-outline mt-3"
    >
      ゲームスタート
    </button>
  </div>
  
</template>

<style>
  
</style>
