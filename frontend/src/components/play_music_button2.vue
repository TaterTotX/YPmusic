<script setup>


</script>

<template>
  <div class="container">
    <label>
      <input class="play-btn" type="checkbox">
      <div class="play-icon"  ></div>
      <div class="pause-icon"></div>
    </label>
  </div>
</template>

<style scoped>
.container {
  width: 80px;
  height: 80px;
  position: relative;
  border-radius: 50%;
  left: -35px; /* 向左移动20像素 */

}

.play-btn {
  position: absolute;
  appearance: none;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  background: conic-gradient(rgb(43, 50, 52), rgb(43, 50, 52));
  cursor: pointer;
  outline: none;
}

.play-btn::before {
  content: "";
  position: absolute;
  width: 93%;
  height: 93%;
  background-color: #000;
  border-radius: 50%;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}

.play-btn:checked {
  animation: borderAnimate 700ms ease-in-out 1;
  animation-fill-mode: forwards;
}

@keyframes borderAnimate {
  0% {
    transform: rotate(0);
    background: conic-gradient(rgb(181, 160, 255), transparent 20%);
  }

  80% {
    background: conic-gradient(rgb(181, 160, 255), transparent 90%);
  }

  100% {
    transform: rotate(360deg);
    background: conic-gradient(rgb(181, 160, 255), rgb(181, 160, 255));
  }
}

.play-icon {
  position: absolute;
  width: 20px;
  height: 20px;
  left: 110%;
  top: 50%;
  background-color: #ffffff;
  transform: translate(-60%, -50%) rotate(90deg);
  clip-path: polygon(50% 15%, 0% 100%, 100% 100%);
  transition: all 400ms ease-in-out;
  cursor: pointer;
}

.play-btn:checked + .play-icon {
  clip-path: polygon(0 100%, 0% 100%, 100% 100%);
}

.pause-icon {
  position: absolute;
  width: 20px;
  height: 20px;
  left: 100%;
  top: 50%;
  transform: translate(-50%, -50%);
  cursor: pointer;
}

.pause-icon::before {
  content: "";
  position: absolute;
  width: 0%;
  height: 100%;
  background-color: #ffffff;
  left: 0;
}

.pause-icon::after {
  content: "";
  position: absolute;
  width: 0;
  height: 100%;
  background-color: #ffffff;
  right: 0;
}

.play-btn:checked ~ .pause-icon::before {
  animation: reveal 300ms ease-in-out 350ms 1;
  animation-fill-mode: forwards;
}

.play-btn:checked ~ .pause-icon::after {
  animation: reveal 300ms ease-in-out 600ms 1;
  animation-fill-mode: forwards;
}

@keyframes reveal {
  0% {
    width: 0;
  }

  100% {
    width: 35%;
  }
}
</style>