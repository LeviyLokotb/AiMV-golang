package data

/*
	F = x1 + 2*x2 -> max

	x1 + x2 <= 5
	3*x1 + 8*x2 <= 24
	0 <= x1 <= 5
	0 <= x2 <= 3
	x1, x2 - целые
*/

// Видов на старте
var FIRSTS = 4

// Сколько особей приходит извне в каждом поколении
var OUTSIDE_STRANGERS = 2

// Шанс мутации одного гена %
var MUTATION_CHANCE = 10

// Поколений
var GENERATIONS_MAX = 10

var X1s = 5
var X2s = 3

func IsGoodSol(x1, x2 int) bool {
	return (x1+x2 <= 5) &&
		(3*x1+8*x2 <= 24)
}

func F(x1, x2 int) int {
	return x1 + 2*x2
}

var Mutate_Counter = 0
var Generations_Counter = 1
var Death_Counter = 0
var Creature_Counter = 0
