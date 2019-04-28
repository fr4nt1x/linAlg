package tensor

import (
	"testing"

	"github.com/fr4nt1x/linAlg/mathutils"
)

func TestDot(t *testing.T) {
	indexX := uint(2)
	indexY := uint(2)
	indexZ := uint(3)
	testTensor1 := createTestTensor([]uint{indexX, indexY, indexZ})
	testTensor2 := createTestTensor([]uint{indexX, indexY, indexZ})

	testTensor1.Dot(&testTensor2, 1, 1)
}

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

func TestGetShape(t *testing.T) {
	indexX := uint(2)
	indexY := uint(2)
	indexZ := uint(3)
	testTensor := createTestTensor([]uint{indexX, indexY, indexZ})
	expectedValue := []uint{2, 2, 3}

	actualValue := testTensor.GetShape()
	for i := 0; i < 3; i++ {
		if actualValue[i] != expectedValue[i] {
			t.Errorf("Wrong shape returned, Got %d, Expected %d, for index %d", actualValue[i], expectedValue[i], i)
		}
	}
}

func TestCheckIndicesInRangeFailingTooManyIndices(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	indexX := uint(2)
	indexY := uint(2)
	indexZ := uint(3)
	testStruct := createZeroTestStruct([]uint{indexX, indexY, indexZ})
	checkIndicesInRange(*testStruct, []uint{1, 1, 1, 1})
}
func TestCheckIndicesInRangeFailingIndexOutOfBounds(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	indexX := uint(2)
	indexY := uint(2)
	indexZ := uint(3)
	testStruct := createZeroTestStruct([]uint{indexX, indexY, indexZ})
	indices := []uint{0, 0, 0}
	for i := 0; i < 3; i++ {
		indices[i] = 4
		checkIndicesInRange(*testStruct, indices)
		indices[i] = 0
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

func TestSet(t *testing.T) {
	indexX := uint(2)
	indexY := uint(2)
	indexZ := uint(3)
	testTensor := createZeroTestTensor([]uint{indexX, indexY, indexZ})
	expectedValue := uint(2)
	var actualValue uint
	for i := uint(0); i < indexX; i++ {
		for j := uint(0); j < indexY; j++ {
			for k := uint(0); k < indexZ; k++ {
				testTensor.Set([]uint{i, j, k}, 2)
				actualValue = uint(testTensor.Get([]uint{i, j, k}))
				if expectedValue != actualValue {
					t.Errorf("Wrong value returned, Got %d, Expected %d, for index (%d,%d,%d)", actualValue, expectedValue, i, j, k)
				}
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

func TestSetList(t *testing.T) {
	indexX := uint(2)
	indexY := uint(2)
	indexZ := uint(3)
	testTensor := createZeroTestTensor([]uint{indexX, indexY, indexZ})
	indices := [][]uint{
		[]uint{0, 0, 1},
		[]uint{0, 1, 0},
		[]uint{0, 1, 1},
		[]uint{1, 0, 2},
	}
	expectedValues := []float64{1, 3, 4, 8}
	testTensor.SetList(indices, expectedValues)
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

func createZeroTestTensor(indices []uint) Tensor {
	numelTensor := mathutils.Prod(indices)
	a := []float64{}
	for i := uint(0); i < numelTensor; i++ {
		a = append(a, float64(0))
	}
	t := New(a, indices...)
	return t
}

func createZeroTestStruct(indices []uint) *Struct {
	numelTensor := mathutils.Prod(indices)
	a := []float64{}
	for i := uint(0); i < numelTensor; i++ {
		a = append(a, float64(0))
	}
	t := New(a, indices...)
	return t
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
