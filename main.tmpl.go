package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	// 创建一个App结构体实例
	app := NewApp()

	// Create application with options
	// 使用选项创建应用
	err := wails.Run(&options.App{
		Title:             "{{.ProjectName}}",
		Width:             1600,
		Height:            1000,
		MinWidth:          1600,
		MinHeight:         1000,
		MaxWidth:          2000,
		MaxHeight:         1200,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         true,
		StartHidden:       false,
		HideWindowOnClose: false,
		RGBA:              &options.RGBA{255,255,255,0},
		Assets:            assets,
		LogLevel:   logger.DEBUG,
		OnStartup:  app.startup,
		OnDomReady: app.domReady,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		// Windows平台特定选项
		Windows: &windows.Options{
			WebviewIsTransparent:          true,
			WindowIsTranslucent: true,
			DisableWindowIcon:             false,
		},
		// Mac: &mac.Options{
		// 	WebviewIsTransparent:          true,
		// 	WindowBackgroundIsTranslucent: true,
		// 	TitleBar:                      mac.TitleBarHiddenInset(),
		// },
	})
	
	if err != nil {
		log.Fatal(err)
	}
}
