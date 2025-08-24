package main

import (
	"fmt"
	"minnelbezogr/fx"
	"minnelbezogr/graphdraw"
)

func main() {

	xHistory, yHistory := fx.SteepestDescent(fx.F)

	xmin, ymin := xHistory[len(xHistory)-1], yHistory[len(yHistory)-1]
	zmin := fx.F(xmin, ymin)
	fmt.Println("Минимум: (", xmin, " ; ", ymin, ")")

	img := graphdraw.DrawFunction3D(fx.F, zmin, "Функция")
	graphdraw.Save(img, "graphic")
	img2 := graphdraw.Trajectory(img, xHistory, yHistory, "С траекторией")
	graphdraw.Save(img2, "trajjectory")
}
