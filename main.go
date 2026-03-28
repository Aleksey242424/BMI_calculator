package main

import (
	"errors"
	"fmt"
	"math"
)

const BMIPower = 2

func main() {
	fmt.Println("____ Калькулятор ИМТ ____")
	for {
		var userHeight, userWeight float64 = getUserInput()
		BMI, err := calculateIMT(userHeight, userWeight)
		if err != nil {
			fmt.Println(err)
			continue
		}
		status := getStatus(BMI)
		fmt.Printf("Ваш ИМТ: %.2f\nСтатус: %v\n", BMI, status)
		fmt.Println("Вы хотите сделать ещё расчёт (y/n): ")
		userChoise := checkRepeatCalculation()
		if !userChoise {
			break
		}
	}
}

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}

func getStatus(BMI float64) string {
	status := ""
	switch {
	case BMI < 16:
		status = "Сильный дефицит массы тела !!!"
	case BMI > 16 && BMI <= 18.5:
		status = "Дефицит массы тела !!"
	case BMI > 18.5 && BMI <= 25:
		status = "Норма."
	case BMI > 25 && BMI <= 30:
		status = "Избыточная масса !"
	case BMI > 30 && BMI <= 35:
		status = "1-я степень ожирения !!"
	case BMI > 35 && BMI <= 40:
		status = "2-я степень ожирения !!"
	default:
		status = "3-я степень ожирения !!!!"
	}
	return status
}

func calculateIMT(userHeight, userWeight float64) (float64, error) {
	if userHeight <= 0 || userWeight <= 0 {
		return 0, errors.New("Не указан вес или высота")
	}
	result := userWeight / math.Pow(userHeight/100, BMIPower)
	return result, nil
}

func getUserInput() (float64, float64) {
	var userHeight, userWeight float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}
