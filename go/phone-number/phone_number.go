package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

// Number return phone number
func Number(input string) (string, error) {
	reg := regexp.MustCompile(`[^\d]`)
	input = reg.ReplaceAllString(input, "")

	if len(input) > 0 && input[0] == '1' {
		input = input[1:] // removes '1'
	}

	if len(input) != 10 {
		return "", errors.New("invalid length")
	}

	if int(input[0]-'0') < 2 {
		return "", errors.New("invalid area code")
	}

	if int(input[3]-'0') < 2 {
		return "", errors.New("invalid number")
	}

	return input, nil
}

// AreaCode return area code
func AreaCode(input string) (string, error) {
	input, err := Number(input)

	if err != nil {
		return "", err
	}

	return input[:3], nil
}

// Format return a formatted phone number with its area code
func Format(input string) (string, error) {
	input, err := Number(input)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", input[:3], input[3:6], input[6:]), nil
}
