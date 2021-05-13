package fizzbuzz_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"code-cadets-2021/homework_1/task_1/fizzbuzz"
)

func TestFizzbuzz(t *testing.T) {
	for index, testCase := range getTestCases() {
		Convey(fmt.Sprintf("Test case %d: %v", index, testCase), t, func() {
			actualOutput, actualErr := fizzbuzz.Fizzbuzz(testCase.inputStart, testCase.inputEnd)

			if testCase.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, testCase.expectedOutput)
			}
		})
	}
}
