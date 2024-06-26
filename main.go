package main

import (
	"fmt"

	"github.com/FLOpeSS/doom/player"
	"github.com/FLOpeSS/doom/tp"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var scr = tp.Screen{Width: 1290, Height: 960}

func createPlayer() *player.Player {
	p := &player.Player{
		Position_x: scr.Width / 3,
		Position_y: scr.Height / 2,
	}
	return p
}

func startScreen() {
	rl.InitWindow(scr.Width, scr.Height, "basic window")
	p := createPlayer()
	fmt.Println(p)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawText("hello world", scr.Width/2, scr.Height/2, 20, rl.Pink)
		rl.DrawText("hello Player", p.Position_x, p.Position_y, 20, rl.Pink)
		rl.DrawFPS(scr.Width-50, 0)
		rl.EndDrawing()
		p.Movement_x()

	}

}

func main() {
	startScreen()
}
