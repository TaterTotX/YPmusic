<script setup>

</script>

<template>
  <button class="left_next_button">
    <svg viewBox="0 0 384 512" class="svgIcon">
      <path
          d="M25.1 239l135.9-136c9.4-9.4 24.6-9.4 33.9 0l22.6 22.6c9.4 9.4 9.4 24.6 0 33.9L109.9 256l107.6 107.6c9.4 9.4 9.4 24.6 0 33.9l-22.6 22.6c-9.4 9.4-24.6 9.4-33.9 0l-135.9-136c-9.4-9.4-9.4-24.6 0-33.9z"
      ></path>
    </svg>
  </button>

</template>

<style scoped>
.left_next_button {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  background-color: rgb(20, 20, 20);
  border: none;
  font-weight: 600;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0px 0px 0px 4px rgb(43, 50, 52);
  cursor: pointer;
  transition-duration: 0.3s;
  overflow: hidden;
  position: relative;
}

.svgIcon {
  width: 12px;
  transition-duration: 0.3s;
}

.svgIcon path {
  fill: white;
}

.left_next_button:hover {

  width: 140px;
  border-radius: 50px;
  transition-duration: 0.3s;
  background-color: rgb(181, 160, 255);
  align-items: center;
}

.left_next_button:hover .svgIcon {
  /* width: 20px; */
  transition-duration: 0.3s;
  transform: translatex(-200%);


}

.left_next_button::before {
  position: absolute;
  bottom: -20px;
  content: "上一首";
  color: white;
  /* transition-duration: .3s; */
  font-size: 0px;
}

.left_next_button:hover::before {
  font-size: 13px;
  opacity: 1;
  bottom: unset;
  /* transform: translateY(-30px); */
  transition-duration: 0.3s;
}


</style>