package main

import (
	"fmt"
	// "math"
)

func main() {

	var revenue, expenses, tax_rate float64
	fmt.Print("Your revenue: ")
	fmt.Scan(&revenue)
	revenue = rev(revenue)
	fmt.Println("Your revenue is :", revenue)

	fmt.Print("Your expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax rate: ")
	fmt.Scan(&tax_rate)

	// ebt := revenue - expenses

	var ebt, profit, ratio = calev(revenue, expenses, tax_rate)
	fmt.Print("Your ebt is ")
	fmt.Println(ebt)
	// fmt.Println()

	// profit := ebt * (1 - tax_rate / 100)
	
	fmt.Print("Your profit is ")
	fmt.Print(profit)
	fmt.Println()

	// ratio := profit / ebt
	fmt.Print("The ratio is ")
	fmt.Print(ratio)


}

func rev(revenue float64) float64 {
	return revenue
}

func calev(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - tax_rate/100)
	ratio := profit / ebt


	return ebt, profit, ratio
}
