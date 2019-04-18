package tensor

import (
	"fmt"

	"github.com/fr4nt1x/linAlg/mathutils"
)

//Struct Datastruct for tensor type
type Struct struct {
	shape   []uint
	entries []float64
}

//Tensor Interface methods for type tensor
type Tensor interface {
	Get([]uint) float64
	GetList([]uint) []float64
}

//New Creates new tensor
//Removes any entries of shape that are equal to 1
//Checks that number of entries matches the product of shape
func New(entries []float64, shape ...uint) Struct {
	if uint(len(entries)) != mathutils.Prod(shape) {
		panic("Wrong amoung of input for given shape.")
	}
	newShape := processShape(shape)
	t := Struct{newShape, entries}
	return t
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

//Get returns the value of the given index
func (inputTensor Struct) Get(indices []uint) float64 {
	checkIndicesInRange(inputTensor, indices)
	vectorizedIndex := uint(0)
	lastIndex := len(indices) - 1
	for i, index := range indices[:lastIndex] {
		vectorizedIndex += index * mathutils.Prod((inputTensor.shape[i+1:]))
	}
	vectorizedIndex += indices[lastIndex]
	return inputTensor.entries[vectorizedIndex]
}

func checkIndicesInRange(inputTensor Struct, indices []uint) {
	if len(indices) != len(inputTensor.shape) {
		panic("Not enough indices given for tensor.")
	}
	for i, index := range indices {
		if index > inputTensor.shape[i]-1 {
			panic(fmt.Sprintf("Index %d out of bounds: %d", i, index))
		}

	}

}
