package framestats

import (
	"log"
	"math"
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
