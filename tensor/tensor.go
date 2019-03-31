package tensor

import (
	"fmt"

	"github.com/fr4nt1x/linAlg/main/mathutils"
)

type tensor struct {
	shape   []uint
	entries []float64
}

type Tensor interface {
	Get() tensor
}

func New(entries []float64, shape ...uint) tensor {
	if uint(len(entries)) != mathutils.Prod(shape) {
		panic("Wrong amoung of input for given shape.")
	}
	newShape := processShape(shape)
	t := tensor{newShape, entries}
	return t
}

func (inputTensor tensor) Get(indices ...uint) float64 {
	checkIndicesInRange(inputTensor, indices)
	vectorizedIndex := uint(0)
	lastIndex := len(indices)
	for i, index := range indices[:lastIndex-1] {
		vectorizedIndex += index * mathutils.Prod((inputTensor.shape[i+1:]))
	}
	vectorizedIndex += indices[lastIndex]
	return inputTensor.entries[vectorizedIndex]
}

func checkIndicesInRange(inputTensor tensor, indices []uint) {
	if len(indices) != len(inputTensor.shape) {
		panic("Not enough indices given for tensor.")
	}
	for i, index := range indices {
		if index > inputTensor.shape[i]-1 {
			panic(fmt.Sprintf("Index %d out of bounds: %d", i, index))
		}

	}

}
func processShape(shape []uint) []uint {
	var newShape []uint
	for _, i := range shape {
		if i != 1 {
			newShape = append(newShape, i)
		}
	}
	return newShape
}
