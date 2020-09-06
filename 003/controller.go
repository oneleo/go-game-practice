package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/oneleo/go-game-practice/env"
)

// var line int = 16
var line int = env.GetEnvAsInt("LINE", 12)

func update(screen *ebiten.Image) error {
	// 顯示一串除錯用的文字
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS=%f", ebiten.CurrentFPS()), 0, 0)
	ebitenutil.DebugPrintAt(screen, "Our first game in Ebiten!", 0, 1*line)

	// Key control
	ebitenutil.DebugPrintAt(screen, "Key information:", 0, 2*line)
	// 當「按鍵上」被按下時⋯⋯
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'UP' button.", 0, 3*line)
	}
	// 當「按鍵下」被按下時⋯⋯
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'DOWN' button.", 0, 4*line)
	}
	// 當「按鍵左」被按下時⋯⋯
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'LEFT' button.", 0, 5*line)
	}
	// 當「按鍵右」被按下時⋯⋯
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'RIGHT' button.", 0, 6*line)
	}

	// Mouse Control
	ebitenutil.DebugPrintAt(screen, "Mouse information:", 0, 7*line)
	// 從 CursorPosition() 取得 X, Y 座標
	mx, my := ebiten.CursorPosition()
	// 顯示一段「X: xx, Y: xx」格式文字
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %d, Y: %d", mx, my), 0, 8*line)
	// 當「滑鼠左鍵」被按下時⋯⋯
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'LEFT' mouse button.", 0, 9*line)
	}
	// 當「滑鼠右鍵」被按下時⋯⋯
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'RIGHT' mouse button.", 0, 10*line)
	}
	// 當「滑鼠中鍵」被按下時⋯⋯
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'MIDDLE' mouse button.", 0, 11*line)
	}

	// Gamepad Control
	ebitenutil.DebugPrintAt(screen, "Gamepad information:", 0, 12*line)
	// 兩個蘑菇頭
	axes := []float64{
		ebiten.GamepadAxis(0, 0), // 0：左邊的水平值
		ebiten.GamepadAxis(0, 1), // 1：左邊的垂直值
		ebiten.GamepadAxis(0, 2), // 2：右邊的水平值
		ebiten.GamepadAxis(0, 3)} // 3：右邊的垂直值
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Left H: %0.6f, Left V: %0.6f", axes[0], axes[1]), 0, 13*line)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Right H: %0.6f, Right V: %0.6f", axes[2], axes[3]), 0, 14*line)
	// 當遊戲手把「方向鍵上」被按下時⋯⋯
	if ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton13) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'UP' gamepad button.", 0, 15*line)
	}
	// 當遊戲手把「方向鍵下」被按下時⋯⋯
	if ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton15) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'DOWN' gamepad button.", 0, 16*line)
	}
	// 當遊戲手把「方向鍵左」被按下時⋯⋯
	if ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton16) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'LEFT' gamepad button.", 0, 17*line)
	}
	// 當遊戲手把「方向鍵右」被按下時⋯⋯
	if ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton14) {
		ebitenutil.DebugPrintAt(screen, "You're pressing the 'RIGHT' gamepad button.", 0, 18*line)
	}

	return nil
}

func main() {
	// 初始化 Ebiten，並不斷呼叫 update()
	ebiten.Run(update, 320, 240, 2, "Hello world!")
}
