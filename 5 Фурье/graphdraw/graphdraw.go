package graphdraw

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	//"gonum.org/v1/plot/vg/draw"
	"Furje/fx"
	"image/color"
)

func DrawFunction(F func(float64) float64, text string) *plot.Plot {
	// Создаем новый график
	p := plot.New()

	// Подписи
	p.Title.Text = text
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	// Массив точек
	points := make(plotter.XYs, 0)
	for x := fx.A; x <= fx.B; x += fx.H {
		points = append(points, plotter.XY{X: x, Y: F(x)})
	}

	// Создаем линию функции
	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}

	line.Color = color.RGBA{R: 0, G: 128, B: 128, A: 255} // Цвет линии
	p.Add(line)

	return p

}

func DrawFunction2(F func(float64) float64, text string, n *plot.Plot) *plot.Plot {
	// Создаем новый график
	p := plot.New()

	// Подписи
	p.Title.Text = text
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	// Массив точек
	points := make(plotter.XYs, 0)
	for x := fx.A; x <= fx.B; x += fx.H {
		y := F(x)
		//fmt.Println(x, " : ", y)
		points = append(points, plotter.XY{X: x, Y: y})

	}

	// Создаем линию функции
	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}

	line.Color = color.RGBA{R: 192, G: 64, B: 0, A: 255} // Цвет линии
	p.Add(line)

	return p

}

func DrawPoint(x, y float64, p *plot.Plot, r, g, b uint8) {
	scatterP := make(plotter.XYs, 0)
	scatterP = append(scatterP, plotter.XY{X: x, Y: y})
	scatter, _ := plotter.NewScatter(scatterP)

	scatter.GlyphStyle.Color = color.RGBA{R: r, G: g, B: b, A: 255} // Цвет точек
	scatter.GlyphStyle.Radius = vg.Points(1)                        // Размер точек
	p.Add(scatter)
}

// Касательная
/*
func DrawKasLine(x0 float64, p *plot.Plot, F func(float64) float64) {
	df := DF(x0)
	k := F(x0) - df*x0
	points := make(plotter.XYs, 0)
	points = append(points, plotter.XY{X: fx.A, Y: df*fx.A + k})
	points = append(points, plotter.XY{X: x0, Y: df*x0 + k})
	points = append(points, plotter.XY{X: fx.B, Y: df*fx.B + k})

	// Создаем линию функции
	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}

	line.Color = color.RGBA{R: 128, G: 0, B: 128, A: 255} // Цвет линии

	p.Add(line)
}
*/
func DrawXmark(x0 float64, p *plot.Plot, F func(float64) float64, r, g, b uint8) {
	y0 := F(x0)
	{
		points := make(plotter.XYs, 0)
		points = append(points, plotter.XY{X: x0 + 0.03, Y: y0 + 0.03})
		points = append(points, plotter.XY{X: x0 - 0.03, Y: y0 - 0.03})

		// Создаем линию функции
		line, err := plotter.NewLine(points)
		if err != nil {
			panic(err)
		}

		line.Color = color.RGBA{R: r, G: g, B: b, A: 255} // Цвет линии

		p.Add(line)
	}
	{
		points := make(plotter.XYs, 0)
		points = append(points, plotter.XY{X: x0 - 0.03, Y: y0 + 0.03})
		points = append(points, plotter.XY{X: x0 + 0.03, Y: y0 - 0.03})

		// Создаем линию функции
		line, err := plotter.NewLine(points)
		if err != nil {
			panic(err)
		}

		line.Color = color.RGBA{R: r, G: g, B: b, A: 255} // Цвет линии

		p.Add(line)
	}
}

func Save(p *plot.Plot, name string) {
	// Сохраняем в файл
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "images/"+name+".png"); err != nil {
		panic(err)
	}
}
