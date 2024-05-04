package main

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// Generate some random data points
	n := 100
	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = float64(i)
		y[i] = 0.5*float64(i) + 10 + rand.NormFloat64()*5 // y = 0.5*x + 10 + noise
	}

	// Perform linear regression
	xMat := mat.NewDense(n, 2, nil)
	for i := 0; i < n; i++ {
		xMat.Set(i, 0, x[i])
		xMat.Set(i, 1, 1) // Bias term
	}
	yVec := mat.NewVecDense(n, y)

	// Calculate weights
	var weights mat.Dense
	weights.Solve(xMat, yVec)

	// Extract the slope and intercept
	slope := weights.At(0, 0)
	intercept := weights.At(1, 0)
	fmt.Printf("Slope: %.2f, Intercept: %.2f\n", slope, intercept)
}
