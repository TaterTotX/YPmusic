<script setup>

</script>

<template>
  <img draggable="false" id="music_logo" alt="Wails logo"  src="../assets/images/logo.png"  />

</template>

<style scoped>
#music_logo {
  display: block; /* 块级元素显示方式，使得元素独占一行 */
  width: 500px; /* 固定容器宽度 */
  height: 300px; /* 固定容器高度 */
  margin: 70px auto 0; /* 上外边距5像素，左右自动，下外边距0，使元素在父元素中水平居中同时向下移动 */
  border-radius: 20px; /* 设置元素的圆角，圆角的半径为20像素 */
  object-fit: cover; /* 控制图片如何填充容器，覆盖整个容器 */
  overflow: hidden; /* 隐藏超出容器大小的图片部分 */
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.7); /* 添加阴影效果 */
}
</style>