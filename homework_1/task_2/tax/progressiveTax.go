package tax

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type TaxBracket struct {
	LowerThreshold float64
	UpperThreshold float64
	Percentage     float64
}


func CalculateTax(inputValue float64, fileName string) (float64, error) {
	var tax float64
	var remainingValue = inputValue

	taxBrackets, err := getTaxBracketsFromFile(fileName)
	if err != nil {
		return 0, errors.Wrap(err, "error while creating tax brackets")
	}

	for index, bracket := range taxBrackets {
		if remainingValue <= 0 {
			break
		}
		if index == len(taxBrackets)-1 {
			tax += remainingValue * bracket.Percentage
			return tax, nil
		}
		tax += math.Min((bracket.UpperThreshold-bracket.LowerThreshold), remainingValue) * bracket.Percentage
		remainingValue = remainingValue - (bracket.UpperThreshold - bracket.LowerThreshold)
	}
	return tax, nil
}

func makeSortedTaxBrackets(text []string) ([]TaxBracket, error) {
	var taxBrackets []TaxBracket
	noOfLines := 0
	for _, line := range text {
		if !strings.HasPrefix(line, "#") {
			noOfLines += 1
			elements := strings.Split(line, ";")
			if len(elements) != 2 {
				return nil, errors.New("file not formatted correctly")
			}

			upperThreshold, _ := strconv.ParseFloat(elements[0], 64)

			lowerTreshold := 0.0
			if noOfLines > 1 {
				lowerTreshold = taxBrackets[noOfLines-2].UpperThreshold
			}
			percentage, err := strconv.ParseFloat(elements[1], 64)
			if err != nil {
				return nil, errors.Wrap(err, "invalid percentage")
			}

			taxBrackets = append(taxBrackets, TaxBracket{
				LowerThreshold: lowerTreshold,
				UpperThreshold: upperThreshold,
				Percentage:     percentage,
			})
		}
	}
	return taxBrackets, nil
}

func getTaxBracketsFromFile(file string) ([]TaxBracket, error) {
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

	var sortedLines []string
	noOfLines := 0
	for _, line := range text {
		if !strings.HasPrefix(line, "#") {
			noOfLines += 1
			elements := strings.Split(line, ";")
			if len(elements) != 2 {
				return nil, errors.New("file not formatted correctly")
			}

			upperThreshold, err := strconv.ParseFloat(elements[0], 64)
			if err != nil {
				return nil, errors.Wrap(err, "invalid upper limit")
			}

			for index, sortedLine := range sortedLines {
				if len(sortedLines) == 0 {
					sortedLines = append(sortedLines, line)
					break
				}
				savedUpperTreshold, _ := strconv.ParseFloat(strings.Split(sortedLine, ";")[0], 64)
				if (upperThreshold < savedUpperTreshold || savedUpperTreshold == -1) && upperThreshold != -1 {
					sortedLines = append(sortedLines[:index+1], sortedLines[index:]...)
					sortedLines[index] = line
					break
				}
				if index == len(sortedLines)-1 || upperThreshold == -1 {
					sortedLines = append(sortedLines, line)
					break
				}
			}
			if len(sortedLines) < noOfLines {
				sortedLines = append(sortedLines, line)
			}
		}
	}

	taxBrackets, err := makeSortedTaxBrackets(sortedLines)
	if err != nil {
		return nil, err
	}

	return taxBrackets, nil
}
