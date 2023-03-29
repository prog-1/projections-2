package main

import (
	"image/color"
	"log"
	"math"
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
	dots          []*Point
	camera        Point
	cameraAngleX  float64
	cameraAngleY  float64
	cameraAngleZ  float64
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func cameraProject(p *Point, camera Point) *Point {
	return &Point{p.x - camera.x, p.y - camera.y, p.z - camera.z}
}
func (g *Game) CameraRotateY(p *Point, angle float64) *Point {
	p.x = p.x*math.Cos(angle) + p.z*math.Sin(angle)
	p.z = -p.x*math.Sin(angle) + p.z*math.Cos(angle)
	return p
}
func (g *Game) CameraRotateX(p *Point, angle float64) *Point {
	p.y = p.y*math.Cos(angle) - p.z*math.Sin(angle)
	p.z = p.y*math.Sin(angle) + p.z*math.Cos(angle)
	return p
}

func (g *Game) CameraRotateZ(p *Point, angle float64) *Point {
	p.x = p.x*math.Cos(angle) - p.y*math.Sin(angle)
	p.y = p.x*math.Sin(angle) + p.y*math.Cos(angle)
	return p
}

func (g *Game) Update() error {
	g.camera = Point{0, 0, 0}
	g.cameraAngleX = 0
	g.cameraAngleY = 0
	g.cameraAngleZ = 0
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.camera.x = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.camera.x = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.camera.z = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.camera.z = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.camera.y = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) {
		g.camera.y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.cameraAngleY = math.Pi / 360
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.cameraAngleY = -math.Pi / 360
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.cameraAngleX = math.Pi / 360
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.cameraAngleX = -math.Pi / 360
	}
	if ebiten.IsKeyPressed(ebiten.KeyPageUp) {
		g.cameraAngleZ = math.Pi / 360
	}
	if ebiten.IsKeyPressed(ebiten.KeyPageDown) {
		g.cameraAngleZ = -math.Pi / 360
	}
	for i := range g.dots {
		g.dots[i] = cameraProject(g.dots[i], g.camera)
		g.dots[i] = g.CameraRotateX(g.dots[i], g.cameraAngleX)
		g.dots[i] = g.CameraRotateY(g.dots[i], g.cameraAngleY)
		g.dots[i] = g.CameraRotateZ(g.dots[i], g.cameraAngleZ)
	}
	return nil
}

func (g *Game) DrawLine(screen *ebiten.Image, a, b *Point) {
	z1 := a.z + 500
	z2 := b.z + 500
	x1 := a.x / z1
	y1 := a.y / z1
	x2 := b.x / z2
	y2 := b.y / z2
	ebitenutil.DrawLine(screen, x1*500+float64(g.width)/2, y1*500+float64(g.height)/2, x2*500+float64(g.width)/2, y2*500+float64(g.height)/2, color.White)

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
		dots: []*Point{
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
