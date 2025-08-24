package genetic

import (
	"genetic/data"
	"math/rand"
	"time"
)

type Genom struct {
	x1      int
	x2      int
	fitness int
}

// Новый случайный геном
func NewGenom() *Genom {
	data.Creature_Counter++
	rand.New(rand.NewSource(time.Now().UnixNano()))
	lilith := new(Genom)
	lilith.x1 = rand.Intn(data.X1s)
	//fmt.Println(lilith.x1)
	lilith.x2 = rand.Intn(data.X2s)
	//fmt.Println(lilith.x2)
	for !lilith.IsAlive() {
		lilith.x1 = rand.Intn(data.X1s)

		lilith.x2 = rand.Intn(data.X2s)
	}

	lilith.fitness = data.F(lilith.x1, lilith.x2)
	return lilith
}

// Скрещивание
func Crossover(adam, eva *Genom) []*Genom {
	data.Creature_Counter += 4
	aa := new(Genom)
	aa.x1, aa.x2, aa.fitness = adam.x1, adam.x2, adam.fitness

	ae := new(Genom)
	ae.x1, ae.x2 = adam.x1, eva.x2
	ae.fitness = data.F(ae.x1, ae.x2)

	ea := new(Genom)
	ea.x1, ea.x2 = eva.x1, adam.x2
	ea.fitness = data.F(ea.x1, ea.x2)

	ee := new(Genom)
	ee.x1, ee.x2, ee.fitness = eva.x1, eva.x2, eva.fitness

	// Мутации
	aa.Mutate()
	ae.Mutate()
	ea.Mutate()
	ee.Mutate()
	return []*Genom{aa, ae, ea, ee}
}

// Мутация с шансом
func (gen *Genom) Mutate() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	mutant := Genom{gen.x1, gen.x2, gen.fitness}
	if rand.Intn(99)+1 <= data.MUTATION_CHANCE {
		data.Mutate_Counter++
		switch rand.Intn(2) {
		case 0:
			mutant.x1 -= 1
		case 1:
			mutant.x1 += 1
		case 2:
			mutant.x1 = rand.Intn(data.X1s)
		}
	}
	if rand.Intn(99)+1 <= data.MUTATION_CHANCE {
		data.Mutate_Counter++
		switch rand.Intn(2) {
		case 0:
			mutant.x2 -= 1
		case 1:
			mutant.x2 += 1
		case 2:
			mutant.x2 = rand.Intn(data.X2s)
		}
	}
	mutant.fitness = data.F(mutant.x1, mutant.x2)
	gen.x1, gen.x2, gen.fitness = mutant.x1, mutant.x2, mutant.fitness
}

func (gen *Genom) IsAlive() bool {
	good := data.IsGoodSol(gen.x1, gen.x2)
	if !good {
		data.Death_Counter++
	}
	return good
}

func BubbleSortGen(g []*Genom) []*Genom {
	n := make([]*Genom, 0)
	for i := range g {
		if g[i].IsAlive() {
			n = append(n, g[i])
		}
	}
	for {
		count := 0
		for i := 0; i < len(g)-2; i++ {
			if n[i+1].fitness < n[i].fitness {
				n[i], n[i+1] = n[i+1], n[i]
				count++
			}
		}
		if count == 0 {
			break
		}
	}
	return n
}

func (gen *Genom) GetInfo() (int, int, int) {
	if gen == nil {
		return 0, 0, -1
	}
	return gen.x1, gen.x2, gen.fitness
}
