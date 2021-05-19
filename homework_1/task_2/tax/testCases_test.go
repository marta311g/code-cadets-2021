package tax_test


type calculateTaxTestCase struct {
	inputValue float64
	inputFile string

	expectedOutput float64
	expectingError bool
}


func getCalculateTaxTestCases() []calculateTaxTestCase {
	return []calculateTaxTestCase {
		{
			inputValue: 7000,
			inputFile: "../brackets.txt",

			expectedOutput: 800.0,
			expectingError: false,
		},
		{
			inputValue: 0,
			inputFile: "../brackets.txt",

			expectedOutput: 0.0,
			expectingError: false,
		},
		{
			inputValue: 1000,
			inputFile: "../brackets.txt",

			expectedOutput: 0.0,
			expectingError: false,
		},
		{
			inputValue: 5000,
			inputFile: "../brackets.txt",

			expectedOutput: 400.0,
			expectingError: false,
		},
		{
			inputValue: 369852,
			inputFile: "../brackets.txt",

			expectedOutput: 109355.59999999999,
			expectingError: false,
		},
		{
			inputValue: 75000,
			inputFile: "../brackets2.txt",

			expectedOutput: 12290.0,
			expectingError: false,
		},
		{
			inputValue: 75045.45,
			inputFile: "../brackets2.txt",

			expectedOutput: 12299.999,
			expectingError: false,
		},
	}
}
