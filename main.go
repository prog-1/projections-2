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
	screenWidth  = 2560
	screenHeight = 1440
)

type point struct {
	x, y, z float64
}

type Game struct {
	o      [8]point
	join   [][2]int
	camera point
}

var col = color.RGBA{244, 212, 124, 255}

func NewGame(width, height int) *Game {
	return &Game{
		o: [8]point{
			{-300, -300, -300},
			{-300, 300, -300},
			{300, 300, -300},
			{300, -300, -300},
			{-300, -300, 300},
			{-300, 300, 300},
			{300, 300, 300},
			{300, -300, 300},
		},
		join: [][2]int{
			{5, 6},
			{5, 4},
			{6, 7},
			{7, 4},

			{1, 2},
			{1, 0},
			{0, 3},
			{3, 2},

			{6, 2},
			{6, 5},
			{5, 1},
			{1, 2},

			{7, 3},
			{7, 4},
			{0, 4},
			{0, 3},

			{5, 1},
			{5, 4},
			{0, 4},
			{0, 1},

			{6, 2},
			{6, 7},
			{3, 7},
			{3, 2},
		},
	}
}

func Add(a, b point) point {
	return point{a.x + b.x, a.y + b.y, a.z + b.z}
}

func Sub(a, b point) point {
	return point{a.x - b.x, a.y - b.y, a.z - b.z}
}

func Divide(v point, a float64) point {
	return point{v.x / a, v.y / a, v.z / a}
}

func Multiply(v point, a float64) point {
	return point{v.x * a, v.y * a, v.z * a}
}

func Mod(a point) float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y + a.z*a.z)
}
func Cross(a, b point) point {
	return point{
		a.y*b.z - b.y*a.z,
		a.z*b.x - b.z*a.x,
		a.x*b.y - b.x*a.y,
	}
}
func Dot(a, b point) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}
func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return screenWidth, screenHeight
}

func CameraProject(o point) point {
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		o.y += 10
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		o.y -= 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		o.x -= 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		o.x += 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		o.z -= 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		o.z += 10
	}
	return o
}

func (g *Game) rotX() {
	for i, j := range g.o {
		g.o[i].y = j.y*math.Cos(math.Pi/180) + j.z*math.Sin(math.Pi/180)
		g.o[i].z = j.y*math.Sin(math.Pi/180) - j.z*math.Cos(math.Pi/180)
	}
}

func (g *Game) rotY() {
	for i, j := range g.o {
		g.o[i].x = j.x*math.Sin(math.Pi/180) - j.z*math.Cos(math.Pi/180)
		g.o[i].z = j.x*math.Cos(math.Pi/180) + j.z*math.Sin(math.Pi/180)
	}
}

func (g *Game) rotZ() {
	for i, j := range g.o {
		g.o[i].x = j.x*math.Cos(math.Pi/180) - j.y*math.Sin(math.Pi/180)
		g.o[i].y = j.x*math.Sin(math.Pi/180) + j.y*math.Cos(math.Pi/180)
	}
}

func (g *Game) Update() error {
	g.rotX()
	g.rotY()
	g.rotZ()
	g.camera = CameraProject(g.camera)
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < len(g.join); i += 4 {
		for i1 := i; i1 < i+4; i1++ {
			ebitenutil.DrawLine(screen,
				((g.o[g.join[i1][0]].x+g.camera.x)/(g.o[g.join[i1][0]].z+g.camera.z+1500))*-900+float64(screenWidth/2),
				((g.o[g.join[i1][0]].y+g.camera.y)/(g.o[g.join[i1][0]].z+g.camera.z+1500))*-900+float64(screenHeight/2),
				((g.o[g.join[i1][1]].x+g.camera.x)/(g.o[g.join[i1][1]].z+g.camera.z+1500))*-900+float64(screenWidth/2),
				((g.o[g.join[i1][1]].y+g.camera.y)/(g.o[g.join[i1][1]].z+g.camera.z+1500))*-900+float64(screenHeight/2),
				col)
		}
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
