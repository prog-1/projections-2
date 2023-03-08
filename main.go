package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 960
	screenHeight = 640
)

type Point struct {
	x, y, z float64
}

type Game struct {
	width, height int
	dots          []Point
	camera        Point
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.camera.x--
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.camera.x++
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.camera.z--
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.camera.z++
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.camera.y--
	}
	if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) {
		g.camera.y++
	}
	return nil
}

func (g *Game) DrawLine(screen *ebiten.Image, a, b Point) {
	z1 := a.z + 500
	z2 := b.z + 500
	x1 := a.x / z1
	y1 := a.y / z1
	x2 := b.x / z2
	y2 := b.y / z2
	ebitenutil.DrawLine(screen, x1*(500-g.camera.z)+float64(g.width)/2-g.camera.x, y1*(500-g.camera.z)+float64(g.height)/2-g.camera.y, x2*(500-g.camera.z)+float64(g.width)/2-g.camera.x, y2*(500-g.camera.z)+float64(g.height)/2-g.camera.y, color.White)

}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawLine(screen, g.dots[0], g.dots[1])
	g.DrawLine(screen, g.dots[1], g.dots[2])
	g.DrawLine(screen, g.dots[2], g.dots[3])
	g.DrawLine(screen, g.dots[3], g.dots[0])

	g.DrawLine(screen, g.dots[4], g.dots[5])
	g.DrawLine(screen, g.dots[5], g.dots[6])
	g.DrawLine(screen, g.dots[6], g.dots[7])
	g.DrawLine(screen, g.dots[7], g.dots[4])

	g.DrawLine(screen, g.dots[0], g.dots[4])
	g.DrawLine(screen, g.dots[1], g.dots[5])
	g.DrawLine(screen, g.dots[2], g.dots[6])
	g.DrawLine(screen, g.dots[3], g.dots[7])
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		dots: []Point{
			{-100, -100, -100},
			{-100, 100, -100},
			{100, 100, -100},
			{100, -100, -100},
			{-100, -100, 100},
			{-100, 100, 100},
			{100, 100, 100},
			{100, -100, 100},
		},
		camera: Point{0, 0, -1},
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
