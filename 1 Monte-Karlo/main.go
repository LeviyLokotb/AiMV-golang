package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg" // Импорт для поддержки JPEG
	"image/png"

	//"image/png"
	"math"
	"math/rand"
	"os"
	"strconv"

	foo "montecarlo/funcmy"
)

const SEArectA = 1200    // км длина карты
const SEArectB = 800     // км ширина карты
var rectA, rectB float64 // стороны фигуры (объявляем глобально для удобства)

type Color struct {
	r int32
	g int32
	b int32
}

// задаём цвета интересующих фигур
var BlackSea Color = Color{78, 98, 157}
var AzovSea Color = Color{163, 73, 163}

// определяем соответствие диапазону цветов
func WhatAColor(test, find Color, luft uint32) bool {
	if math.Abs(float64(test.r-find.r)) <= float64(luft) {
		if math.Abs(float64(test.g-find.g)) <= float64(luft) {
			if math.Abs(float64(test.b-find.b)) <= float64(luft) {
				return true
			}
		}
	}
	return false
}

func RandomTest(count int, img image.Image, find Color, front *image.RGBA) int {
	countPixels := 0
	size := img.Bounds().Size()
	// fmt.Println("sizeX: ", size.X)
	// fmt.Println("sizeY: ", size.Y)

	// Цвет для замены
	//white := color.RGBA{255, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}

	for ; count > 0; count-- {
		// Задаем координаты пикселя (x, y)
		x, y := rand.Intn(size.X-1)+1, rand.Intn(size.Y-1)+1

		// Получаем цвет пикселя
		color1 := img.At(x, y)

		// Преобразуем цвет в модель RGBA
		r, g, b, a := color1.RGBA()

		// Приводим значения к диапазону 0-255
		r = r >> 8
		g = g >> 8
		b = b >> 8
		a = a >> 8
		old := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		var pixel Color = Color{int32(r), int32(g), int32(b)}
		if WhatAColor(pixel, find, 5) {
			countPixels++
			// закрашиваем пиксели
			front.Set(x, y, red)
		} else {
			front.Set(x, y, old)
		}

	}
	return countPixels
}

func FRandomTest(count int, front *image.RGBA) int {
	countPoints := 0
	min := foo.MinY()
	blue := color.RGBA{0, 0, 255, 255}
	red := color.RGBA{255, 0, 0, 255}

	for ; count > 0; count-- {
		//Intn генерирует целые числа, для этого умножаем и делим на 300
		//x и y локальные для "рамки"
		x := float64(rand.Intn(int(rectA*1000.0))) / 1000.0
		y := float64(rand.Intn(int(rectB*1000.0))) / 1000.0
		// передаём уже глобальные значения
		ok, sign := foo.IsPointInside(x+foo.A, y+min)
		if ok {
			countPoints += sign
			front.Set(int(300*(x)), int(300*(rectB-y)), red)
		} else {
			front.Set(int(300*(x)), int(300*(rectB-y)), blue)
		}
	}
	return countPoints
}

