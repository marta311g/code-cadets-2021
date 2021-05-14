package tax_test

import "code-cadets-2021/homework_1/task_1/tax"


type calculateTaxTestCase struct {
	inputValue float64
	inputTaxBrackets []tax.TaxBracket

	expectedOutput float64
	expectingError bool
}

type getTaxBracketsTestCase struct {
	inputFile string

	expectedOutput []tax.TaxBracket
	expectingError bool
}


func getCalculateTaxTestCases() []calculateTaxTestCase {
	return []calculateTaxTestCase {
		{
			inputValue: 7000,
			inputTaxBrackets: []tax.TaxBracket{
				tax.TaxBracket{0, 1000, 0},
				tax.TaxBracket{1000, 5000, 0.1},
				tax.TaxBracket{5000, 10000, 0.2},
				tax.TaxBracket{10000, -1, 0.3},
			},

			expectedOutput: 800.0,
			expectingError: false,
		},
		{
			inputValue: 0,
			inputTaxBrackets: []tax.TaxBracket{
				tax.TaxBracket{0, 1000, 0},
				tax.TaxBracket{1000, 5000, 0.1},
				tax.TaxBracket{5000, 10000, 0.2},
				tax.TaxBracket{10000, -1, 0.3},
			},

			expectedOutput: 0.0,
			expectingError: false,
		},
		{
			inputValue: 1000,
			inputTaxBrackets: []tax.TaxBracket{
				tax.TaxBracket{0, 1000, 0},
				tax.TaxBracket{1000, 5000, 0.1},
				tax.TaxBracket{5000, 10000, 0.2},
				tax.TaxBracket{10000, -1, 0.3},
			},

			expectedOutput: 0.0,
			expectingError: false,
		},
		{
			inputValue: 5000,
			inputTaxBrackets: []tax.TaxBracket{
				tax.TaxBracket{0, 1000, 0},
				tax.TaxBracket{1000, 5000, 0.1},
				tax.TaxBracket{5000, 10000, 0.2},
				tax.TaxBracket{10000, -1, 0.3},
			},

			expectedOutput: 400.0,
			expectingError: false,
		},
		{
			inputValue: 369852,
			inputTaxBrackets: []tax.TaxBracket{
				tax.TaxBracket{0, 1000, 0},
				tax.TaxBracket{1000, 5000, 0.1},
				tax.TaxBracket{5000, 10000, 0.2},
				tax.TaxBracket{10000, -1, 0.3},
			},

			expectedOutput: 109355.59999999999,
			expectingError: false,
		},
		{
			inputValue: 75000,
			inputTaxBrackets: []tax.TaxBracket{
				tax.TaxBracket{0,9875,0.1},
				tax.TaxBracket{9875,40125,0.12},
				tax.TaxBracket{40125,85525,0.22},
				tax.TaxBracket{85525,163300,0.24},
				tax.TaxBracket{163300,207350,0.32},
				tax.TaxBracket{207350,518400,0.35},
				tax.TaxBracket{518400,-1,0.37},
			},

			expectedOutput: 12290.0,
			expectingError: false,
		},
		{
			inputValue: 75045.45,
			inputTaxBrackets: []tax.TaxBracket{
				tax.TaxBracket{0,9875,0.1},
				tax.TaxBracket{9875,40125,0.12},
				tax.TaxBracket{40125,85525,0.22},
				tax.TaxBracket{85525,163300,0.24},
				tax.TaxBracket{163300,207350,0.32},
				tax.TaxBracket{207350,518400,0.35},
				tax.TaxBracket{518400,-1,0.37},
			},

			expectedOutput: 12299.999,
			expectingError: false,
		},
	}
}

func getGetTaxBracketsTestCases() []getTaxBracketsTestCase {
	return []getTaxBracketsTestCase {
		{
			inputFile: "../brackets.txt",

			expectedOutput: []tax.TaxBracket{
				tax.TaxBracket{0, 1000, 0},
				tax.TaxBracket{1000, 5000, 0.1},
				tax.TaxBracket{5000, 10000, 0.2},
				tax.TaxBracket{10000, -1, 0.3},
			},
			expectingError: false,
		},
	}
}
