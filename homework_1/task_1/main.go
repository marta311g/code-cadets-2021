package main

import (
	"flag"
	"log"
	"strings"

	"code-cadets-2021/homework_1/task_1/fizzbuzz"
)

func main() {
	start := flag.Int("start", 1, "fizzbuzz start argument")
	end := flag.Int("end", 20, "fizzbuzz end argument")
	flag.Parse()

	fizzbuzzOutputList, err := fizzbuzz.Fizzbuzz(*start, *end)
	if err != nil {
		log.Fatal(err)
	}

	println(strings.Join(fizzbuzzOutputList, " "))
}
