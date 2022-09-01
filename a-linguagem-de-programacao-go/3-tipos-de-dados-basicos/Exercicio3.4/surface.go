/* Exercicio 3.4 - . */
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type FrameStats struct {
	Width   int
	Height  int
	Cells   int
	XYrange float64
	XYscale float64
	Zscale  float64
	Angle   float64
	Color   string
}

func main() {
	http.HandleFunc("/", surfaceHandler)
	log.Printf("listening on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func surfaceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fs := FrameStats{}
	fs.GetDefaultValues()
	sin30, cos30 := math.Sin(fs.Angle), math.Cos(fs.Angle) //seno(30°), cosseno(30°)

	params := r.URL.Query()
	fs.Update(params)

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill:white; strokewidth: 0.5' "+
		"width='%d' height='%d'>", fs.Color, fs.Width, fs.Height)
	for i := 0; i < fs.Cells; i++ {
		for j := 0; j < fs.Cells; j++ {
			ax, ay := corner(i+1, j, sin30, cos30, fs)
			bx, by := corner(i, j, sin30, cos30, fs)
			cx, cy := corner(i, j+1, sin30, cos30, fs)
			dx, dy := corner(i+1, j+1, sin30, cos30, fs)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprint(w, "</svg>")
}

func corner(i, j int, sin30, cos30 float64, fs FrameStats) (float64, float64) {
	// Encontra o ponto (x,y) no canto da celula (i,j)
	x := fs.XYrange * (float64(i)/float64(fs.Cells) - 0.5)
	y := fs.XYrange * (float64(j)/float64(fs.Cells) - 0.5)

	// Calcula a altura z da superficie
	z := f(x, y)
	if math.IsInf(z, 0) {
		return 0.0, 0.0
	}

	// Faz uma projeção isometrica de (x,y,z) sobre (sx,sy) do canvas SVG 2D
	sx := float64(fs.Width/2) + (x-y)*cos30*fs.XYscale
	sy := float64(fs.Height/2) + (x+y)*sin30*fs.XYscale - z*fs.Zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distancia de (0,0)
	return math.Sin(r) / r
}

/******************* framestats: */

func (fs *FrameStats) GetDefaultValues() {
	fs.Width = 600
	fs.Height = 320
	fs.Cells = 100
	fs.XYrange = 30
	fs.XYscale = float64(fs.Width) / 2 / fs.XYrange
	fs.Zscale = float64(fs.Height) * 0.4
	fs.Angle = math.Pi / 6
	fs.Color = "black"
}

func (fs *FrameStats) Update(params url.Values) {
	if len(params) == 0 {
		return
	} else {
		height, err := strconv.Atoi(strings.Join((params["height"]), ""))
		if err != nil {
			log.Fatalf("Erro: %v", err)
			goto getWidth
		}
		fs.Height = height

	getWidth:
		width, err := strconv.Atoi(strings.Join((params["width"]), ""))
		if err != nil {
			log.Fatalf("Erro: %v", err)
			goto getColor
		}
		fs.Width = width

	getColor:
		color := strings.Join((params["width"]), "")
		if color != "" {
			fs.Color = color
		}
	}

}
