package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	// 先幫畫布填上 #FF0000 顏色
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})

	// 顯示一串除錯用的文字
	ebitenutil.DebugPrint(screen, "Our first game in Ebiten!")

	if err := drawSquare(screen, 64, 64, 16, 16, color.White); err != nil {
		log.Fatal(err)
	}

	if err := drawSquare(screen, 64, 128, 16, 32, color.Black); err != nil {
		log.Fatal(err)
	}

	if err := drawSquare(screen, 128, 64, 32, 16, color.NRGBA{0xff, 0xff, 0x00, 0xff}); err != nil {
		log.Fatal(err)
	}

	return nil
}

func drawSquare(screen *ebiten.Image, tx float64, ty float64, width int, height int, clr color.Color) error {
	// 建立一個 width x height 的新畫布
	square, err := ebiten.NewImage(width, height, ebiten.FilterNearest)
	// 替 square 畫布填上白色
	square.Fill(clr)
	// 建立一個空的選項建構體
	opts := &ebiten.DrawImageOptions{}
	// 修改選項，新增 Translate 變形效果
	opts.GeoM.Translate(tx, ty)
	// 渲染 square 畫布到 screen 主畫布上並套用空的選項
	screen.DrawImage(square, opts)

	return err
}

func main() {
	// 初始化 Ebiten，並不斷呼叫 update()
	ebiten.Run(update, 320, 240, 2, "Hello world!")
}
