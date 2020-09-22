package main

import (
	"fmt"
	"image/color"
	"log"
	"reflect"

	"github.com/hajimehoshi/ebiten"
)

const (
	dftLen = 20
	dftWid = 10
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

// Block 儲存貪吃蛇（影像區塊）的參數
type Block struct {
	// 縮放（Scale）及縮放濾鏡（filter）
	img *ebiten.Image
	// 位置（Position）
	x float64
	y float64
}

// Game 要實現 Update、Layout 函數，若有 Draw 函數則會根據參數將指定影像顯示至遊戲畫面上
type Game struct {
	blocks    []*Block
	direction Direction
}

func newBlock(x, y float64) (*Block, error) {
	// 縮放（Scale）及縮放濾鏡（filter）
	img, err := ebiten.NewImage(1, 10, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	// 顏色白色
	img.Fill(color.White)
	return &Block{
		img: img,
		x:   x,
		y:   y,
	}, nil
}

// initBlocks 在遊戲開始前設置 i 個影像區塊
func initBlocks() ([]*Block, error) {
	var blocks []*Block
	for i := 0; i < 20; i++ {
		block, err := newBlock(float64(i), 0)
		if err != nil {
			return nil, err
		}
		// 寫入
		blocks = append(blocks, block)
	}
	return blocks, nil
}

// Init 將初始化的 Block（貪吃蛇影像區塊）寫入 Game.blocks
func (g *Game) Init() error {
	blocks, err := initBlocks()
	if err != nil {
		return err
	}
	g.blocks = blocks
	g.direction = Right
	return nil
}

// Update 在這邊用不到
func (g *Game) Update(screen *ebiten.Image) error {

	head := g.blocks[len(g.blocks)-1]
	block, err := newBlock(head.x+1, head.y)
	if err != nil {
		return err
	}
	g.blocks = append(g.blocks, block)
	return nil
}

// Draw 將 block 內的參數取出，並根據設計圖把圖繪出
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

// Layout 設置遊戲整體外觀的解析度
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
	//return 800,600
}

func main() {
	// 設置遊戲整體外觀的大小
	ebiten.SetWindowSize(320, 240)
	//ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Snake")
	game := &Game{}

	// 初始化 Block（貪吃蛇影像區塊）
	if err := game.Init(); err != nil {
		log.Fatalln(err)
	}

	// 開始遊戲，根據實現的 Update、Layout、 Draw 函數，來進行執行
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

// MemAdr 取得變數型能及記憶體位址
func MemAdr(p interface{}) {
	fmt.Printf("%v %p\n", reflect.TypeOf(p), &p)
}
