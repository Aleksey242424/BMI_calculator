package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight, userWeight float64 = 1.8, 100
	BMI := userWeight / math.Pow(userHeight, 2)
	fmt.Print(BMI)
}
