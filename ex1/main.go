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
	screenWidth  = 640
	screenHeight = 480
)

type Point struct {
	x, y, z float64
}

type Game struct {
	points [8]Point
	faces  [][4]int
}

func (g *Game) RotationX(ang float64) {
	for i, p := range g.points {
		g.points[i].x = p.x*math.Cos(ang) - p.y*math.Sin(ang)
		g.points[i].y = p.x*math.Sin(ang) + p.y*math.Cos(ang)
	}
}

func (g *Game) RotationY(ang float64) {
	for i, p := range g.points {
		g.points[i].x = p.x*math.Cos(ang) - p.z*math.Sin(ang)
		g.points[i].z = p.x*math.Sin(ang) + p.z*math.Cos(ang)
	}
}

func (g *Game) RotationZ(ang float64) {
	for i, p := range g.points {
		g.points[i].y = p.y*math.Cos(ang) + p.z*math.Sin(ang)
		g.points[i].z = -p.y*math.Sin(ang) + p.z*math.Cos(ang)
	}
}

func (g *Game) CameraProject(p Point) {
	for i := range g.points {
		g.points[i].x -= p.x
		g.points[i].y -= p.y
		g.points[i].z -= p.z
	}
}

func (g *Game) Update() error {
	var p Point

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.x = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.x = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.y = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.y = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.RotationX(math.Pi / 360)
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.RotationX(-math.Pi / 360)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.RotationY(-math.Pi / 360)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.RotationY(math.Pi / 360)
	}

	if ebiten.IsKeyPressed(ebiten.KeyZ) && ebiten.IsKeyPressed(ebiten.Key1) {
		p.z = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyZ) && ebiten.IsKeyPressed(ebiten.Key2) {
		p.z = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyZ) && ebiten.IsKeyPressed(ebiten.Key3) {
		g.RotationZ(-math.Pi / 360)
	}
	if ebiten.IsKeyPressed(ebiten.KeyZ) && ebiten.IsKeyPressed(ebiten.Key4) {
		g.RotationZ(math.Pi / 360)
	}
	g.CameraProject(p)

	return nil
}

func Cross(u, v Point) Point {
	return Point{u.y*v.z - u.z*v.y, u.z*v.x - u.x*v.z, u.x*v.y - u.y*v.x}
}

func Sub(a, b Point) Point {
	return Point{a.x - b.x, a.y - b.y, a.z - b.z}
}

func (g *Game) Draw(screen *ebiten.Image) {
	a, b := screenWidth, screenHeight
	for i, j := range g.faces {
		n := Cross(Sub(g.points[j[1]], g.points[j[0]]), Sub(g.points[j[1]], g.points[j[2]]))
		if n.z < 0 {
			ebitenutil.DrawLine(screen,
				(g.points[g.faces[i][0]].x/(g.points[g.faces[i][0]].z-1000))*-900+float64(a/2),
				(g.points[g.faces[i][0]].y/(g.points[g.faces[i][0]].z-1000))*-900+float64(b/2),
				(g.points[g.faces[i][1]].x/(g.points[g.faces[i][1]].z-1000))*-900+float64(a/2),
				(g.points[g.faces[i][1]].y/(g.points[g.faces[i][1]].z-1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
			ebitenutil.DrawLine(screen,
				(g.points[g.faces[i][1]].x/(g.points[g.faces[i][1]].z-1000))*-900+float64(a/2),
				(g.points[g.faces[i][1]].y/(g.points[g.faces[i][1]].z-1000))*-900+float64(b/2),
				(g.points[g.faces[i][2]].x/(g.points[g.faces[i][2]].z-1000))*-900+float64(a/2),
				(g.points[g.faces[i][2]].y/(g.points[g.faces[i][2]].z-1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
			ebitenutil.DrawLine(screen,
				(g.points[g.faces[i][2]].x/(g.points[g.faces[i][2]].z-1000))*-900+float64(a/2),
				(g.points[g.faces[i][2]].y/(g.points[g.faces[i][2]].z-1000))*-900+float64(b/2),
				(g.points[g.faces[i][3]].x/(g.points[g.faces[i][3]].z-1000))*-900+float64(a/2),
				(g.points[g.faces[i][3]].y/(g.points[g.faces[i][3]].z-1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
			ebitenutil.DrawLine(screen,
				(g.points[g.faces[i][3]].x/(g.points[g.faces[i][3]].z-1000))*-900+float64(a/2),
				(g.points[g.faces[i][3]].y/(g.points[g.faces[i][3]].z-1000))*-900+float64(b/2),
				(g.points[g.faces[i][0]].x/(g.points[g.faces[i][0]].z-1000))*-900+float64(a/2),
				(g.points[g.faces[i][0]].y/(g.points[g.faces[i][0]].z-1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
		}
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := ebiten.RunGame(&Game{
		//    Pos6________Pos7
		//   /|          / |
		//Pos4_|_______Pos5|
		//  |  |       |   |
		//  |  Pos3____|___| Pos2
		//  | /        |  /
		//  Pos0_______Pos1
		points: [8]Point{
			{x: 100, y: 100, z: 100},
			{x: 100, y: -100, z: 100},
			{x: -100, y: -100, z: 100},
			{x: -100, y: 100, z: 100},
			{x: 100, y: 100, z: -100},
			{x: 100, y: -100, z: -100},
			{x: -100, y: 100, z: -100},
			{x: -100, y: -100, z: -100},
		},
		faces: [][4]int{
			{5, 4, 0, 1},
			{7, 2, 3, 6},
			{6, 4, 5, 7},
			{3, 0, 4, 6},
			{7, 5, 1, 2},
			{3, 2, 1, 0},
		}}); err != nil {
		log.Fatal(err)
	}
}
