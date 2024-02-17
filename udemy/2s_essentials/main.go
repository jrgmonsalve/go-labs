package main

import (
	"fmt"
	"math"
)

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

func profitCalculator() {
	var revenue float64
	var expenses float64
	var taxRate float64
	var ebt float64 // Earnings before tax
	var profit float64
	var ratio float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	ebt = revenue - expenses
	profit = ebt - (ebt * taxRate / 100)
	ratio = profit / revenue * 100

	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Profit ratio: ", ratio)

}

func investCalculator() {
	const inflationRate = 2.5
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
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
