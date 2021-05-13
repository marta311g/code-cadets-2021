package main

import (
	"fmt"
	"math"
)

type taxBracket struct {
	LowerLimitInclusive float64
	UpperLimitExclusive float64
	Percentage float64
}

func calculateTax(inputValue float64, taxBrackets []taxBracket) float64 {
	var tax float64
	for index, bracket := range taxBrackets {
		if inputValue <= 0 {
			break
		}
		if index == len(taxBrackets)-1 {
			tax += inputValue * bracket.Percentage
			return tax
		}
		tax += math.Min((bracket.UpperLimitExclusive - bracket.LowerLimitInclusive), inputValue) * bracket.Percentage
		inputValue = inputValue - (bracket.UpperLimitExclusive - bracket.LowerLimitInclusive)
	}
	return tax
}

func addTaxBracket(lowerLimit, upperLimit, percentage float64) taxBracket {
	return taxBracket{ lowerLimit, upperLimit, percentage}
}

func main() {
	inputValue := 7000.0

	var taxBrackets []taxBracket
	taxBrackets = append(taxBrackets, addTaxBracket(0, 1000, 0))
	taxBrackets = append(taxBrackets, addTaxBracket(1000, 5000, 0.1))
	taxBrackets = append(taxBrackets, addTaxBracket(5000, 10000, 0.2))
	taxBrackets = append(taxBrackets, addTaxBracket(10000, -1, 0.3))

	fmt.Printf("Za ulaznu vrijednost %.2f iznos poreza je: %.2f\n", inputValue, calculateTax(inputValue, taxBrackets))
}
