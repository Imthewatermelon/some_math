package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/imthewatermelon/some_math/sketcher"
)

func main() {
	Xs := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	Ys := []float64{50, 62, 73, 80, 71, 60, 51, 43, 29, 20, 28, 41, 49}

	s, err := sketcher.NewSketcher(Xs, Ys, []func(x float64) float64{
		func(x float64) float64 { return 30*math.Sin(math.Pi*x/6) + 50 },
	},
		0, 12, 20, 80, "Somethin2g.png")
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
