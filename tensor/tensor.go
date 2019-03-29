package tensor

import "github.com/fr4nt1x/linAlg/main/mathutils"

type tensor struct {
	shape   []int
	entries []float64
}

func New(entries []float64, shape ...int) tensor {
	if len(entries) != mathutils.Prod(shape) {
		panic("Wrong amoung of input for given shape.")
	}
	newShape := processShape(shape)
	t := tensor{newShape, entries}
	return t
}

func processShape(shape []int) []int {
	var newShape []int
	for _, i := range shape {
		if i != 1 {
			newShape = append(newShape, i)
		}
	}
	return newShape
}
