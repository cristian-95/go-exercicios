// Exercicio 1.12 - Modifique o servidor Lissajous para ler valores de parametros
// do URL. Por exemplo [...] localhost:8000/?cycles=20 Utilize a função strconv.Atoi
// [...]

// Em progresso
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
)

type Options struct {
	nframes int
	delay   int
	cycles  float64
	size    float64
	res     float64
}

// do Exercicio 1.6:
var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{255, 0, 0, 255},     // Red
	color.RGBA{0, 255, 0, 255},     // Green
	color.RGBA{0, 0, 255, 255},     // Blue
	color.RGBA{232, 121, 23, 255},  // Orange
	color.RGBA{255, 0, 221, 255},   // Pink
	color.RGBA{0, 251, 255, 255},   // Cyan
	color.RGBA{111, 255, 0, 255},   // Neon green
	color.RGBA{128, 128, 128, 255}, // Gray
}

const (
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", myHandler)
	log.Println("Serving at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	opt := Options{
		delay:   8,
		nframes: 64,
		cycles:  5,
		size:    200,
		res:     0.001,
	}

	err := r.PostForm
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parsing form: %v", err)
	}

	lissajous(w, opt)
}

func lissajous(out io.Writer, opt Options) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: opt.nframes}
	phase := 0.0
	colorIndex := uint8(blackIndex)

	for i := 0; i < opt.nframes; i++ {

		rect := image.Rect(0, 0, int(2*opt.size+1), int(2*opt.size+1))
		img := image.NewPaletted(rect, palette)
		if i%2 == 0 {
			colorIndex = randomColor()
		}

		for t := 0.0; t < opt.cycles*2*math.Pi; t += opt.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(opt.size+x*opt.size+0.5), int(opt.size+y*opt.size+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, opt.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func randomColor() uint8 {
	return uint8(rand.Intn(len(palette)))
}
