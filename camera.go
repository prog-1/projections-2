package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//---------------------------Declaration--------------------------------

const (
	sW = 640
	sH = 480
)

type Game struct {
	width, height int      //screen width and height
	cube          [8]point // cube points
	edges         [][2]int // cube edges
	cp            point    // cube center point
	cam           camera   //camera
}

type point struct {
	x, y, z float64
}

type camera struct {
	pos point //camera position
	rot point //camera rotation angle
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.cam.pos.x += 7
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.cam.pos.x -= 7
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.cam.pos.y -= 7
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.cam.pos.y += 7
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.cam.pos.z += 7
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.cam.pos.z -= 7
	}
	if ebiten.IsKeyPressed(ebiten.KeyNumpad8) { // - X
		g.cam.rot.x -= math.Pi / 200
	}
	if ebiten.IsKeyPressed(ebiten.KeyNumpad2) { // + X
		g.cam.rot.x += math.Pi / 200
	}
	if ebiten.IsKeyPressed(ebiten.KeyNumpad4) { // - Y
		g.cam.rot.y += math.Pi / 200
	}
	if ebiten.IsKeyPressed(ebiten.KeyNumpad6) { // + Y
		g.cam.rot.y -= math.Pi / 200
	}
	if ebiten.IsKeyPressed(ebiten.KeyNumpad9) { // + Z
		g.cam.rot.z += math.Pi / 200
	}
	if ebiten.IsKeyPressed(ebiten.KeyNumpad7) { // - Z
		g.cam.rot.z -= math.Pi / 200
	}
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {

	//cube line draw
	for i := range g.edges { //for all cube edges
		g.drawLine(screen, g.cube[g.edges[i][0]], g.cube[g.edges[i][1]])
	}

}

//line draw with central projection
func (g *Game) drawLine(screen *ebiten.Image, a, b point) {

	a = updatePoint(a, g.cam, g.cp)
	b = updatePoint(b, g.cam, g.cp)
	//draw function
	ebitenutil.DrawLine(screen, a.x, a.y, b.x, b.y, color.White)
}

//-------------------------Functions----------------------------------

func updatePoint(p point, cam camera, cp point) point {

	//adding cube center coordinates to point
	p = add(p, cp)

	//subtracting camera position from point
	p = subtract(p, cam.pos)

	//point rotation by camera
	p.rotateX(-cam.rot.x)
	p.rotateY(-cam.rot.y)
	p.rotateZ(-cam.rot.z)

	//central projections
	k := float64(250)
	centralProjection(&p, k)

	//p.y *= -1

	//moving point to center
	p.x += sW / 2
	p.y += sH / 2

	return p
}

func add(a, b point) (res point) {
	res.x = a.x + b.x
	res.y = a.y + b.y
	res.z = a.z + b.z
	return res
}

func subtract(a, b point) (res point) {
	res.x = a.x - b.x
	res.y = a.y - b.y
	res.z = a.z - b.z
	return res
}

func centralProjection(p *point, k float64) {
	//k - scaling koefficient

	//formulas
	x1 := (p.x / p.z) * k
	y1 := (p.y / p.z) * k

	p.x, p.y = x1, y1
}

//rotates point around X axis
func (p *point) rotateX(angle float64) {

	p.y = (p.y)*math.Cos(angle) - (p.z)*math.Sin(angle)
	p.z = (p.y)*math.Sin(angle) + (p.z)*math.Cos(angle)

}

//rotates point around Y axis
func (p *point) rotateY(angle float64) {

	p.x = (p.x)*math.Cos(angle) + (p.z)*math.Sin(angle)
	p.z = (p.x)*-math.Sin(angle) + (p.z)*math.Cos(angle)

}

//rotates point around Z axis
func (p *point) rotateZ(angle float64) {

	p.x = (p.x)*math.Cos(angle) - (p.y)*math.Sin(angle)
	p.y = (p.x)*math.Sin(angle) + (p.y)*math.Cos(angle)

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

	//cube center point
	cp := point{0, 0, 300}

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
	cam := camera{pos: point{0, 0, 0}, rot: point{0, 0, 0}}

	//RETURN
	return &Game{width: width, height: height, cp: cp, cube: cube, edges: edges, cam: cam}
}
