package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	s "github.com/oneleo/go-game-practice/screen"
	"github.com/oneleo/go-game-practice/transform"
)

// // Position is
// type Position struct {
// 	tx float64
// 	ty float64
// }

// // Scale is
// type Scale struct {
// 	width  int
// 	height int
// }

func update(screen *ebiten.Image) error {
	// 先幫畫布填上 #FF0000 顏色
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})

	// 顯示一串除錯用的文字
	ebitenutil.DebugPrint(screen, "Our first game in Ebiten!")

	if err := s.DrawSquare(screen, &transform.Position{Tx: 64, Ty: 64}, &transform.Scale{Width: 16, Height: 16}, color.White); err != nil {
		log.Fatal(err)
	}

	if err := s.DrawSquare(screen, &transform.Position{Tx: 64, Ty: 128}, &transform.Scale{Width: 16, Height: 32}, color.Black); err != nil {
		log.Fatal(err)
	}

	if err := s.DrawSquare(screen, &transform.Position{Tx: 128, Ty: 64}, &transform.Scale{Width: 32, Height: 16}, color.NRGBA{0xff, 0xff, 0x00, 0xff}); err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	// 初始化 Ebiten，並不斷呼叫 update()
	ebiten.Run(update, 320, 240, 2, "Hello world!")
}
