package main

import (
	"Furje/fx"
	"Furje/graphdraw"
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("--=[ Разложение в ряд Фурье ]=--")

	image1 := graphdraw.DrawFunction(fx.F, "Разложение в ряд Фурье")
	graphdraw.Save(image1, "1-graphik")

	fx.ChetOrNot()

	fmt.Print("\nПериод (*π): ")
	var Ts string
	fmt.Scan(&Ts)
	if Ts != "d" {
		te, _ := strconv.ParseFloat(Ts, 64)
		fx.T *= te / 2.0
	}

	fx.Wre()

	fmt.Print("\nСтепень: ")
	var Ms string
	fmt.Scan(&Ms)
	if Ms != "d" {
		fx.M, _ = strconv.ParseFloat(Ms, 64)
	}

	image2 := graphdraw.DrawFunction2(fx.Furje, "Разложение в ряд Фурье", image1)
	graphdraw.Save(image2, "2-furje")
}
