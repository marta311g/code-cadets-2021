package tax

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type TaxBracket struct {
	LowerThreshold float64
	UpperThreshold float64
	Percentage float64
}

func CalculateTax(inputValue float64, taxBrackets []TaxBracket) (float64, error) {
	var tax float64
	var remainingValue = inputValue

	if inputValue < 0 {
		return 0, errors.New("The input value should be greater than zero.")
	}

	for index, bracket := range taxBrackets {
		if remainingValue <= 0 {
			break
		}
		if index == len(taxBrackets)-1 {
			tax += remainingValue * bracket.Percentage
			return tax, nil
		}
		tax += math.Min((bracket.UpperThreshold - bracket.LowerThreshold), remainingValue) * bracket.Percentage
		remainingValue = remainingValue - (bracket.UpperThreshold - bracket.LowerThreshold)
	}
	return tax, nil
}

func makeTaxBracket(lowerTresholdString, upperThresholdString, percentageString string) TaxBracket {
	lowerTreshold, err := strconv.ParseFloat(lowerTresholdString, 64)
	if err != nil {
		log.Fatal( errors.WithMessage(err, "invalid lower limit"))
	}
	upperThreshold, err := strconv.ParseFloat(upperThresholdString, 64)
	if err != nil {
		log.Fatal( errors.WithMessage(err, "invalid upper limit"))
	}
	percentage, err := strconv.ParseFloat(percentageString, 64)
	if err != nil {
		log.Fatal( errors.WithMessage(err, "invalid percentage"))
	}

	if upperThreshold < lowerTreshold && upperThreshold != -1 {
		upperThreshold, lowerTreshold = lowerTreshold, upperThreshold
	}

	return TaxBracket{ lowerTreshold, upperThreshold, percentage}
}

func GetTaxBracketsFromFile(file string) ([]TaxBracket, error) {
	var taxBrackets []TaxBracket

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	defer f.Close()

	for index, line := range text {
		if !strings.HasPrefix(line, "#") {
			lowerTreshold := "0"
			if index > 1 {
				lowerTreshold = strings.Split(text[index-1], ";")[0]
			}
			upperThreshold := strings.Split(line, ";")[0]
			percentage := strings.Split(line, ";")[1]

			taxBrackets = append(taxBrackets, makeTaxBracket(lowerTreshold, upperThreshold, percentage))
		}
	}
	return taxBrackets, nil
}
