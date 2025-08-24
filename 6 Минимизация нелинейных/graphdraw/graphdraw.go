package graphdraw

import (
	"image/color"
	"minnelbezogr/fx"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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
	for x := fx.A; x <= fx.B; x += 0.01 {
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

func DrawPoint(x, y float64, p *plot.Plot, r, g, b uint8) {
	scatterP := make(plotter.XYs, 0)
	scatterP = append(scatterP, plotter.XY{X: x, Y: y})
	scatter, _ := plotter.NewScatter(scatterP)

	scatter.GlyphStyle.Color = color.RGBA{R: r, G: g, B: b, A: 255} // Цвет точек
	scatter.GlyphStyle.Radius = vg.Points(1)                        // Размер точек
	p.Add(scatter)
}

func Save(p *plot.Plot, name string) {
	// Сохраняем в файл
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "images/"+name+".png"); err != nil {
		panic(err)
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Реализуем интерфейс для нашей функции
type functionGrid struct {
	xmin, xmax     float64
	ymin, ymax     float64
	xsteps, ysteps int
	f              func(x, y float64) float64
}

func (g *functionGrid) Dims() (c, r int) { return g.xsteps, g.ysteps }
func (g *functionGrid) X(c int) float64 {
	return g.xmin + (g.xmax-g.xmin)*float64(c)/float64(g.xsteps-1)
}
func (g *functionGrid) Y(r int) float64 {
	return g.ymin + (g.ymax-g.ymin)*float64(r)/float64(g.ysteps-1)
}
func (g *functionGrid) Z(c, r int) float64 {
	x := g.X(c)
	y := g.Y(r)
	return g.f(x, y)
}

func DrawFunction3D(F func(float64, float64) float64, zmin float64, title string) *plot.Plot {
	// Создаем новый график
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	grid := &functionGrid{
		xmin:   -4,
		xmax:   4,
		ymin:   -4,
		ymax:   4,
		xsteps: 400,
		ysteps: 400,
		f:      fx.F,
	}
	// Создаем контурный график
	var levels []float64
	for i := zmin; i <= zmin+3.2; i += 0.2 {
		levels = append(levels, i)
	}
	contour := plotter.NewContour(grid, levels, palette.Heat(16, 1))

	p.Add(contour)

	return p
}

func Trajectory(p *plot.Plot, xHist, yHist []float64, title string) *plot.Plot {
	//p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Создаем точки для траектории
	pts := make(plotter.XYs, len(xHist))
	for i := range xHist {
		DrawPoint(xHist[i], yHist[i], p, 0, 0, 0)
		pts[i].X = xHist[i]
		pts[i].Y = yHist[i]
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	line.Color = color.RGBA{R: 0, G: 128, B: 128, A: 255}

	p.Add(line)
	return p
}

func Save3D(p *plot.Plot, name string) {
	// Сохраняем в файл
	err := p.Save(10*vg.Inch, 8*vg.Inch, "images/"+name+".png")
	if err != nil {
		panic(err)
	}
}
