package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("____ Калькулятор ИМТ ____")
	var userHeight, userWeight float64 = getUserInput()
	BMI := calculateIMT(userHeight, userWeight)
	fmt.Printf("Ваш ИМТ: %.2f", BMI)
}

func calculateIMT(userHeight, userWeight float64) float64 {
	const BMIPower = 2
	result := userWeight / math.Pow(userHeight/100, BMIPower)
	return result
}

func getUserInput() (float64, float64) {
	var userHeight, userWeight float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}
