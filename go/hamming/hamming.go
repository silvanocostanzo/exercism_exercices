package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("err")
	}

	v := strings.Compare(a, b)

	if v == 0 {
		return v, nil
	}

	counter := 0
	for i, c := range a {
		if string(c) != string(b[i]) {
			counter++
		}
	}
	return counter, nil
}
