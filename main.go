package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/imthewatermelon/some_math/sketcher"
)

func main() {
	Xs := []float64{0, 0.5, 1.0, 1.5, 2.0, 2.5, 3.0}
	Ys := []float64{4.2, 26.1, 40.1, 46.0, 43.9, 33.7, 15.8}
	s, err := sketcher.NewSketcher(Xs, Ys, []func(x float64) float64{func(x float64) float64 { return 275.428 * math.Pow(x, 2) }}, 0, 10, 0, 50, "Somethin2g.png")
	if err != nil {
		panic(err)
	}
	err = s.Sketch()
	if err != nil {
		panic(err)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
