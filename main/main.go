package main

import (
	"fmt"

	"github.com/fr4nt1x/linAlg/tensor"
)

func main() {
	a := []float64{1, 2, 3, 4}
	fmt.Println(tensor.New(a, 1, 4, 1, 1, 1))
}
