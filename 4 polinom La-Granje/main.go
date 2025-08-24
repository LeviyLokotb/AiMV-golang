package main

import (
	"fmt"
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Point struct {
	X, Y float64
}

// Значение полинома Лагранжа в точке x
func L(x float64, points []Point) float64 {
	n := len(points)
	result := 0.0
	for j := 0; j < n; j++ {
		l := 1.0
		for i := 0; i < n; i++ {
			if i == j {
				continue
			}
			l *= (x - points[i].X) / (points[j].X - points[i].X)
		}
		result += l * points[j].Y
	}
	return result
}

// Вычисление полинома Лагранжа

func main() {
	points := make([]Point, 0)
	fmt.Println("================\nПолином Лагранжа\n================")
	fmt.Println("Введите режим custom/default")
	var mode string
	fmt.Scan(mode)
	if mode != "c" {
		mode = "d"
	}

	var count_of_points int
	var min, max float64
	if mode == "d" {
		fmt.Println("Введите количество точек: ")
		fmt.Scan(&count_of_points)
		max = -1024.0
		min = 1024.0
		for i := 0; i < count_of_points; i++ {
			fmt.Print("Точка #", i+1, ": ")
			var x, y float64
			fmt.Scan(&x, &y)
			if x > max {
				max = x
			}
			if x < min {
				min = x
			}
			points = append(points, Point{X: x, Y: y})
		}
	} else {
		count_of_points = 5
		max = 21
		min = 1.9

	}

	fmt.Println("Полином Лагранжа - полином ", count_of_points-1, " степени")
	image := DrawFunction(min, max, points)
	for _, dot := range points {
		DrawPoint(dot.X, dot.Y, image, 0, 0, 0)
	}
	Save(image, "LaGrange")
}

func DrawFunction(minX, maxX float64, mypoints []Point) *plot.Plot {
	// Создаем новый график
	p := plot.New()

	// Подписи
	p.Title.Text = "Полином Лагранжа"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	// Массив точек (массив большой, поэтому создаётся в хипе воизбежание)
	points := make(plotter.XYs, 0)
	for x := minX - (math.Abs(minX) * 0.1); x <= maxX*1.1; x += 0.1 {
		points = append(points, plotter.XY{X: x, Y: L(x, mypoints)})
	}

	// Создаем линию функции
	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}

	line.Color = color.RGBA{R: 191, G: 64, B: 0, A: 255} // Цвет линии
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
	if err := p.Save(6*vg.Inch, 4*vg.Inch, name+".png"); err != nil {
		panic(err)
	}
}
