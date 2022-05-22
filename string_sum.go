package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	plus    = "+"
	minus   = "-"
	blank   = " "
	numbers = "0123456789"
	n       = 2 //Number of numbers
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	sign                bool
	value               bool
	contains            bool
	sum                 int
	gap                 bool
	st                  string

	str = make([]string, 0, 100)

	builder strings.Builder
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {

	str = nil
	sum = 0

	intput1 := strings.TrimSpace(input)

	//is blank
	if intput1 == "" {
		err := fmt.Errorf("the string is empty: \"%w\"", errorEmptyInput)
		return "", err
	}

	for _, r := range intput1 {

		contains = strings.ContainsRune(numbers, r)

		symbol := string(r)

		if symbol == plus || symbol == minus {

			if sign {
				if len(str) > 0 {
					prev_sign := str[len(str)-1]
					err := fmt.Errorf("two signs %s successively", "\""+prev_sign+symbol+"\"")
					return "", err
				}
			}

			if value {
				if len(str) > 0 {
					str = append(str, ";", symbol)
				} else {
					str = append(str, symbol)
				}
			} else {
				str = append(str, symbol)
			}

			sign = true
			value = false
			gap = false
			continue
		}

		if contains {

			if len(str) > 0 {
				pre_simbol := str[len(str)-1]

				if pre_simbol != plus && pre_simbol != minus && gap {
					if n == 2 {
						err := fmt.Errorf("two number successively without sign: \"%w\"", errorNotTwoOperands)
						return "", err
					} else {
						err := fmt.Errorf("two number successively without sign")
						return "", err
					}
				}
			}

			str = append(str, string(r))

			value = true
			sign = false
			gap = false

			continue
		}

		if symbol == blank {
			gap = true
		}

		if symbol != plus && symbol != minus && symbol != blank && !contains {
			err := fmt.Errorf("\"%s\" is invalid character", string(r))
			return "", err
		}

	}

	builder.Reset()

	for _, w := range str {
		builder.WriteString(w)
	}

	st = builder.String()

	stsplit := strings.Split(st, ";")

	c := 0

	for _, p := range stsplit {

		i, err := strconv.Atoi(p)
		if err != nil {
			if n == 2 {
				err := fmt.Errorf("\"%s\" symbol can't be converted to Int: \"%w\"", p, errorNotTwoOperands)
				return "", err
			} else {
				err := fmt.Errorf("\"%s\" symbol can't be converted to Int", p)
				return "", err
			}
		}

		c++
		sum += i
	}

	if c == 1 {
		if n == 2 {
			err := fmt.Errorf("the string contains only one numbers \"%s\": \"%w\"", strconv.Itoa(sum), errorNotTwoOperands)
			return "", err
		}
		err := fmt.Errorf("the string contains only one numbers \"%s\"", strconv.Itoa(sum))
		return "", err
	}

	if c > n {
		if n == 2 {
			err := fmt.Errorf("the string contains more than %s numbers: %w", strconv.Itoa(n), errorNotTwoOperands)
			return "", err
		}
		err := fmt.Errorf("the string contains more than %s numbers", strconv.Itoa(n))
		return "", err
	}

	output = strconv.Itoa(sum)

	return output, nil
}
