package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

	"code-cadets-2021/homework_1/task_1/progressive_tax"
)

func main() {
	var inputValue float64
	fmt.Print("Please, enter input value: ")
	fmt.Scanf("%f", &inputValue)

	if inputValue < 0 {
		log.Fatal(errors.New("The input value should be greater than zero."))
	}

	taxBrackets, err := tax.GetTaxBrackets()
	if err != nil {
		log.Fatal(errors.WithMessage(err, "Error while creating tax brackets."))
	}
	fmt.Printf("Za ulaznu vrijednost %.2f iznos poreza je: %.2f\n", inputValue, tax.CalculateTax(inputValue, taxBrackets))
}
