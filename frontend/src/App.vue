<script setup>
import {onMounted, reactive} from 'vue';
import {Greet,Get_music_path} from '../wailsjs/go/main/App'
import Left_next_button from "./components/left_next_button.vue";
import Right_next_button from "./components/right_next_button.vue";
import Logo from "./components/logo.vue";
import Play_music_button2 from "./components/play_music_button2.vue";
import Setting_button from "./components/setting_button.vue";
import Setting_page from "./components/setting_page.vue";
import {LogPrint} from "../wailsjs/runtime/runtime.js";

const data = reactive({
  music:"",
  play_status: "暂停",
  audio: null,
  now_music: "请配置音乐路径哦❤️",
  num:0,
  end_status:0,
  showSetting :true,
  music_path:"",
  music_mode:"",
  current_index:0,
  is_first:0
});



// 在组件挂载后获取音乐列表
onMounted(() => {
  fetchMusicList();
});

// 定义一个函数来从服务器获取音乐列表
async function fetchMusicList() {
    const response = await fetch('http://localhost:8080/music/list');
    const musicList = await response.json();
    console.log(response.json())
    data.music = musicList;
    data.now_music=musicList[0]
}

//定义上一首歌的按钮
function up_music() {
  if (data.num > 0) {
    data.num -= 1;
  } else {
    // 如果已经在第一首，就跳到最后一首
    data.num = data.music.length - 1;
  }
  data.now_music = data.music[data.num];
  releaseMusic()
  stopMusic()
  if (data.play_status==="播放"&&data.music_mode==="") {
    playMusic(data.now_music, "切换");
  }else if (data.play_status==="播放"&&data.music_mode==="单曲") {
    playMusicLoop(data.now_music, "切换")
  }else if (data.play_status==="播放"&&data.music_mode==="随机") {
    playMusicRandom2("切换")
  }else if (data.play_status==="播放"&&data.music_mode==="循环") {
    playMusic_circulate("切换")
  }

}

// 定义下一首歌的按钮
function down_music() {
  if (data.num < data.music.length - 1) {
    data.num += 1;
  } else {
    // 如果已经在最后一首，就跳到第一首
    data.num = 0;
  }
  data.now_music = data.music[data.num];
  releaseMusic()
  stopMusic()
  if (data.play_status==="播放"&&data.music_mode==="") {
    playMusic(data.now_music, "切换");
  }else if (data.play_status==="播放"&&data.music_mode==="单曲") {
    playMusicLoop(data.now_music, "切换")
  }else if (data.play_status==="播放"&&data.music_mode==="随机") {
    playMusicRandom2("切换")
  }else if (data.play_status==="播放"&&data.music_mode==="循环") {
    playMusic_circulate("切换")
  }
}


//连接go函数
function greet() {
  Get_music_path(data.music_path).then(
      music_path=>{
        data.music_path=music_path
      }
  )


  Greet(data.play_status).then(result => {
    data.play_status=result
    if (data.play_status==="播放" && data.music_mode===""){
      playMusic(data.now_music,"播放");
    }else if(data.play_status==="播放" && data.music_mode==="单曲"){
      playMusicLoop(data.now_music,"播放")
    }else if(data.play_status==="播放" && data.music_mode==="随机"){
      playMusicRandom2("播放")
    }else if(data.play_status==="播放" && data.music_mode==="循环"){
      playMusic_circulate("播放")
    }



    else {
      stopMusic();
    }
  })

}


// 定义播放音乐的函数
function playMusic(music_name,tag) {
  if (data.music.length===0){
    fetchMusicList()
  }
  if (!data.audio ||tag==="切换") {
    data.audio = new Audio();

    data.audio.src = 'http://localhost:8080/music?song='+music_name;

    console.log(data.audio.src,"获取到的音乐路径为")
    data.audio.onloadedmetadata = function() {
      data.audio.play();

    };

    data.audio.onended = function () {
      data.end_status+=1
      data.play_status="暂停"
    };
  }else {
    data.audio.play();
  }
}





// 定义停止音乐播放的函数
function stopMusic() {
  if (data.audio) {
    data.audio.pause();
  }
}


//清空音乐播放缓存
function releaseMusic(){
  data.audio.src = ''; // 清除音频源
  data.audio = null; // 释放audio对象

}

//展示设置页面
function ShowSet() {
  data.showSetting = !data.showSetting;
}





// 获取音乐路径
const handleInputChange = (value) => {
  if (value.length > 4){
    data.music_path = value;
    data.now_music="";
    data.music=""
    stopMusic()
    releaseMusic()
    data.end_status+=1;
    data.play_status="暂停"


  }


};


// 获取循环/单曲/随机信息
const music_mode = (value) => {
    data.music_mode = value;
};



//单曲循环方法
function playMusicLoop(music_name, tag) {
  if (data.music.length === 0) {
    fetchMusicList();
  }
  // 如果audio对象已经存在，且标签不是“切换”，则直接播放
  if (data.audio && tag !== "切换") {
    data.audio.play();
  } else {
    data.audio = new Audio();
    data.audio.src = 'http://localhost:8080/music?song=' + music_name;
    data.audio.onloadedmetadata = function() {
      data.audio.play();
    };


    data.audio.onended = function() {
      data.audio.play();
    };
  }
}



//随机播放方法
function playMusicRandom2(tag) {
  if (data.music.length === 0) {
    fetchMusicList();
  }

  function playRandomSong() {
    const randomIndex = Math.floor(Math.random() * data.music.length);
    const randomMusicName = data.music[randomIndex];
    data.audio.src = 'http://localhost:8080/music?song=' + randomMusicName;
    data.now_music=randomMusicName
    data.audio.onloadedmetadata = function() {
      data.audio.play();
    };
  }

  if (!data.audio || tag === "切换") {
    data.audio = new Audio();
    playRandomSong();
  } else {
    data.audio.play();
  }

  data.audio.onended = playRandomSong;
}




//循环方法
function playMusic_circulate(tag) {
  if (data.music.length === 0) {
    fetchMusicList();
  }

  function playNextSong() {
    data.current_index = (data.current_index + 1) % data.music.length;
    const nextMusicName = data.music[data.current_index];
    data.audio.src = 'http://localhost:8080/music?song=' + nextMusicName;
    data.audio.onloadedmetadata = function() {
      data.audio.play();
    };
  }

  if (!data.audio || tag === "切换") {
    data.current_index = data.music.indexOf(data.now_music)-1;
    data.audio = new Audio();
    playNextSong();
  } else {
    data.audio.play();
  }

  data.audio.onended = playNextSong;
}

</script>
<template>
  <div id="result" class="result animate__animated animate__rubberBand" :key="data.num" style="position: relative;top: 30px;">{{ data.now_music }}{{data.music_mode}}</div>
  <keep-alive>
  <logo v-if="data.showSetting"/>
  <setting_page  v-else @input-change="handleInputChange"  @music-mode="music_mode" />
  </keep-alive>
  <div style="display: flex; justify-content: center; align-items: center; margin-top: 50px;">
    <div>
      <left_next_button @click="up_music"/>
    </div>
    <div style="margin: 0 130px;">
      <play_music_button2 :key="data.end_status" @click="greet" />
    </div>
    <div>
      <right_next_button @click="down_music"/>
    </div>
  </div>
  <setting_button @click="ShowSet" style="position: fixed; top: 15px; right: 15px; " />
</template>




<style>
* {
  -webkit-user-select: none;  /* Safari */
  -moz-user-select: none;     /* Firefox */
  -ms-user-select: none;      /* IE10+/Edge */
  user-select: none;          /* Standard */
}
</style>