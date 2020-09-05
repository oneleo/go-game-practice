package screen

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/oneleo/go-game-practice/transform"
)

// DrawSquare will draw a square image
func DrawSquare(screen *ebiten.Image, pos *transform.Position, scl *transform.Scale, clr color.Color) error {
	// 建立一個 width x height 的新畫布
	square, err := ebiten.NewImage(scl.Width, scl.Height, ebiten.FilterNearest)
	// 替 square 畫布填上白色
	square.Fill(clr)
	// 建立一個空的選項建構體
	opts := &ebiten.DrawImageOptions{}
	// 修改選項，新增 Translate 變形效果
	opts.GeoM.Translate(pos.Tx, pos.Ty)
	// 渲染 square 畫布到 screen 主畫布上並套用空的選項
	screen.DrawImage(square, opts)

	return err
}
