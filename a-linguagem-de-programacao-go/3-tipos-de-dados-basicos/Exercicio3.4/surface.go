/* Exercicio 3.4 - . */
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", surfaceHandler)
	log.Printf("listening on http://localhost:8000")
	http.ListenAndServe(":8000", nil)

}

func surfaceHandler(w http.ResponseWriter, r *http.Request) {
	fs := FrameStats{}
	fs.GetDefaultValues()
	sin30, cos30 := math.Sin(fs.angle), math.Cos(fs.angle) //seno(30°), cosseno(30°)

	//fmt.Println(params["color"])
	//fmt.Printf("\n\n***** %T -> %v len:(%d)*****************\n\n", params, params, len(params))
	params := r.URL.Query()
	fs.Update(params)

	dumy := ""
	// fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
	// 	"style='stroke: black; fill:white; strokewidth: 0.5' "+
	// 	"width='%d' height='%d'>", width, height)
	for i := 0; i < fs.cells; i++ {
		for j := 0; j < fs.cells; j++ {
			ax, ay := corner(i+1, j, sin30, cos30, fs)
			bx, by := corner(i, j, sin30, cos30, fs)
			cx, cy := corner(i, j+1, sin30, cos30, fs)
			dx, dy := corner(i+1, j+1, sin30, cos30, fs)
			dumy = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	//fmt.Println("</svg>")
	fmt.Printf("%T", dumy)
}

func corner(i, j int, sin30, cos30 float64, fs FrameStats) (float64, float64) {
	// Encontra o ponto (x,y) no canto da celula (i,j)
	x := fs.xyrange * (float64(i)/float64(fs.cells) - 0.5)
	y := fs.xyrange * (float64(j)/float64(fs.cells) - 0.5)

	// Calcula a altura z da superficie
	z := f(x, y)
	if math.IsInf(z, 0) {
		return 0.0, 0.0
	}

	// Faz uma projeção isometrica de (x,y,z) sobre (sx,sy) do canvas SVG 2D
	sx := float64(fs.width/2) + (x-y)*cos30*fs.xyscale
	sy := float64(fs.height/2) + (x+y)*sin30*fs.xyscale - z*fs.zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distancia de (0,0)
	return math.Sin(r) / r
}
