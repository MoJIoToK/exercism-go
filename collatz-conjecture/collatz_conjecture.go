package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	var step int

	if n <= 0 {
		return 0, errors.New("negative or zero")
	}

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		step++
	}
	return step, nil
}
