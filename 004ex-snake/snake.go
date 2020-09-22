package main

import (
	"fmt"
	"image/color"
	"log"
	"reflect"

	"github.com/hajimehoshi/ebiten"
)

// Block 儲存貪吃蛇（圖片）的參數
type Block struct {
	// 縮放（Scale）及縮放濾鏡（filter）
	img *ebiten.Image
	// 位置（Position）
	x float64
	y float64
}

type Game struct {
	blocks []*Block
}

func initBlocks() ([]*Block, error) {
	var blocks []*Block
	for i := 0; i < 20; i++ {
		img, err := ebiten.NewImage(1, 10, ebiten.FilterDefault)
		if err != nil {
			return nil, err
		}
		img.Fill(color.White)
		blocks = append(blocks, &Block{
			img: img,
			x:   float64(i),
			y:   0,
		})
	}
	return blocks, nil
}
func (g *Game) Init() error {
	blocks, err := initBlocks()
	if err != nil {
		return err
	}
	g.blocks = blocks
	MemAdr(g.blocks)
	MemAdr(blocks)
	return nil
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Fill(color.White)
	for _, block := range g.blocks {
		geom := ebiten.GeoM{}
		geom.Translate(block.x, block.y)
		options := ebiten.DrawImageOptions{
			GeoM: geom,
		}
		screen.DrawImage(block.img, &options)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240

}

func main() {
	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Snake")
	game := &Game{}

	if err := game.Init(); err != nil {
		log.Fatalln(err)
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func MemAdr(p interface{}) {
	fmt.Println(fmt.Sprintf("%v %p", reflect.TypeOf(p), &p))
}
