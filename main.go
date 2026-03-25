package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower = 2
	var userHeight, userWeight float64
	fmt.Println("____ Калькулятор ИМТ ____")
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	BMI := userWeight / math.Pow(userHeight/100, BMIPower)
	fmt.Printf("Ваш ИМТ: %.2f", BMI)
}
