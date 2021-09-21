package main

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
// NewApp 创建一个新的 App 应用程序
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
// startup 在应用程序启动时调用
func (b *App) startup(ctx context.Context) {
	// Perform your setup here
	// 在这里执行初始化设置
	b.ctx = ctx
}

// shutdown is called at application termination
// 在应用程序终止时被调用
func (b *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	// 在此处做一些资源释放的操作
}
