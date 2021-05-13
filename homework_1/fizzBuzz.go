package main

import (
	"flag"
	"strconv"
	"strings"
)

func main() {
	start := flag.Int("start", 0, "fizzbuzz start argument")
	end := flag.Int("end", 20, "fizzbuzz end argument")
	flag.Parse()
	fizzbuzzOutputList := []string{}

	// version 1: with array - saving elements then printing
	for i := *start; i <= *end; i++ {
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
	println(strings.Join(fizzbuzzOutputList, " "))

	// version 2: without array - printing elements
	for i := *start; i <= *end; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			print("FizzBuzz ")
		} else if  i % 3 == 0 {
			print("Fizz ")
		} else if i % 5 == 0 {
			print("Buzz ")
		} else {
			print(i)
			print(" ")
		}
	}
	print("\n")
}
