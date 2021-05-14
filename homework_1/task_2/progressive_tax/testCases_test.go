package tax_test

import "code-cadets-2021/homework_1/task_1/progressive_tax"


type calculateTaxTestCase struct {
	inputValue float64
	inputTextBrackets []tax.TaxBracket

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
			inputTextBrackets: []tax.TaxBracket{
				tax.TaxBracket{0, 1000, 0},
				tax.TaxBracket{1000, 5000, 0.1},
				tax.TaxBracket{5000, 10000, 0.2},
				tax.TaxBracket{10000, -1, 0.3},
			},

			expectedOutput: 800.0,
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