package tax_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	tax "code-cadets-2021/homework_1/task_1/progressive_tax"
)

func TestCalculateTax(t *testing.T) {
	for index, testCase := range getCalculateTaxTestCases() {
		Convey(fmt.Sprintf("Test case %d: %v", index, testCase), t, func() {

			actualOutput, actualErr := tax.CalculateTax(testCase.inputValue, testCase.inputTextBrackets)

			if testCase.expectingError {
				So(actualErr, ShouldBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, testCase.expectedOutput)
			}
		})
	}
}

func TestGetTaxBrackets(t *testing.T) {
	for index, testCase := range getGetTaxBracketsTestCases() {
		Convey(fmt.Sprintf("Test case %d: %v", index, testCase), t, func() {

			actualOutput, actualErr := tax.GetTaxBrackets(testCase.inputFile)

			if testCase.expectingError {
				So(actualErr, ShouldBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, testCase.expectedOutput)
			}
		})
	}
}
