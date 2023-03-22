package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	Height = 480
	Width  = 640
)

type Game struct {
	c      cube
	camera Point
}

type Point struct {
	x, y, z float64
}

type cube struct {
	Points []Point
	Joints [][4]int
}

func (c *cube) RotateX(angle float64) {
	for i := range c.Points {
		c.Points[i].y = c.Points[i].y*math.Cos(angle) + c.Points[i].z*math.Sin(angle)
		c.Points[i].z = -c.Points[i].y*math.Sin(angle) + c.Points[i].z*math.Cos(angle)
	}
}

func (c *cube) RotateY(angle float64) {
	for i := range c.Points {
		c.Points[i].x = c.Points[i].x*math.Cos(angle) - c.Points[i].z*math.Sin(angle)
		c.Points[i].z = c.Points[i].x*math.Sin(angle) + c.Points[i].z*math.Cos(angle)
	}
}

func (c *cube) RotateZ(angle float64) {
	for i := range c.Points {
		c.Points[i].x = c.Points[i].x*math.Cos(angle) - c.Points[i].y*math.Sin(angle)
		c.Points[i].y = c.Points[i].x*math.Sin(angle) + c.Points[i].y*math.Cos(angle)
	}
}

func Dot(a, b Point) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func Sub(a, b Point) Point {
	return Point{a.x - b.x, a.y - b.y, a.z - b.z}
}

func (c *cube) Draw(screen *ebiten.Image, camera Point) {
	for _, v := range c.Joints {
		// a := crossProduct(Sub(c.Points[v[1]], c.Points[v[0]]), Sub(c.Points[v[1]], c.Points[v[2]]))
		// if Dot(Point{0, 0, 1}, a) > 0 {
		// 	continue
		// }
		for i := 1; i <= 3; i++ {

			z1 := (c.Points[v[i-1]].z - camera.z) + 500
			z2 := (c.Points[v[i]].z - camera.z) + 500
			x1 := c.Points[v[i-1]].x / z1
			y1 := c.Points[v[i-1]].y / z1
			x2 := c.Points[v[i]].x / z2
			y2 := c.Points[v[i]].y / z2
			ebitenutil.DrawLine(screen, x1*500+Width/2-camera.x, y1*500+Height/2-camera.y, x2*500+Width/2-camera.x, y2*500+Height/2-camera.y, color.White)
		}
		z1 := (c.Points[v[0]].z - camera.z) + 500
		z2 := (c.Points[v[3]].z - camera.z) + 500
		x1 := c.Points[v[0]].x / z1
		y1 := c.Points[v[0]].y / z1
		x2 := c.Points[v[3]].x / z2
		y2 := c.Points[v[3]].y / z2
		ebitenutil.DrawLine(screen, x1*500+Width/2-camera.x, y1*500+Height/2-camera.y, x2*500+Width/2-camera.x, y2*500+Height/2-camera.y, color.White)
	}

	// for _, v := range c.Joints {
	// }
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.camera.z += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.camera.z -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.camera.x += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.camera.x -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.camera.y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.camera.y -= 1
	}
	g.c.RotateX(math.Pi / 1000)
	g.c.RotateY(math.Pi / 750)
	g.c.RotateZ(math.Pi / 450)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.c.Draw(screen, g.camera)

}

func crossProduct(u, v Point) Point {
	return Point{u.y*v.z - u.z*v.y, u.z*v.x - u.x*v.z, u.x*v.y - u.y*v.x}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func main() {
	c := cube{
		Points: []Point{
			{-100, -100, 100},
			{-100, 100, 100},
			{-100, -100, -100},
			{100, -100, -100},
			{100, -100, 100},
			{100, 100, -100},
			{100, 100, 100},
			{-100, 100, -100},
		},
		Joints: [][4]int{
			{4, 0, 1, 6},
			{3, 5, 7, 2},
			{3, 4, 6, 5},
			{2, 0, 4, 3},
			{7, 1, 0, 2},
			{7, 5, 6, 1},
		},
	}
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{c, Point{0, 0, 0}}); err != nil {
		log.Fatal(err)
	}
}
