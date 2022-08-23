/* Exercicio 1.6: Modifique o programa Lissajous para gerar imagens em varias cores adicionando
 * mais valores a palette para então exibi-las alterando o terceiro argumento de
 * SetColorIndex
 */
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

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
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	colorIndex := uint8(blackIndex)

	for i := 0; i < nframes; i++ {

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		if i%2 == 0 {
			colorIndex = randomColor()
		}

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func randomColor() uint8 {
	return uint8(rand.Intn(len(palette)))
}
