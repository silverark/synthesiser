// Package chart will chart audio data and output it as an image.
package chart

import (
	"github.com/wcharczuk/go-chart"
	"os"
)

func ChartIt(filename string, data []float64) {
	var x []float64
	var y []float64

	for i := 0.0; i < float64(len(data)); i++ {
		x = append(x, i)
	}
	y = data
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: y,
			},
		},
	}

	f, _ := os.Create(filename + ".png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
