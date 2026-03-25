package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight float64 = 1.8
	var userWeight float64 = 100
	var BMI float64 = userWeight / math.Pow(userHeight, 2)
	fmt.Print(BMI)
}
