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
	LowerLimitInclusive float64
	UpperLimitExclusive float64
	Percentage float64
}

func CalculateTax(inputValue float64, taxBrackets []TaxBracket) (float64, error) {
	var tax float64

	if inputValue < 0 {
		return 0, errors.New("The input value should be greater than zero.")
	}

	for index, bracket := range taxBrackets {
		if inputValue <= 0 {
			break
		}
		if index == len(taxBrackets)-1 {
			tax += inputValue * bracket.Percentage
			return tax, nil
		}
		tax += math.Min((bracket.UpperLimitExclusive - bracket.LowerLimitInclusive), inputValue) * bracket.Percentage
		inputValue = inputValue - (bracket.UpperLimitExclusive - bracket.LowerLimitInclusive)
	}
	return tax, nil
}

func makeTaxBracket(lineElements []string) TaxBracket {
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
	return TaxBracket{ lowerLimit, upperLimit, percentage}
}

func GetTaxBrackets(file string) ([]TaxBracket, error) {
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

	for _, line := range text {
		taxBrackets = append(taxBrackets, makeTaxBracket(strings.Split(line, ";")))
	}
	return taxBrackets, nil
}
