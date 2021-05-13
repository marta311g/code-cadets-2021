package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
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

func makeTaxBracket(lineElements []string) taxBracket {
	lowerLimit, err := strconv.ParseFloat(lineElements[0], 64)
	if err != nil {
		log.Fatal( errors.WithMessage(err, "invalid lower limit"), )
	}
	upperLimit, err := strconv.ParseFloat(lineElements[1], 64)
	if err != nil {
		log.Fatal( errors.WithMessage(err, "invalid lower limit"), )
	}
	percentage, err := strconv.ParseFloat(lineElements[2], 64)
	if err != nil {
		log.Fatal( errors.WithMessage(err, "invalid lower limit"), )
	}
	return taxBracket{ lowerLimit, upperLimit, percentage}
}

func getTaxBrackets() []taxBracket {
	var taxBrackets []taxBracket

	f, err := os.Open("./homework_1/brackets.txt")
	if err != nil {
		log.Fatal( errors.WithMessage(err, "error while opening a file"), )
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	defer f.Close()

	for _, line := range text {
		taxBrackets = append(taxBrackets, makeTaxBracket(strings.Split(line, ";")))
	}
	return taxBrackets
}

func main() {
	var inputValue float64
	fmt.Print("Please, enter input value: ")
	fmt.Scanf("%f", &inputValue)

	var taxBrackets = getTaxBrackets()

	fmt.Printf("Za ulaznu vrijednost %.2f iznos poreza je: %.2f\n", inputValue, calculateTax(inputValue, taxBrackets))
}
