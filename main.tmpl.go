package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {

	// Create application with options
	app := NewApp()

	err := wails.Run(&options.App{
		Title:             "{{.ProjectName}}",
		Width:             1600,
		Height:            1000,
		MinWidth:          400,
		MinHeight:         400,
		MaxWidth:          2000,
		MaxHeight:         1500,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		RGBA:              0x000000FF,
		Assets:            assets,
		Windows: &windows.Options{
			WebviewIsTransparent:          false,
			WindowBackgroundIsTranslucent: false,
			DisableWindowIcon:             false,
		},
		// Mac: &mac.Options{
		// 	WebviewIsTransparent:          true,
		// 	WindowBackgroundIsTranslucent: true,
		// 	TitleBar:                      mac.TitleBarHiddenInset(),
		// },
		LogLevel:   logger.DEBUG,
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
