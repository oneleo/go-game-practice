package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

//Snake 00:14:10
type Game struct {
}

var blocks []*ebiten.Image

func init() {
	blocks = initBlockImage()
}

func initBlockImage() []*ebiten.Image {
	var images []*ebiten.Image
	for i := 0; i < 20; {
		img, _ := ebiten.NewImage(1, 10, ebiten.FilterDefault)
		img.Fill(color.White)
		images = append(images, img)
	}
	return images
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Fill(color.White)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240

}

func main() {
	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
