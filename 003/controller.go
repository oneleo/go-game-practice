package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var line int = 16

func update(screen *ebiten.Image) error {
	// 顯示一串除錯用的文字
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS=%f", ebiten.CurrentFPS()), 0, 0)
	ebitenutil.DebugPrintAt(screen, "Our first game in Ebiten!", 0, 1*line)

	// 當「按鍵上」被按下時⋯⋯
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'UP' button.", 0, 2*line)
	}

	// 當「按鍵下」被按下時⋯⋯
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'DOWN' button.", 0, 3*line)
	}

	// 當「按鍵左」被按下時⋯⋯
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'LEFT' button.", 0, 4*line)
	}

	// 當「按鍵右」被按下時⋯⋯
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'RIGHT' button.", 0, 5*line)
	}

	return nil
}

func main() {
	// 初始化 Ebiten，並不斷呼叫 update()
	ebiten.Run(update, 320, 240, 2, "Hello world!")
}
