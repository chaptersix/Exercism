package hamming

import "errors"

// Distance finds the hamming distance between 2 dna strands formated as a string
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Strands must be of equal length")
	}
	dist := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
