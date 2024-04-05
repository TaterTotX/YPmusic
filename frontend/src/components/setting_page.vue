<script setup>
import { defineEmits } from 'vue'



// 使用 defineEmits 来定义组件可以触发的自定义事件
const emit = defineEmits(['input-change', 'music-mode']);

function handleInput(event) {
  emit('input-change', event.target.value)

}


function music_mode(event) {
  emit('music-mode', event.target.value)
}


</script>

<template>
  <div class="image-container">
    <img draggable="false" id="music_logo" class="animate__animated animate__rubberBand" />

    <div class="container overlay-text">
      <div class="pane">
        <label class="label">
          <span>循环</span>
          <input id="left" class="input" name="radio" type="radio" :value="'循环'"  @input="music_mode">
        </label>
        <label class="label">
          <span>随机</span>
          <input id="middle" class="input" checked="checked" name="radio" type="radio" :value="'随机'" @input="music_mode">
        </label>
        <label class="label">
          <span>单曲</span>
          <input id="right" class="input" name="radio" type="radio" :value="'单曲'" @input="music_mode">
        </label>
        <span class="selection"></span>
      </div>

      <input   class="music_path" placeholder="请拖入音乐文件夹"  @input="handleInput" >


    </div>



  </div>


</template>

<style scoped>
.image-container {
  position: relative; /* 使得子元素可以绝对定位相对于该容器 */
  width: 500px; /* 固定容器宽度 */
  height: 300px; /* 固定容器高度 */
  margin: 70px auto 0; /* 上外边距70像素，左右自动，下外边距0 */
}

#music_logo {
  display: block;
  width: 100%; /* 容器的宽度 */
  height: 100%; /* 容器的高度 */
  border-radius: 20px;
  object-fit: cover;
  overflow: hidden;
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.7);
}

.overlay-text {
  position: absolute; /* 绝对定位 */
  top: 40%; /* 垂直居中 */
  left: 50%; /* 水平居中 */
  transform: translate(-50%, -50%); /* 偏移自身宽高的一半来确保居中 */
  color: white; /* 文字颜色 */
  text-align: center; /* 文字居中 */
  font-size: 24px; /* 文字大小 */
  /* 这里可以添加更多样式，如字体、阴影等 */
}










.container {
  transform-style: preserve-3d;
  perspective: 1000px;
}

.pane {
  outline: 2px solid #191e1e;
  box-shadow: 0 0 10px #191e1e, inset 0 0 10px #191e1e;
  height: 1cm;
  width: 4.5cm;
  border-radius: 5px;
  position: relative;
  overflow: hidden;
  transition: 0.7s ease;
}

.input {
  display: none;
}

.label {
  height: 1cm;
  width: 1.5cm;
  float: left;
  font-weight: 600;
  letter-spacing: -1px;
  font-size: 14px;
  padding: 0px;
  position: relative;
  z-index: 1;
  color: #ffffff;
  text-align: center;
  padding-top: 10px;
}

.selection {
  display: none;
  position: absolute;
  height: 1cm;
  width: calc(4.5cm / 3);
  z-index: 0;
  left: 0;
  top: 0;
  box-shadow: 0 0 10px #191e1e;
  transition: .15s ease;
}

.label:has(input:checked) {
  color: #ffffff;
}

.pane:has(.label:nth-child(1):hover) {
  transform: rotateY(-30deg);
}

.pane:has(.label:nth-child(3):hover) {
  transform: rotateY(35deg);
}

.label:has(input:checked) ~ .selection {
  background-color: #191e1e;
  display: inline-block;
}

.label:nth-child(1):has(input:checked) ~ .selection {
  transform: translateX(calc(4.5cm * 0/3));
}

.label:nth-child(2):has(input:checked) ~ .selection {
  transform: translateX(calc(4.5cm * 1/3));
}

.label:nth-child(3):has(input:checked) ~ .selection {
  transform: translateX(calc(4.5cm * 2/3));
}





input[class*="music_path"] {
  height: 30px;
  width: 160px;
  border-radius: 5px;
  border: 2px solid black; /* 默认状态下透明边框 */
  background-color: transparent; /* 设置透明背景 */
  position: relative;
  top: 40px;
  transition: border-color 0.3s; /* 添加过渡效果使变化更平滑 */
}

input[class*="music_path"]:focus {
  border-color: black; /* 获得焦点时边框变为黑色 */
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.5); /* 添加黑色的光圈效果 */
  outline: none; /* 移除默认的焦点轮廓线 */
}


</style>