func main() {

	// Цвет
	var color Color

	// Режим работы
	fmt.Println("Set mode (BlackSea, AzovSea, custom, func, func+):")
	var mode string
	fmt.Scan(&mode)

	var front *image.RGBA
	var bounds image.Rectangle
	var img image.Image

	if mode != "func" {
		// Открываем файл изображения
		file, err := os.Open("./image.jpeg") // Замените на путь к вашему изображению
		if err != nil {
			fmt.Println("Error with open file:", err)
			return
		}
		defer file.Close()

		// Декодируем изображение
		img0, _, err := image.Decode(file)
		if err != nil {
			fmt.Println("Error with image decoding:", err)
			return
		}
		img = img0

		// Создаём пользовательское изображение
		bounds = img.Bounds()

	}

	switch mode {
	case "BlackSea":
		var newEPS float64
		fmt.Print("ε=")
		fmt.Scan(&newEPS)
		foo.SetEPS(newEPS)

		color = BlackSea
		rectA, rectB = SEArectA, SEArectB
		front = image.NewRGBA(bounds)
	case "AzovSea":
		var newEPS float64
		fmt.Print("ε=")
		fmt.Scan(&newEPS)
		foo.SetEPS(newEPS)

		color = AzovSea
		rectA, rectB = SEArectA, SEArectB
		front = image.NewRGBA(bounds)
	case "custom":
		fmt.Println("Enter rgb-code of color (78 98 157)")
		fmt.Scan(&(color.r), &(color.g), &(color.b))

		var rectAN, rectBN float64
		fmt.Println("Set length of OX:")
		fmt.Scan(&rectAN)
		rectA = rectAN
		fmt.Println("Set length of OY:")
		fmt.Scan(&rectBN)
		rectB = rectBN
		front = image.NewRGBA(bounds)
	case "func":
		fmt.Println("Area of function:\n y = ln(x+sqrt(x^2-0.25))/2x^2\nin ab section, \na=0.5 \nb=1.7 \nε=1E-4")
		rectA = foo.B - foo.A
		rectB = foo.MaxY() - foo.MinY()
		n := image.Rectangle{image.Point{0, 0}, image.Point{int(300 * rectA), int(300 * rectB)}}
		front = image.NewRGBA(n)
	case "func+":
		fmt.Println("Area of function:\n y = ln(x+sqrt(x^2-0.25))/2x^2\nin ab section")
		fmt.Print("a=")
		var newA, newB, newEPS float64
		fmt.Scan(&newA)
		foo.SetA(newA)

		fmt.Print("b=")
		fmt.Scan(&newB)
		foo.SetB(newB)

		fmt.Print("ε=")
		fmt.Scan(&newEPS)
		foo.SetEPS(newEPS)

		rectA = foo.B - foo.A
		rectB = foo.MaxY() - foo.MinY()

		n := image.Rectangle{image.Point{0, 0}, image.Point{int(300 * rectA), int(300 * rectB)}}
		front = image.NewRGBA(n)
	}

	fmt.Println("Count of points (number/runge):")
	// Получаем количество точек
	var s string
	fmt.Scan(&s)
	count, errconv := strconv.Atoi(s)
	if errconv != nil {
		if s != "runge" && s != "serial" && s != "srunge" {
			fmt.Println("Error: invalid input")
			return
		}
		count = 1000
	}

	rectS := rectA * rectB
	var black, sum, sqrsum float64

	var try float64
	fmt.Println("Count of tryes: ")
	fmt.Scan(&try)

	for n := try; n > 0; n-- {

		switch s {
		case "runge":
			switch mode {
			case "func", "func+":
				black1 := float64(FRandomTest(count, front))
				for {

					// когда вычислили количество точек под погрешность не увеличиваем
					count *= 2

					black2 := float64(FRandomTest(count, front))
					if math.Abs(black2-black1*2)*rectS/float64(count) <= foo.EPS {
						black = black2
						// как только нашли оптимальное количество точек
						// считаем по нему (без Рунге)
						s = "d"
						//count = int(black)
						break
					}
					black1 = black2
				}

			case "BlackSea", "AzovSea", "custom":
				black1 := float64(RandomTest(count, img, color, front))
				for {

					count *= 2

					black2 := float64(RandomTest(count, img, color, front))
					if black1 == 0 || black2 == 0 {
						continue
					}
					if math.Abs(black2-black1*2)*rectS/float64(count) <= foo.EPS {
						black = black2
						s = "d"
						//count = int(black)
						break
					}
					black1 = black2
				}

			}
		default:
			switch mode {
			case "BlackSea", "AzovSea", "custom":
				black = float64(RandomTest(count, img, color, front))
			case "func", "func+":
				black = float64(FRandomTest(count, front))
			}

		}

		// суммируем число попавших точек
		sum += black

		fmt.Print(" • ")
		sqrsum += black * black

	}

	z := rectS / float64(count)

	sum /= try             // среднее количество точек
	S := sum * z           // средняя площадь
	sqrsum /= try          // среднее количество точек в квадрате
	sqrS := sqrsum * z * z // средняя площадь в квадрате

	fmt.Println("\nHits: ", sum, " / ", count)
	fmt.Println("Full area: ", rectS)

	fmt.Println("\nσ = ±", math.Sqrt(sqrS-S*S))
	fmt.Println("Result: ", S)
	if mode != "func" {
		// Сохраняем новое изображение
		outFile, err := os.Create("output.png")
		if err != nil {
			panic(err)
		}
		defer outFile.Close()
		// Записываем данные
		err = png.Encode(outFile, front)
		if err != nil {
			panic(err)
		}
	}

}
