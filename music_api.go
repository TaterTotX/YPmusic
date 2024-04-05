package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func serveMusic(w http.ResponseWriter, r *http.Request) {

	// 从请求中获取歌曲名称
	songName := r.URL.Query().Get("song")
	if songName == "" {
		http.Error(w, "Song name not provided", http.StatusBadRequest)
		return
	}

	// 构建完整的音乐文件路径
	musicFilePath := filepath.Join(musicFolderPath, songName)

	// 打开音乐文件
	musicFile, err := os.Open(musicFilePath)
	if err != nil {
		log.Println("Error opening music file:", err, musicFilePath)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer musicFile.Close()

	// 获取音乐文件的大小
	fileInfo, err := musicFile.Stat()
	if err != nil {
		log.Println("Error getting music file info:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fileSize := fileInfo.Size()

	// 设置响应头
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(musicFilePath))
	w.Header().Set("Content-Length", strconv.FormatInt(fileSize, 10))
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 将音乐文件作为流发送给客户端
	_, err = io.Copy(w, musicFile)
	if err != nil {
		log.Println("Error copying music file to response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func serveMusicList(w http.ResponseWriter, r *http.Request) {

	log.Println("serveMusicList", musicFolderPath)
	// 读取音乐文件夹
	entries, err := os.ReadDir(musicFolderPath)
	if err != nil {
		log.Println("Error reading music folder:", err, musicFolderPath)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 创建一个用于存储音乐文件名的切片
	var musicFiles []string

	// 遍历文件夹条目，收集音乐文件
	for _, entry := range entries {
		if !entry.IsDir() {
			musicFiles = append(musicFiles, entry.Name())
		}
	}

	// 将音乐文件列表作为JSON响应返回
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(musicFiles)
}

func music_api() {
	musicFolderPath, _ = ReadInfo()
	http.HandleFunc("/music/list", serveMusicList) // 添加新的端点
	http.HandleFunc("/music", serveMusic)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
