package main

import (
	"image/color"
	"log"
	"math"
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

	a, b = WorldToCamera(&g.Cam, a), WorldToCamera(&g.Cam, b)

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
	Pos           v.Vec
	Rad           v.Rotator
	MovementSpeed float64
	RotationSpeed float64
}

func (Cam *Camera) Update(dt float64) {
	// Movement:
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		Cam.Pos.X += Cam.MovementSpeed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		Cam.Pos.X -= Cam.MovementSpeed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		Cam.Pos.Y += Cam.MovementSpeed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) {
		Cam.Pos.Y -= Cam.MovementSpeed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		Cam.Pos.Z += Cam.MovementSpeed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		Cam.Pos.Z -= Cam.MovementSpeed * dt
	}

	// Rotation:
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		Cam.Rad.Z += Cam.RotationSpeed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		Cam.Rad.Z -= Cam.RotationSpeed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		Cam.Rad.Y -= Cam.RotationSpeed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		Cam.Rad.Y += Cam.RotationSpeed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		Cam.Rad.X -= Cam.RotationSpeed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		Cam.Rad.X += Cam.RotationSpeed * dt
	}
}

func WorldToCamera(Cam *Camera, a v.Vec) v.Vec {
	a = v.Sub(a, Cam.Pos)
	a.RotateX(-Cam.Rad.X)
	a.RotateY(-Cam.Rad.Y)
	a.RotateZ(-Cam.Rad.Z)
	return a
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

					{-200, -200, 700}, // FarBottomLeft
					{-200, 200, 700},  // FarTopLeft
					{200, 200, 700},   // FarTopRight
					{200, -200, 700},  // FarBottomRight
				},
			},
		},
		Camera{
			v.Vec{X: 0, Y: 0, Z: 0},
			v.Rotator{X: 0, Y: 0, Z: 0},
			1,
			math.Pi / 720,
		},
		0,
		ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	dt := float64(time.Now().UnixMilli() - g.prevFrameTime)
	g.prevFrameTime = time.Now().UnixMilli()

	g.Cam.Update(dt)

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
