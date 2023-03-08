package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	winTitle     = "Cube"
	screenWidth  = 1000
	screenHeight = 1000
	dpi          = 100
)

var c = color.RGBA{R: 255, G: 255, B: 255, A: 255}

type (
	point struct {
		x, y, z float64
	}
	game struct {
		p      [8]point
		planes [][2]int
	}
)

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

func (g *game) rotateX() {
	for i, v := range g.p {
		g.p[i].x = v.x*math.Cos(0.0025*5) - v.y*math.Sin(0.0025*5)
		g.p[i].y = v.x*math.Sin(0.0025*5) + v.y*math.Cos(0.0025*5)
	}
}

func (g *game) rotateY() {
	for i, v := range g.p {
		g.p[i].x = v.x*math.Cos(0.00474533*5) - v.z*math.Sin(0.00474533*5)
		g.p[i].z = v.x*math.Sin(0.00474533*5) + v.z*math.Cos(0.00474533*5)
	}

}
func (g *game) rotateZ() {
	for i, v := range g.p {
		g.p[i].y = v.y*math.Cos(-0.00474533*5) - v.z*math.Sin(-0.00474533*5)
		g.p[i].z = v.y*math.Sin(-0.00474533*5) + v.z*math.Cos(-0.00474533*5)
	}

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

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }

func (g *game) Update() error {
	// g.rotateX()
	// g.rotateY()
	// g.rotateZ()
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		for i := range g.p {
			g.p[i].z -= 10
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		for i := range g.p {
			g.p[i].x -= 10
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		for i := range g.p {
			g.p[i].z += 10
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		for i := range g.p {
			g.p[i].x += 10
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		for i := range g.p {
			g.p[i].y -= 10
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyControl) {
		for i := range g.p {
			g.p[i].y += 10
		}
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	for i := 0; i < len(g.planes); i += 4 {
		// a := Sub(g.p[g.planes[i][1]], g.p[g.planes[i][0]])
		// b := Sub(g.p[g.planes[i+1][1]], g.p[g.planes[i+1][0]])
		// center := Divide(Add(g.p[g.planes[i+1][1]], g.p[g.planes[i][0]]), 2)
		// screen.Set(int(center.x)+screenWidth/2, int(center.y)+screenHeight/2, c)
		// cross := Cross(a, b)
		// p := point{float64(screenWidth / 2), float64(screenHeight / 2), 0}
		// if Dot(cross, Add(Divide(center, center.z), p)) < 0 {
		for i1 := i; i1 < i+4; i1++ {
			ebitenutil.DrawLine(screen,
				(g.p[g.planes[i1][0]].x/(g.p[g.planes[i1][0]].z+1500))*-900+float64(screenWidth/2),
				(g.p[g.planes[i1][0]].y/(g.p[g.planes[i1][0]].z+1500))*-900+float64(screenHeight/2),
				(g.p[g.planes[i1][1]].x/(g.p[g.planes[i1][1]].z+1500))*-900+float64(screenWidth/2),
				(g.p[g.planes[i1][1]].y/(g.p[g.planes[i1][1]].z+1500))*-900+float64(screenHeight/2),
				color.White)
		}

		// } else {
		// 	for i1 := i; i1 < i+4; i1++ {
		// 		ebitenutil.DrawLine(screen,
		// 			(g.p[g.planes[i1][0]].x/(g.p[g.planes[i1][0]].z+1500))*-900+float64(screenWidth/2),
		// 			(g.p[g.planes[i1][0]].y/(g.p[g.planes[i1][0]].z+1500))*-900+float64(screenHeight/2),
		// 			(g.p[g.planes[i1][1]].x/(g.p[g.planes[i1][1]].z+1500))*-900+float64(screenWidth/2),
		// 			(g.p[g.planes[i1][1]].y/(g.p[g.planes[i1][1]].z+1500))*-900+float64(screenHeight/2),
		// 			color.RGBA{0xff, 0xff, 0xff, 28})

		// 	}
		// }
	}

}

func main() {
	ebiten.SetWindowTitle(winTitle)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizable(true)
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
func NewGame() *game {

	return &game{
		p: [8]point{
			{300, -300, -300},  //0
			{-300, -300, -300}, //1
			{-300, 300, -300},  //2
			{300, 300, -300},   //3

			{300, -300, 300},  //4
			{-300, -300, 300}, //5
			{-300, 300, 300},  //6
			{300, 300, 300},   //7
		},
		planes: [][2]int{
			// near plane
			{0, 1},
			{1, 2},
			{2, 3},
			{3, 0},
			// far plane
			{4, 7},
			{7, 6},
			{6, 5},
			{5, 4},
			//  top plane
			{5, 1},
			{1, 0},
			{0, 4},
			{4, 5},
			//left plane
			{6, 2},
			{2, 1},
			{1, 5},
			{5, 6},
			//right plane
			{4, 0},
			{0, 3},
			{3, 7},
			{7, 4},
			//bottom plane
			{6, 7},
			{7, 3},
			{3, 2},
			{2, 6},
		},
	}
}
