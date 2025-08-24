package fu

// Функции, описывающие СНАУ:
func F(x float64, num int) float64 {
	switch num {
	case 1:
		return 1
	case 2:
		return 1
	case 3:
		return 1
	case 4:
		return 1
	case 5:
		return 1
	case 6:
		return 1
	default:
		panic("Нет такого уравнения!")
	}
}

// Производная (считаем вручную, к сожалению)
func Dx(x float64, num int) float64 {
	x1 = x+0.01
	x2 = x-0.01
	return (F(x2,))
}

