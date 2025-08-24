package processing

import (
	"fmt"
	"math"
)

// Является ли решением
func IsSolution(sol [][]float64, sols map[int]int) bool {
	l := len(sol[1])
	if sol[1][l-1] != 0 {
		return false
	}
	for i := 0; i < 6; i++ {
		if sol[1][i] < 0 {
			return false
		}
		/*
			for _, j := range sols {
				if i == j {
					if sol[i][l-1] != float64(int(sol[i][l-1])) {
						return false
					}
				}
			}
		*/

	}

	return true
}

func FindMinInCol(m []float64) (index int, min float64) {
	min = math.Inf(1)
	index = -1
	for i := 0; i < len(m)-1; i++ {
		if m[i] < min && m[i] < 0 {
			min = m[i]
			index = i
		}
	}
	return
}

func FindMinInRow(m []float64) (index int, min float64) {
	min = math.Inf(1)
	index = -1
	for i := 2; i < len(m); i++ {
		if m[i] < min {
			min = m[i]
			index = i
		}
	}
	return
}

func PrintMx(m [][]float64) {
	for i := range m {
		fmt.Print("[")
		for j := range m[i] {
			el := m[i][j]
			if el-float64(int(el)) == 0.0 {
				fmt.Print(" ", el, " ")
			} else {
				fmt.Printf(" %.1f ", el)
			}

		}
		fmt.Println("]")
	}
}
