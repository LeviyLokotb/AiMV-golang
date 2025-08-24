package main

import (
	"fmt"
	"minnel2/fx"
	"minnel2/graphdraw"
)

// Метод Ньютона (2 порядка, с использованием матрицы Гессе)
func main() {

	x, y := fx.Newton(fx.F)
	zmin := fx.F(x, y)
	fmt.Println("Минимум: (", x, " ; ", y, ")")

	img := graphdraw.DrawFunction3D(fx.F, zmin, "Функция")
	graphdraw.DrawPoint(x, y, img, 0, 0, 0)
	graphdraw.Save(img, "graphic")
}
