package data

/*
	F = x1 + 2*x2 -> max

	x1 + x2 <= 5
	3*x1 + 8*x2 <= 24
	0 <= x1 <= 5
	0 <= x2 <= 3
	x1, x2 - целые
*/

// Муравьёв (итерации поиска решения без изменения феромонов)
var ANTS = 4

// Ходок (итерации поиска решения с изменением феромонов)
var ITER = 4

// Влияние феромона на вероятность (степень)
var FeromonPOWER = 1.0

// Коэффициент усиления феромонов
var FeromonUP = 1.0

// Коэффициент испарения феромонов
var FeromonDOWN = 0.5

var X1s = []int{0, 1, 2, 3, 4, 5, 6}
var X2s = []int{0, 1, 2, 3}

func IsGoodSol(x1, x2 int) bool {
	return (x1+x2 <= 5) &&
		(3*x1+8*x2 <= 24)
}

func F(x1, x2 int) int {
	return x1 + 2*x2
}
