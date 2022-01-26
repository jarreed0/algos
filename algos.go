package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"github.com/faiface/pixel/imdraw"
	"math/rand"
	"time"
)

type bar struct {
        rect  pixel.Rect
        color color.Color
}

var (
	dataSize = 600
	barWidth float64
	imd *imdraw.IMDraw
)

const (
	WIDTH = 1024
	HEIGHT = 700
)

func randData(size int) []float64 {
	data := make([]float64, size)
        for i := 0; i < len(data); i++ {
                data[i] = float64(HEIGHT/dataSize * (dataSize-i-1))//float64(dataSize)//float64((i+1) * ((HEIGHT-50)/dataSize))
        }
        rand.Seed(time.Now().UnixNano())
        rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	return data
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Visualizing Algos",
		Bounds: pixel.R(0, 0, WIDTH, HEIGHT),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	barWidth = float64(WIDTH) / float64(dataSize)

	bars := make([]bar, dataSize)

	data := randData(dataSize)
	count := 0

	for !win.Closed() {
		win.Update()
		win.Clear(colornames.Lightslategray)
		for i := 0; i < len(bars); i++ {
			//bars[i].rect = pixel.R(barWidth*float64(i), HEIGHT-data[i], barWidth*float64(i) + barWidth, 0)
			bars[i].rect = pixel.R(barWidth*float64(i), data[i], barWidth*float64(i) + barWidth, 0)
			bars[i].color = colornames.Lightblue
		}
		for b := 0; b < len(bars); b++ {
                	imd := imdraw.New(nil)
                	imd.Color = bars[b].color
                	imd.Push(bars[b].rect.Min, bars[b].rect.Max)
                	imd.Rectangle(0)
                	imd.Draw(win)
        	}
		reset := false
		if count > 45 {
			reset = true
			for i := 0; i < len(data)-1; i++ {
				if(data[i] > data[i+1]) {
					old := data[i]
					data[i] = data[i+1]
					data[i+1] = old
					reset = false
					//i = len(data)+1
				}
			}
		}
		count++
		if reset {
			data = randData(dataSize);
			count = 0
		}
	}
}

func main() {
	pixelgl.Run(run)
}
