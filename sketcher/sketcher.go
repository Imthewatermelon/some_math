package sketcher

import (
	"errors"
	"image/color"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func customize_plane(p *plot.Plot, s *plotter.Scatter) *plot.Plot {
	grid := &plotter.Grid{
		Vertical: draw.LineStyle{
			Color: color.Gray{128},
			Width: vg.Points(0.5),
		},
		Horizontal: draw.LineStyle{
			Color: color.Gray{128},
			Width: vg.Points(0.5),
		},
	}

	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	p.Add(grid)
	p.Add(s)

	return p
}

type Sketcher struct {
	Xs, Ys                     []float64
	Functions                  []func(x float64) float64
	X_min, X_max, Y_min, Y_max float64
	Filename                   string
}

// Returns a plotter.XYs type to use  plot/plotter package.
func get_XYs(Xs, Ys []float64) plotter.XYs {
	XYs := make(plotter.XYs, len(Xs))
	for i := range XYs {
		XYs[i].X = Xs[i]
		XYs[i].Y = Ys[i]
	}
	return XYs
}

func NewSketcher(Xs, Ys []float64, Functions []func(x float64) float64,
	X_min, X_max, Y_min, Y_max float64, Filename string) (*Sketcher, error) {

	if len(Xs) != len(Ys) {
		return nil, errors.New("the number of x coordinates is not equal to the number of y coordinates")
	} else if X_min >= X_max || Y_min >= Y_max {
		return nil, errors.New("the maximum limit is less then or equal to the minimum limit.")
	}

	// TODO: Check if the filename must contain the extension of the result file
	return &Sketcher{Xs, Ys, Functions, X_min, X_max, Y_min, Y_max, Filename}, nil
}

func (sketcher *Sketcher) set_ranges_of_axes(p *plot.Plot) *plot.Plot {
	// Set the axis ranges.  Unlike other data sets,
	// functions don't set the axis ranges automatically
	// since functions don't necessarily have a
	// finite range of x and y values.
	p.X.Min = sketcher.X_min
	p.X.Max = sketcher.X_max
	p.Y.Min = sketcher.Y_min
	p.Y.Max = sketcher.Y_max

	return p
}

func add_functions_into_plotter(p *plot.Plot, functions []func(float64) float64) *plot.Plot {
	for _, fun := range functions {
		random_function := plotter.NewFunction(fun)
		// random_function.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
		random_function.Dashes = []vg.Length{vg.Points(2), vg.Points(0)}
		random_function.Width = vg.Points(2)
		random_function.Color = color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)), A: 255}
		p.Add(random_function)
	}
	return p
}

func (sketcher *Sketcher) Sketch() error {
	p := plot.New()
	XYs := get_XYs(sketcher.Xs, sketcher.Ys)

	// Make a scatter plotter and set its style.
	scatter, err := plotter.NewScatter(XYs)
	if err != nil {
		return err
	}

	p = add_functions_into_plotter(p, sketcher.Functions)
	p = customize_plane(p, scatter)
	p = sketcher.set_ranges_of_axes(p)

	// Save the plot to a PNG file.
	p.Save(8*vg.Inch, 8*vg.Inch, sketcher.Filename)

	return nil
}
