package tensor

import (
	"testing"

	"github.com/fr4nt1x/linAlg/mathutils"
)

func TestProcessShape(t *testing.T) {
	tables := []struct {
		inputShapes    []uint
		expectedShapes []uint
	}{
		{[]uint{1, 2, 3}, []uint{2, 3}},
		{[]uint{1, 1, 1, 1}, []uint{1}},
		{[]uint{2, 1, 3, 1}, []uint{2, 3}},
	}
	for _, table := range tables {
		actualShapes := processShape(table.inputShapes)
		for i, expectedValue := range table.expectedShapes {

			if expectedValue != actualShapes[i] {
				t.Errorf("Actual shape was incorrect, got %d, want %d at index %d", actualShapes[i], expectedValue, i)
			}
		}
	}
}

func TestGet(t *testing.T) {
	indexX := uint(2)
	indexY := uint(2)
	indexZ := uint(3)
	testTensor := createTestTensor([]uint{indexX, indexY, indexZ})
	expectedValue := uint(0)
	var actualValue uint
	for i := uint(0); i < indexX; i++ {
		for j := uint(0); j < indexY; j++ {
			for k := uint(0); k < indexZ; k++ {
				actualValue = uint(testTensor.Get([]uint{i, j, k}))
				if expectedValue != actualValue {
					t.Errorf("Wrong value returned, Got %d, Expected %d, for index (%d,%d,%d)", actualValue, expectedValue, i, j, k)
				}
				expectedValue++
			}
		}
	}
}

func TestGetList(t *testing.T) {
	indexX := uint(2)
	indexY := uint(2)
	indexZ := uint(3)
	testTensor := createTestTensor([]uint{indexX, indexY, indexZ})
	indices := [][]uint{
		[]uint{0, 0, 1},
		[]uint{0, 1, 0},
		[]uint{0, 1, 1},
		[]uint{1, 0, 2},
	}
	expectedValues := []float64{1, 3, 4, 8}
	actualValues := testTensor.GetList(indices)
	for i, v := range expectedValues {
		if v != actualValues[i] {
			t.Errorf("Wrong value returned, expected %f, got %f at index %d", v, actualValues[i], i)
		}
	}
}

func createTestTensor(indices []uint) Tensor {
	numelTensor := mathutils.Prod(indices)
	a := []float64{}
	for i := uint(0); i < numelTensor; i++ {
		a = append(a, float64(i))
	}
	t := New(a, indices...)
	return t
}
