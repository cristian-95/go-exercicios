package framestats

import (
	"math"
	"net/url"
)

type FrameStats struct {
	width, height, cells            int
	xyrange, xyscale, zscale, angle float64
}

func (fs *FrameStats) GetDefaultValues() {
	fs.width = 600
	fs.height = 320
	fs.cells = 100
	fs.xyrange = 30
	fs.xyscale = float64(fs.width) / 2 / fs.xyrange
	fs.zscale = float64(fs.height) * 0.4
	fs.angle = math.Pi / 6
}

func (fs *FrameStats) Update(url.Values) {

}
