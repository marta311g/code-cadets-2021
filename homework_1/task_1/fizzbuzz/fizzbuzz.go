package fizzbuzz

import (
	"strconv"

	"github.com/pkg/errors"
)

func Fizzbuzz(start, end int) ([]string, error){
	if end < start {
		return nil, errors.New("The end flag should be greater that the start flag.")
	}
	if start < 1 {
		return nil, errors.New("The start flag should be greater that zero.")
	}

	fizzbuzzOutputList := []string{}

	for i := start; i <= end; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fizzbuzzOutputList = append(fizzbuzzOutputList, "FizzBuzz")
		} else if  i % 3 == 0 {
			fizzbuzzOutputList = append(fizzbuzzOutputList, "Fizz")
		} else if i % 5 == 0 {
			fizzbuzzOutputList = append(fizzbuzzOutputList, "Buzz")
		} else {
			fizzbuzzOutputList = append(fizzbuzzOutputList, strconv.Itoa(i))
		}
	}

	return fizzbuzzOutputList, nil
}
