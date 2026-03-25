package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower = 2
	var userHeight, userWeight float64
	fmt.Print("____ Калькулятор ИМТ ____\n\n")
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	BMI := userWeight / math.Pow(userHeight, BMIPower)
	fmt.Print("Ваш ИМТ: ", BMI)
}
