<script setup>

</script>

<template>
  <button class="set_button">
  </button>

</template>

<style scoped>
.set_button {
  width: 30px;
  height: 30px;
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





.set_button:hover {
  width: 140px;
  border-radius: 50px;
  transition-duration: 0.3s;
  background-color: rgb(181, 160, 255);
  align-items: center;
}

.set_button:hover .svgIcon {
  /* width: 20px; */
  transition-duration: 0.3s;
  transform: translatex(200%);
}

.set_button::before {
  position: absolute;
  bottom: -20px;
  content: "设置";
  color: white;
  /* transition-duration: .3s; */
  font-size: 0px;
}

.set_button:hover::before {
  font-size: 13px;
  opacity: 1;
  bottom: unset;
  /* transform: translateY(-30px); */
  transition-duration: 0.3s;
}


</style>