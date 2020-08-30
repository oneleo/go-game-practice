package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	// 顯示一串除錯用的文字
	ebitenutil.DebugPrint(screen, "Our first game in Ebiten!")

	return nil
}

func main() {
	// 初始化 Ebiten，並不斷呼叫 update()
	ebiten.Run(update, 320, 240, 2, "Hello world!")
}
