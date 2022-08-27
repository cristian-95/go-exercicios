/* Exercicio 3.2 - Faça experimentos com visualizações de outras funções do pacote math.
 * Você pode gerar padroes como caixa de ovo, morrinhos ou uma sela?*/
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
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
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

	// Faz uma projeção isometrica de (x,y,z) sobre (sx,sy) do canvas SVG 2D
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	return -(x * y) / 100 // out.svg

}
