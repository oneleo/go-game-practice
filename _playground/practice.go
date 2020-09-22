package main

import (
	"fmt"
)

type Game interface {
	Update() error
	Layout() error
	//Draw() error
}

type MyGame struct {
}

func (g *MyGame) Init() error {
	return nil
}

func (g *MyGame) Update() error {
	return nil
}

func (g *MyGame) Draw() error {
	return nil
}

func (g *MyGame) Layout() error {
	return nil
}

func testFunc(game Game) {
	if _, ok := game.(interface{ Draw() error }); ok {
		fmt.Println("ok:", ok)
	}
}

func main() {
	g := &MyGame{}
	testFunc(g)
}
