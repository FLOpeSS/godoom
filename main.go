package main

import (
	"github.com/FLOpeSS/doom/tp"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func startScreen() {
	scr := tp.Screen{Width: 1290, Height: 960}
	rl.InitWindow(scr.Width, scr.Height, "basic window")
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawText("hello world", scr.Width/2, scr.Height/2, 20, rl.Pink)
		rl.EndDrawing()
	}

}

func main() {
	startScreen()
}
