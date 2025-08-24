package main

import (
	"splinErmit/fx"
	"splinErmit/graphdraw"
)

// Сплайн Эрмита -  задан: n точек и I производные
//

func main() {

	// Рисуем графики
	fu1 := graphdraw.DrawFunction(fx.F1, "e^x")
	fu2 := graphdraw.DrawFunction(fx.F2, "e^(-x)")
	fu3 := graphdraw.DrawFunction(fx.F3, "sinh(x)")
	fu4 := graphdraw.DrawFunction(fx.F4, "cosh(x)")
	fu5 := graphdraw.DrawFunction(fx.F5, "sin(x)")
	fu6 := graphdraw.DrawFunction(fx.F6, "cos(x)")
	fu7 := graphdraw.DrawFunction(fx.F7, "ln(x)")

	// Сохраняем изображения
	graphdraw.Save(fu1, "1 e^x")
	graphdraw.Save(fu2, "2 e^(-x)")
	graphdraw.Save(fu3, "3 sinh(x)")
	graphdraw.Save(fu4, "4 cosh(x)")
	graphdraw.Save(fu5, "5 sin(x)")
	graphdraw.Save(fu6, "6 cos(x)")
	graphdraw.Save(fu7, "7 ln(x)")

	// Строим и рисуем сплайн
	Sfu1 := graphdraw.DrawSpline(fx.CubeSpline, fx.F1, "e^x", fu1)
	Sfu2 := graphdraw.DrawSpline(fx.CubeSpline, fx.F2, "e^(-x)", fu2)
	Sfu3 := graphdraw.DrawSpline(fx.CubeSpline, fx.F3, "sinh(x)", fu3)
	Sfu4 := graphdraw.DrawSpline(fx.CubeSpline, fx.F4, "cosh(x", fu4)
	Sfu5 := graphdraw.DrawSpline(fx.CubeSpline, fx.F5, "sin(x)", fu5)
	Sfu6 := graphdraw.DrawSpline(fx.CubeSpline, fx.F6, "cos(x)", fu6)
	Sfu7 := graphdraw.DrawSpline(fx.CubeSpline, fx.F7, "ln(x)", fu7)

	// Сохраняем изображения
	graphdraw.Save(Sfu1, "1S e^x")
	graphdraw.Save(Sfu2, "2S e^(-x)")
	graphdraw.Save(Sfu3, "3S sinh(x)")
	graphdraw.Save(Sfu4, "4S cosh(x)")
	graphdraw.Save(Sfu5, "5S sin(x)")
	graphdraw.Save(Sfu6, "6S cos(x)")
	graphdraw.Save(Sfu7, "7S ln(x)")

}
