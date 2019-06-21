package tensor

import (
	"github.com/fr4nt1x/linAlg/mathutils"
)

//Struct Datastruct for tensor type
type Struct struct {
	shape   []uint
	entries []float64
}

//Tensor Interface methods for type tensor
type Tensor interface {
	GetShape() []uint
	Get([]uint) float64
	GetList([][]uint) []float64
	Set([]uint, float64)
	SetList([][]uint, []float64)
	Dot(*Tensor, uint, uint) Tensor
}

//New Creates new tensor
//Removes any entries of shape that are equal to 1
//Checks that number of entries matches the product of shape
func New(entries []float64, shape ...uint) *Struct {
	if uint(len(entries)) != mathutils.Prod(shape) {
		panic("Wrong amoung of input for given shape.")
	}
	newShape := processShape(shape)
	t := Struct{newShape, entries}
	return &t
}

//Get returns the value of the given index
func (inputTensor *Struct) Get(indices []uint) float64 {
	checkIndicesInRange(*inputTensor, indices)
	vectorizedIndex := uint(0)
	lastIndex := len(indices) - 1
	for i, index := range indices[:lastIndex] {
		vectorizedIndex += index * mathutils.Prod((inputTensor.shape[i+1:]))
	}
	vectorizedIndex += indices[lastIndex]
	return inputTensor.entries[vectorizedIndex]
}

//Set sets the value of the given index to value
func (inputTensor *Struct) Set(indices []uint, value float64) {
	checkIndicesInRange(*inputTensor, indices)
	vectorizedIndex := uint(0)
	lastIndex := len(indices) - 1
	for i, index := range indices[:lastIndex] {
		vectorizedIndex += index * mathutils.Prod((inputTensor.shape[i+1:]))
	}
	vectorizedIndex += indices[lastIndex]
	inputTensor.entries[vectorizedIndex] = value
}

//GetList returns the values for the the given indices
func (inputTensor *Struct) GetList(indicesList [][]uint) []float64 {
	outputList := make([]float64, len(indicesList))
	for i, indices := range indicesList {
		checkIndicesInRange(*inputTensor, indices)
		outputList[i] = inputTensor.Get(indices)
	}
	return outputList
}

//SetList set the values to the the given indices
func (inputTensor *Struct) SetList(indicesList [][]uint, values []float64) {
	if len(indicesList) != len(values) {
		panic("Number of values does not fit the given indices.")
	}
	for i, indices := range indicesList {
		checkIndicesInRange(*inputTensor, indices)
		inputTensor.Set(indices, values[i])
	}
}

//GetShape returns the tensor shape
func (inputTensor *Struct) GetShape() []uint {
	return inputTensor.shape
}

//Dot Calculate the dot product for two Tensors
func (inputTensor *Struct) Dot(secondInputTensor *Tensor, firstAxis uint, secondAxis uint) Tensor {
	firstShape := inputTensor.GetShape()
	secondShape := (*secondInputTensor).GetShape()
	var axisLength uint
	if firstShape[firstAxis] != secondShape[secondAxis] {
		panic("The given Axes cannot be used for contraction.")
	} else {
		axisLength = firstShape[firstAxis]
	}

	newShapeLength := len(firstShape) + len(secondShape) - 2
	newShape := make([]uint, newShapeLength)
	newIndex := 0
	for i := uint(0); i < uint(len(firstShape)); i++ {
		if i == firstAxis {
			continue
		} else {
			newShape[newIndex] = firstShape[i]
			newIndex++
		}
	}

	for i := uint(0); i < uint(len(secondShape)); i++ {
		if i == secondAxis {
			continue
		} else {
			newShape[newIndex] = secondShape[i]
			newIndex++
		}
	}

	resultTensorLength := mathutils.Prod(newShape)
	resultTensor := New(make([]float64, resultTensorLength), newShape...)
	var indicesList []uint

	for i := uint(0); i < resultTensorLength; i++ {
		indicesList = getIndicesFromVectorizedIndex(newShape, i)
		indicesFirstTensor := make([]uint, len(firstShape)-1)
		indicesSecondTensor := make([]uint, len(secondShape)-1)
		copy(indicesFirstTensor, indicesList[0:len(firstShape)-1])
		copy(indicesSecondTensor, indicesList[len(firstShape)-1:])

		/*Prepare insertion of indices*/
		indicesFirstTensor = append(indicesFirstTensor, 0 /* use the zero value of the element type */)
		copy(indicesFirstTensor[firstAxis+1:], indicesFirstTensor[firstAxis:])

		indicesSecondTensor = append(indicesSecondTensor, 0 /* use the zero value of the element type */)
		copy(indicesSecondTensor[secondAxis+1:], indicesSecondTensor[secondAxis:])

		newValue := float64(0)
		for sumIndex := uint(0); sumIndex < axisLength; sumIndex++ {
			indicesFirstTensor[firstAxis] = sumIndex
			indicesSecondTensor[secondAxis] = sumIndex
			newValue += inputTensor.Get(indicesFirstTensor) * (*secondInputTensor).Get(indicesSecondTensor)
		}
		resultTensor.Set(indicesList, newValue)
	}
	return resultTensor
}
