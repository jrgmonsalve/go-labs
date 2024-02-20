package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

const inflationRate = 2.5

func getValueFormFile() (float64, error) {
	data, err := os.ReadFile("profit.txt")
	if err != nil {
		return 0, errors.New("Profit file not found!!!!!")
	}
	value, _ := strconv.ParseFloat(string(data), 64)
	return value, nil
}

func wirteValueToFile(value float64) {
	valueText := fmt.Sprintf("%.2f", value)
	os.WriteFile("profit.txt", []byte(valueText), 0644)
}

func profitCalculator() {
	var revenue float64
	var expenses float64
	var taxRate float64
	var ebt float64 // Earnings before tax
	var profit float64
	var ratio float64

	lastProfit, err := getValueFormFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------")
	}
	fmt.Println("----->>> Last profit: ", lastProfit)

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	ebt = revenue - expenses
	profit = ebt - (ebt * taxRate / 100)
	ratio = profit / revenue * 100
	wirteValueToFile(profit)
	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Profit ratio: ", ratio)

}

func investCalculator() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Println("")

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Println("")

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)
	fmt.Println("")

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)
	fmt.Printf("Future value: %.2f\n", futureValue) // f is the float and .2 is the number of decimal places

	fmt.Println(futureRealValue)
}

func main() {
	var option int
	fmt.Println("Choose an option: 1) Invest calculator 2) Profit calculator 3) Exit")
	fmt.Scan(&option)

	switch option {
	case 1:
		investCalculator()
	case 2:
		profitCalculator()
	}

}
