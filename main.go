package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight, userWeight float64
	fmt.Println("____ Калькулятор ИМТ ____")
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	BMI := calculateIMT(userHeight, userWeight)
	fmt.Printf("Ваш ИМТ: %.2f", BMI)
}

func calculateIMT(userHeight, userWeight float64) float64 {
	const BMIPower = 2
	result := userWeight / math.Pow(userHeight/100, BMIPower)
	return result
}
