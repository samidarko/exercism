package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

// NumberBox interface
type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

// FancyNumber type
type FancyNumber struct {
	n string
}

// Value method
func (i FancyNumber) Value() string {
	return i.n
}

// FancyNumberBox interface
type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	value, err := strconv.Atoi(fnb.Value())
	if err != nil {
		return 0
	}
	return value
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	if number, ok := i.(int); ok {
		return DescribeNumber(float64(number))
	}
	if number, ok := i.(float64); ok {
		return DescribeNumber(number)
	}
	if number, ok := i.(NumberBox); ok {
		return DescribeNumberBox(number)
	}
	if number, ok := i.(FancyNumberBox); ok {
		return DescribeFancyNumberBox(number)
	}
	return "Return to sender"
}
