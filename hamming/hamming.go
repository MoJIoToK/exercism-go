package hamming

import "fmt"

// Distance calculates The Hamming Distance between two strands of DNA.
// Two strands are represented by string (a, b string).
// If two strands are not the same length, Distance return error.
func Distance(a, b string) (int, error) {

	//Differences is a count of mistakes between two strands of DNA equal length.
	var differences int = 0
	if len(a) != len(b) {
		return 0, fmt.Errorf("hamming distance is only defined for " +
			"sequences of equal length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			differences++
		}
	}
	return differences, nil
}
