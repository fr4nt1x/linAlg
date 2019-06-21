package tensor

import (
	"fmt"

	"github.com/fr4nt1x/linAlg/mathutils"
)

func processShape(shape []uint) []uint {
	var newShape []uint
	for _, i := range shape {
		if i != 1 {
			newShape = append(newShape, i)
		}
	}
	if len(newShape) == 0 {
		newShape = append(newShape, 1)
	}
	return newShape
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

func getIndicesFromVectorizedIndex(inputShape []uint, vectorizedIndex uint) []uint {
	lastIndex := len(inputShape) - 1
	indices := make([]uint, len(inputShape))
	/*TODO Not working right now*/
	for i := 0; i < len(inputShape)-1; i++ {
		indices[lastIndex-i] = mathutils.Prod(inputShape[0:lastIndex-i]) % vectorizedIndex
		vectorizedIndex -= indices[lastIndex-i]
	}

	return indices
}
