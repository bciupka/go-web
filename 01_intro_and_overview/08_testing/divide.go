package main

import "errors"

func Divide(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("0 division error")
	}
	return x / y, nil
}
