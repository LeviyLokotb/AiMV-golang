package main

import (
	"fmt"
	"math"
	foo "simpson/funcmy"
	"strconv"
)

// обычный
func big_simpson(n, h float64) float64 {
	J := (foo.B - foo.A) / (3 * n)
	// cчитаем сумму с i=0
	sum := foo.F(foo.A)
	fmt.Println("in A: ", sum)
	for i := 1.0; i < n; i++ {
		// получаем коэффициент
		k := float64(2 + (int(i)%2)*2)
		sum += k * foo.F(foo.A+i*h)
	}
	// последний элемент
	sum += foo.F(foo.B)
	fmt.Println("in B: ", foo.F(foo.B))
	J *= sum
	return J
}

// для одного интервала
func simpson(a, b float64) float64 {
	center := (a + b) / 2.0
	h := b - a
	return (h / 6.0) * (foo.F(a) + 4*foo.F(center) + foo.F(b))
}

// адаптивный (автоматическое деление на неравные интервалы)
func adapt_simpson(a, b, eps float64) float64 {
	// находим интеграл на всём интервале
	J := simpson(a, b)
	// а потом для каждой половины
	center := (a + b) / 2
	J1 := simpson(a, center)
	J2 := simpson(center, b)
	Jsum := J1 + J2
	// если соответствует точности
	if math.Abs(Jsum-J)/15 <= eps {
		return Jsum
	}
	// иначе повторяем рекурсивно
	return adapt_simpson(a, center, eps) + adapt_simpson(center, b, eps)
}

func main() {
	var s string
	fmt.Print("Enter ε:\nε=")
	fmt.Scan(&s)
	eps, err := strconv.ParseFloat(s, 64)
	if s != "d" && s != "def" && s != "default" && s != "" {
		if err != nil {
			fmt.Println(err)
			return
		}
		foo.SetEPS(eps)
	}

	// Режим работы
	fmt.Println("Set mode (default/auto/adaptive):")
	var mode string
	fmt.Scan(&mode)
	switch mode {
	case "d", "def":
		mode = "default"
	case "a", "adapt":
		mode = "adaptive"
	}

	var sectN, sectH float64
	var J float64
	switch mode {
	case "default":
		// число секций
		fmt.Print("Amount of sections: ")
		fmt.Scan(&sectN)
		// в Go строгая типизация, после деления округлится
		sectN = (sectN / 2) * 2
		if sectN <= 2 {
			J = simpson(foo.A, foo.B)
			break
		}
		// длина 1 секции
		sectH = (foo.B - foo.A) / sectN
		J = big_simpson(sectN, sectH)
	case "auto":
		sectN = 2
		sectH = (foo.B - foo.A) / sectN
		J1 := big_simpson(sectN, sectH)
		for {
			sectN *= 2
			sectH /= 2
			J2 := big_simpson(sectN, sectH)
			// k=4, 2^k-1=15
			if (J2-J1)/15 <= foo.EPS {
				J = J2
				break
			}
			J1 = J2
		}
		fmt.Println("N = ", sectN)
	case "adaptive":
		J = adapt_simpson(foo.A, foo.B, foo.EPS)
	}
	fmt.Println("\nResult: ", J)

}
