package main

import (
	"fmt"

	"github.com/fr4nt1x/linAlg/tensor"
)

func main() {
	a := []float64{1, 2, 3, 4}
	t := tensor.New(a, 2, 2)
	fmt.Println(t)
	for i := uint(0); i < 1; i++ {
		for j := uint(0); j < 1; j++ {
			fmt.Printf("T[%d,%d] : %f", i, j, t.Get(i, j))
		}
	}
}
