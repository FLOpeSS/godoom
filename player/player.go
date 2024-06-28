package player

import (
	"fmt"

	"github.com/FLOpeSS/doom/tp"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position_y int32
	Position_x int32
	Size       int32
	life       int32
}

func (p *Player) Movement(scr tp.Screen) {
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		p.Position_x -= 1
		fmt.Println("p.Position to left", p.Position_x)
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		p.Position_x += 1
		fmt.Println("p.Position to right: ", p.Position_x)
	}

	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		p.Position_y -= 1
		fmt.Println("p.Position to Up: ", p.Position_y)
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		p.Position_y += 1
		fmt.Println("p.Position to Down: ", p.Position_y)
	}
}
