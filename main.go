package main

import (
	"fmt"
	"math"
)

const BMIPower = 2

func main() {
	fmt.Println("____ Калькулятор ИМТ ____")
	var userHeight, userWeight float64 = getUserInput()
	BMI := calculateIMT(userHeight, userWeight)
	status := ""
	if BMI < 16 {
		status = "Сильный дефицит массы тела !!!"
	} else if BMI > 16 && BMI <= 18.5 {
		status = "Дефицит массы тела !!"
	} else if BMI > 18.5 && BMI <= 25 {
		status = "Норма."
	} else if BMI > 25 && BMI <= 30 {
		status = "Избыточная масса !"
	} else if BMI > 30 && BMI <= 35 {
		status = "1-я степень ожирения !!"
	} else if BMI > 35 && BMI <= 40 {
		status = "2-я степень ожирения !!"
	} else {
		status = "3-я степень ожирения !!!!"
	}
	fmt.Printf("Ваш ИМТ: %.2f\nСтатус: %v", BMI, status)
}

func calculateIMT(userHeight, userWeight float64) float64 {
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
