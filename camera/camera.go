package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//---------------------------Declaration--------------------------------

const (
	sW = 640
	sH = 480
)

type Game struct {
	width, height int //screen width and height
	//global variables
	cp    point    //cube central point
	cube  [8]point // cube points
	edges [][2]int // cube edges
	cam   point    //camera
}

type point struct {
	x, y, z float64
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.cam.x++
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.cam.x--
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.cam.y++
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.cam.y--
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.cam.z++
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.cam.z--
	}

	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
	//cube drawing without normals and stuff
	for i := range g.edges {
		g.drawLine(screen, g.cube[g.edges[i][0]], g.cube[g.edges[i][1]])
	}
}

//line draw with central projection
func (g *Game) drawLine(screen *ebiten.Image, a, b point) {

	//camera modification
	a.x -= g.cam.x
	a.y += g.cam.y
	a.z -= g.cam.z

	b.x -= g.cam.x
	b.y += g.cam.y
	b.z -= g.cam.z

	//central projection
	k := float64(400)
	cproj(&a, g.cp, k)
	cproj(&b, g.cp, k)

	//adding central point
	a.x += g.cp.x
	a.y += g.cp.y

	b.x += g.cp.x
	b.y += g.cp.y

	//draw function
	ebitenutil.DrawLine(screen, a.x, a.y, b.x, b.y, color.White)
}

//-------------------------Functions----------------------------------

//central projection
func cproj(p *point, cp point, k float64) {
	//k - scaling koefficient

	//formulas
	z1 := p.z + k
	x1 := (p.x / z1) * k
	y1 := (p.y / z1) * k

	p.x, p.y = x1, y1
}

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {

	//Window
	ebiten.SetWindowSize(sW, sH)
	ebiten.SetWindowTitle("Cube x Camera")

	//Game instance
	g := NewGame(sW, sH)                      //creating game instance
	if err := ebiten.RunGame(g); err != nil { //running game
		log.Fatal(err)
	}
}

//New game instance function
func NewGame(width, height int) *Game {

	//GAME
	//center point
	cp := point{sW / 2, sH / 2, 100}

	//cube points
	cube := [8]point{
		/*0*/ {-100, -100, -100},
		/*1*/ {100, -100, -100},
		/*2*/ {100, 100, -100},
		/*3*/ {-100, 100, -100},

		/*4*/ {-100, -100, 100},
		/*5*/ {100, -100, 100},
		/*6*/ {100, 100, 100},
		/*7*/ {-100, 100, 100},
	}

	edges := [][2]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0},
		{4, 5}, {5, 6}, {6, 7}, {7, 4},
		{0, 4}, {1, 5}, {2, 6}, {3, 7},
	}

	//CAMERA
	cam := point{0, 0, 0}

	//RETURN
	return &Game{width: width, height: height, cp: cp, cube: cube, edges: edges, cam: cam}
}
