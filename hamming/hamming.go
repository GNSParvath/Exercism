package hamming

import "errors"

func Distance(a, b string) (int, error) {
	c := 0
	var err error = nil
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				c++
			}
		}
	} else {
		err = errors.New("unable to calculate distance for inputs of different lengths.")
	}

	return c, err
}
