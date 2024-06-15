 <template>
  <div class="big-circle-container flex justify-center items-center">
    <div class="big-circle">
      <div v-for="(image, index) in images" :key="index" class="circle-item" :style="{ backgroundImage: 'url(' + image + ')' }"></div>
    </div>
  </div>
</template>


<script>
import { ref } from 'vue';

export default {
  name: 'CircleImages',
  setup() {
    const images = ref([
      './image1.jpg',
      './image2.jpg',
      './image3.jpg',
      './image4.jpg',
      './image5.jpg',
      './image6.jpg',
    ]);

    return { images };
  },
};
</script>


<style scoped>
.big-circle {
	width: 400px;
	aspect-ratio: 1;
	border-radius: 50%;
	border: dashed 1px #777;
	margin: 50px;
	position: relative;
}

.small-circle {
	width: 80px;
	height: 80px;
	border-radius: 50%;
	background: #777;
	font-size: 32px;
	font-weight: 700;
	color: #fff;
	line-height: 80px;
	text-align: center;
	position: absolute;
	-webkit-transform: translate(-50%, -50%);
	transform: translate(-50%, -50%);
}
</style>