package fizzbuzz_test

type testCase struct {
	inputStart int
	inputEnd int

	expectedOutput []string
	expectingError bool
}

func getTestCases() []testCase {
	return []testCase {
		{
			inputStart: 1,
			inputEnd: 10,

			expectedOutput: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"},
			expectingError:  false,
		},
		{
			inputStart: 1,
			inputEnd: 15,

			expectedOutput: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
			expectingError:  false,
		},
		{
			inputStart: 5,
			inputEnd: 18,

			expectedOutput: []string{"Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz"},
			expectingError:  false,
		},
		{
			inputStart: 1,
			inputEnd: 1,

			expectedOutput: []string{"1"},
			expectingError:  false,
		},
		{
			inputStart: 10,
			inputEnd: 1,

			expectingError:  true,
		},
		{
			inputStart: 0,
			inputEnd: 1,

			expectingError:  true,
		},
		{
			inputStart: 0,
			inputEnd: 0,

			expectingError:  true,
		},
	}
}
