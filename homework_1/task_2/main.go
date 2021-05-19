package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

	"code-cadets-2021/homework_1/task_1/tax"
)

func main() {
	var inputValue float64
	fmt.Print("Please, enter input value: ")
	fmt.Scanf("%f", &inputValue)

	if inputValue < 0 {
		log.Fatal(errors.New("the input value should be greater than zero"))
	}

	//the brackets are defined in file "brackets.txt"
	valueOfTax, err := tax.CalculateTax(inputValue, "brackets.txt")
	if err != nil {
		log.Fatal(errors.WithMessage(err, "error while calculating tax"))
	}

	fmt.Printf("Za ulaznu vrijednost %.2f iznos poreza je: %.2f\n", inputValue, valueOfTax)
}
