package main

import (
	"image/color"
	"log"
	"math"

	v "github.com/34thSchool/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

func CentralProjection(a v.Vec, k float64) v.Vec {
	return v.Vec{
		-(a.X / a.Z) * k,
		-(a.Y / a.Z) * k,
		-k,
	}
}

func DrawLine(screen *ebiten.Image, a, b v.Vec, clr color.Color) {
	k := float64(-250)
	a = CentralProjection(a, k)
	b = CentralProjection(b, k)
	ebitenutil.DrawLine(screen, a.X, a.Y, b.X, b.Y, color.RGBA{255, 102, 204, 255})
}

func RotateCube(c *v.Cube, r v.Rotator, screen *ebiten.Image) {
	&c.
	ctr := v.Add(v.Div(v.Sub(c.p[6], c.p[0]), 2), c.p[0])
	for i := range c.p {
		c.p[i] = v.Sub(c.p[i], ctr)
		c.p[i].Rotate(r)
		c.p[i] = v.Add(c.p[i], ctr)
	}
}

type game struct {
	c            []Cube
	screenBuffer *ebiten.Image
}

func NewGame() *game {
	halfWidth, halfHeight := float64(screenWidth/2), float64(screenHeight/2)
	return &game{
		[]Cube{
			{
				[8]v.Vec{
					{-200 + halfWidth, -200 + halfHeight, 200}, // NearBottomLeft
					{-200 + halfWidth, 200 + halfHeight, 200},  // NearTopLeft
					{200 + halfWidth, 200 + halfHeight, 200},   // NearTopRight
					{200 + halfWidth, -200 + halfHeight, 200},  // NearBottomRightS

					{-200 + halfWidth, -200 + halfHeight, 250}, // FarBottomLeft
					{-200 + halfWidth, 200 + halfHeight, 250},  // FarTopLeft
					{200 + halfWidth, 200 + halfHeight, 250},   // FarTopRight
					{200 + halfWidth, -200 + halfHeight, 250},  // FarBottomRight
				},
			},
			// {
			// 	[8]v.Vec{
			// 		{-400 + halfWidth, -200 + halfHeight, 200}, // NearBottomLeft
			// 		{-400 + halfWidth, 200 + halfHeight, 200},  // NearTopLeft
			// 		{halfWidth, 200 + halfHeight, 200},         // NearTopRight
			// 		{halfWidth, -200 + halfHeight, 200},        // NearBottomRight

			// 		{-400 + halfWidth, -200 + halfHeight, 600}, // FarBottomLeft
			// 		{-400 + halfWidth, 200 + halfHeight, 600},  // FarTopLeft
			// 		{halfWidth, 200 + halfHeight, 600},         // FarTopRight
			// 		{halfWidth, -200 + halfHeight, 600},        // FarBottomRight
			// 	},
			// },
		},
		ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	for i := range g.c {
		g.c[i].Rotate(g.screenBuffer, Rotator{0, 0, math.Pi / 180})
	}
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	for i := range g.c {
		g.c[i].Draw(screen, color.RGBA{255, 102, 204, 255})
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
