package main

import (
	"fmt"

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
		rl.DrawFPS(scr.Width-50, 0)
		rl.EndDrawing()

		if rl.IsKeyDown(rl.KeyA) {
			rl.DrawText("hello from A", 500, 250, 26, rl.Blue)
			fmt.Println("A was pressed")
		}
	}

}

func main() {
	startScreen()
}
