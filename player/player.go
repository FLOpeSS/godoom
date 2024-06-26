package player

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position_y int32
	Position_x int32
	Size       int
	life       int
}

func (p *Player) Positioning() {
}

func (p *Player) Movement_x() int32 {
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		p.Position_x -= 1
		fmt.Println("p.Position to left", p.Position_y)
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		p.Position_x += 1
		fmt.Println("p.Position to right: ", p.Position_x)
	}
	return p.Position_x
}
