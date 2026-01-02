//go:build js && wasm

package main

import (
	"embed"

	rl "github.com/BrownNPC/Raylib-Go-Wasm/raylib"
)

//go:embed assets
var ASSETS embed.FS

func main() {
	rl.InitWindow(800, 500, "Hello World")
	rl.InitAudioDevice()
	rl.SetTargetFPS(30)

	game := InitGame()

	var update = func() {
		game.handleInput()
		game.update()
		game.render()
	}

	rl.SetMainLoop(update)
	for !rl.WindowShouldClose() {
		update()
	}

	rl.CloseAudioDevice()
	rl.CloseWindow()
}
