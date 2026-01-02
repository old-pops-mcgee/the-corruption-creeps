//go:build js && wasm

package main

import (
	"fmt"

	rl "github.com/BrownNPC/Raylib-Go-Wasm/raylib"
)

type GameState int

const (
	Waiting GameState = iota
	PressedButton
)

const (
	baseTurns       int = 20
	levelMultiplier int = 5
)

type Game struct {
	Level          int
	CountdownTimer int
	GameState      GameState
}

func calculateTurns(level int) int {
	return baseTurns + (level-1)*levelMultiplier
}

func InitGame() Game {
	return Game{
		Level:          1,
		CountdownTimer: calculateTurns(1),
		GameState:      Waiting,
	}
}

func (g *Game) render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.DrawText(fmt.Sprintf("Turns Remaining: %d", g.CountdownTimer), 50, 50, 20, rl.RayWhite)
	rl.EndDrawing()
}

func (g *Game) update() {
	if g.GameState == PressedButton {
		// Process the button press

		// Lower the countdown timer
		g.CountdownTimer -= 1

		if g.CountdownTimer == 0 {
			// TODO: The final boss fight
		}

		// Update the state back to waiting
		g.GameState = Waiting
	}

}

func (g *Game) handleInput() {
	if rl.IsKeyDown(rl.KeyP) {
		g.GameState = PressedButton
	}
}
