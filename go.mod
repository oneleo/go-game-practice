module github.com/oneleo/go-game-practice

go 1.15

require (
	github.com/hajimehoshi/ebiten v1.11.8
	github.com/joho/godotenv v1.3.0
	github.com/oneleo/go-game-practice/env v0.0.0
	github.com/oneleo/go-game-practice/screen v0.0.0
	github.com/oneleo/go-game-practice/transform v0.0.0
)

replace (
	github.com/oneleo/go-game-practice/env => ./env
	github.com/oneleo/go-game-practice/screen => ./screen
	github.com/oneleo/go-game-practice/transform => ./transform
)
