package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	go music_api()

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "playmusic",
		Width:  700,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			// WebviewIsTransparent: 设置Webview是否透明
			WebviewIsTransparent: false,
			// WindowIsTranslucent: 设置窗口是否半透明
			WindowIsTranslucent: false,
			// BackdropType: 设置窗口的背景类型，这里使用的是Mica效果，Mica是一种材料效果，提供一些透明度
			BackdropType: windows.Mica,
			// DisableWindowIcon: 设置是否禁用窗口图标
			DisableWindowIcon: true,
			// DisableFramelessWindowDecorations: 设置是否禁用无框窗口装饰（会移除窗口边框和标题栏）
			DisableFramelessWindowDecorations: true,
			// WebviewUserDataPath: Webview的用户数据路径，可以用来存储cookies，缓存等信息
			WebviewUserDataPath: "",
			// WebviewBrowserPath: Webview的浏览器路径，可以指定自定义的浏览器执行文件路径
			WebviewBrowserPath: "",
			// Theme: 设置应用的主题，这里使用系统默认主题
			Theme: windows.SystemDefault,
		},

		Mac: &mac.Options{
			TitleBar: mac.TitleBarHiddenInset(),

			// 应用程序的外观设置，这里设置为深色系
			Appearance: mac.NSAppearanceNameDarkAqua,

			// 网页视图是否透明
			WebviewIsTransparent: true,
			// 窗口是否半透明
			WindowIsTranslucent: true,
			// 关于窗口的信息
			About: &mac.AboutInfo{
				Title:   "bossapp",  // 应用程序的名称
				Message: "土豆© 2024", // 关于窗口中显示的信息，例如版权信息

			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
