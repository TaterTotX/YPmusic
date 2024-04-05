package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(play_status string) string {

	if play_status == "暂停" {
		return fmt.Sprintf("播放")
	} else {
		return fmt.Sprintf("暂停")
	}

}

// 定义全局变量
var musicFolderPath string

// 获取前端返回的音乐文件路径
func (a *App) Get_music_path(music_path string) string {
	if len(music_path) > 5 {
		SaveInfo(music_path)
	}

	musicFolderPath, _ = ReadInfo()
	return musicFolderPath
}
