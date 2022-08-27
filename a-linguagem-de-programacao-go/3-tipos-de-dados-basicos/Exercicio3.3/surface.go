/* Exercicio 3.3 - Pinte cada poligono de acordo com sua altura, de modo que os picos
 * tenham cor vermelha((#ff0000) e os vales sejam azuis(#0000ff))
 * OBS: mudei para as cores do Brasil por motivos de Copa do Mundo. */
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // tamanho do canvas
	cells         = 100                 // numero de células da grade
	xyrange       = 30.0                // intervalos dos eixos (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels por unidade de x ou y
	zscale        = height * 0.4        // pixels por unidade de z
	angle         = math.Pi / 6         // ângulo dos eixos, x,y (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //seno(30°), cosseno(30°)
var color string = "black"

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: black; fill:white; strokewidth: 0.5' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
			color = "black"
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Encontra o ponto (x,y) no canto da celula (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Calcula a altura z da superficie
	z := f(x, y)
	if math.IsInf(z, 0) {
		return 0.0, 0.0
	}
	if z >= .01 {
		if z >= 0.25 {
			color = "blue"
		} else {
			color = "yellow"
		}
	} else if z <= -.01 {
		color = "green"
	}

	// Faz uma projeção isometrica de (x,y,z) sobre (sx,sy) do canvas SVG 2D
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distancia de (0,0)
	return math.Sin(r) / r
}
