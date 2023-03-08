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
	cp        point    //cube central point
	cube     [8]point // cube points
	edges      [][2]int // cube edges
}

type point struct {
	x, y, z float64
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
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

	//central projection
	k := float64(400)
	cproj(&a, g.cp, k)
	cproj(&b, g.cp, k)

	//draw function
	ebitenutil.DrawLine(screen, g.cp.x+a.x, g.cp.y+a.y, g.cp.x+b.x, g.cp.y+b.y, color.White)
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
	ebiten.SetWindowResizable(true) //enablening window resize

	//Game instance
	g := NewGame(sW, sH)                      //creating game instance
	if err := ebiten.RunGame(g); err != nil { //running game
		log.Fatal(err)
	}
}

//New game instance function
func NewGame(width, height int) *Game {

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
		{0,1}, {1,2}, {2,3}, {3,0},
		{4,5}, {5,6}, {6,7}, {7,4},
		{0,4}, {1,5}, {2,6}, {3,7},
	}

	return &Game{width: width, height: height, cp: cp, cube: cube, edges: edges}
}
