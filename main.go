package main

import (
	"image/color"
	"log"
	"time"

	v "github.com/34thSchool/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

func CentralProjection(a v.Vec, k float64) v.Vec {
	return v.Vec{
		X: a.X / a.Z * k,
		Y: a.Y / a.Z * k,
	}

}

func DrawLine(screen *ebiten.Image, g *game, a, b v.Vec, clr color.Color) {
	halfWidth, halfHeight := float64(screenWidth/2), float64(screenHeight/2)

	a, b = v.Sub(a, g.Cam.Pos), v.Sub(b, g.Cam.Pos)

	k := float64(250)
	a, b = CentralProjection(a, k), CentralProjection(b, k)
	ebitenutil.DrawLine(screen, a.X+halfWidth, -a.Y+halfHeight, b.X+halfWidth, -b.Y+halfHeight, clr)
}

type Cube struct {
	P [8]v.Vec
}

func (c *Cube) Draw(screen *ebiten.Image, g *game, clr color.Color) {
	for _, f := range [][]int{
		{0, 1, 2, 3}, // Near
		{7, 6, 5, 4}, // Far
		{4, 5, 1, 0}, // Left
		{1, 5, 6, 2}, // Top
		{3, 2, 6, 7}, // Right
		{4, 0, 3, 7}, // Bottom
	} {
		DrawLine(screen, g, c.P[f[0]], c.P[f[1]], clr)
		DrawLine(screen, g, c.P[f[1]], c.P[f[2]], clr)
		DrawLine(screen, g, c.P[f[2]], c.P[f[3]], clr)
		DrawLine(screen, g, c.P[f[3]], c.P[f[0]], clr)
	}
}

type Camera struct {
	Pos   v.Vec
	Speed float64
}

func (Cam *Camera) ProcessMovement(dt float64) {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		Cam.Pos.X += Cam.Speed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		Cam.Pos.X -= Cam.Speed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		Cam.Pos.Y += Cam.Speed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		Cam.Pos.Y -= Cam.Speed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		Cam.Pos.Z += Cam.Speed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		Cam.Pos.Z -= Cam.Speed * dt
	}
}

type game struct {
	C             []Cube
	Cam           Camera
	prevFrameTime int64
	screenBuffer  *ebiten.Image
}

func NewGame() *game {
	return &game{
		[]Cube{
			{
				[8]v.Vec{
					{-200, -200, 400}, // NearBottomLeft
					{-200, 200, 400},  // NearTopLeft
					{200, 200, 400},   // NearTopRight
					{200, -200, 400},  // NearBottomRight

					{-200, -200, 800}, // FarBottomLeft
					{-200, 200, 800},  // FarTopLeft
					{200, 200, 800},   // FarTopRight
					{200, -200, 800},  // FarBottomRight
				},
			},
		},
		Camera{v.Vec{0, 0, 0}, 1},
		0,
		ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	dt := float64(time.Now().UnixMilli() - g.prevFrameTime)
	g.prevFrameTime = time.Now().UnixMilli()

	g.Cam.ProcessMovement(dt)

	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	for i := range g.C {
		g.C[i].Draw(screen, g, color.RGBA{255, 102, 204, 255})
	}
	screen.DrawImage(g.screenBuffer, &ebiten.DrawImageOptions{})
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
