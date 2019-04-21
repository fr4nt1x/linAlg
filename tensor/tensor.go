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
	Get([]uint) float64
	GetList([][]uint) []float64
	Set([]uint, float64)
	SetList([][]uint, []float64)
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

//Set sets the value of the given index to value
func (inputTensor Struct) Set(indices []uint, value float64) {
	checkIndicesInRange(inputTensor, indices)
	vectorizedIndex := uint(0)
	lastIndex := len(indices) - 1
	for i, index := range indices[:lastIndex] {
		vectorizedIndex += index * mathutils.Prod((inputTensor.shape[i+1:]))
	}
	vectorizedIndex += indices[lastIndex]
	inputTensor.entries[vectorizedIndex] = value
}

//GetList returns the values for the the given indices
func (inputTensor Struct) GetList(indicesList [][]uint) []float64 {
	outputList := make([]float64, len(indicesList))
	for i, indices := range indicesList {
		checkIndicesInRange(inputTensor, indices)
		outputList[i] = inputTensor.Get(indices)
	}
	return outputList
}

//SetList set the values to the the given indices
func (inputTensor Struct) SetList(indicesList [][]uint, values []float64) {
	if len(indicesList) != len(values) {
		panic("Number of values does not fit the given indices.")
	}
	for i, indices := range indicesList {
		checkIndicesInRange(inputTensor, indices)
		inputTensor.Set(indices, values[i])
	}
}
