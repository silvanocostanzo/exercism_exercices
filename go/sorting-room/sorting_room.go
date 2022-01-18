package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	n := nb.Number()
	return fmt.Sprintf("This is a box containing the number %.1f", float64(n))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {

	s, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}

	intVal, err := strconv.Atoi(s.Value())
	if err != nil {
		return 0
	}
	return intVal
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	nb := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(nb))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case int, float64:
		s, ok := v.(float64)
		if ok {
			return DescribeNumber(s)
		}
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	}
	return "Return to sender"
}
