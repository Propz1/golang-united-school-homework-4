package string_sum

import (
	"fmt"
	"testing"
)

func TestSortMapValues(t *testing.T) {

	t.Run("Positive", func(t *testing.T) {

		var input string = "\t\t\t\t\n\n\n\n -               98961230121985+  1   \t\t\t\n "

		//input := "\t\t\t\t\n\n\n\n -               98961230121985-  3 +! +  5 -2-1   -1 -6+25-9+8-9-188-1+1 \t\t\t\n "
		//var input string = "\t\t\t\t\n\n\n\n -               98961230121985-  3 +6  \t\t\t\n "
		//input := "\t\t\t\t\n\n\n\n        \t\t\t\n "
		//input := "\t\t\t\t\n\n\n\n -               98961230121985 +8 + \t\t\t\n "
		//input := "\t\t\t\t\n\n\n\n -               98961230121985 8\t\t\t\n "
		//input := "\t\t\t\t\n\n\n\n -               98961230121985  \t\t\t\n "

		var result = "-98961230121984"

		realResult, err := StringSum(input)

		if err != nil {
			fmt.Println(err)
		}

		if realResult != result {
			t.Errorf("Expected result %v != %v real result", result, realResult)
		}

	})

	t.Run("Negative_1", func(t *testing.T) {

		var (
			input  string = "\t\t\t\t\n\n\n\n        \t\t\t\n "
			result        = ""
		)

		realResult, err := StringSum(input)

		if err != nil {

			fmt.Println(err.Error())

		}

		if realResult != result {
			t.Errorf("Expected result %s != %s real result", result, realResult)
		}

	})

	t.Run("Negative_2", func(t *testing.T) {

		var (
			input  string = "\t\t\t\t\n\n\n\n -               98961230121985-  3 +6  \t\t\t\n "
			result        = ""
		)

		realResult, err := StringSum(input)
		if err != nil {
			fmt.Printf("%s\n", err)

		}

		if realResult != result {
			t.Errorf("Expected result %s != %s real result", result, realResult)
		}

	})

}
