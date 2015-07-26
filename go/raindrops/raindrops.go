package raindrops

import "strconv"

var (
	primes = [...]int{3, 5, 7}
	words  = [...][]byte{[]byte("Pling"), []byte("Plang"), []byte("Plong")}
	maxLen = 15
)

// Convert returns a string, the contents of which depends on the number's prime factors.
//
// - If the number contains 3 as a prime factor, output 'Pling'.
// - If the number contains 5 as a prime factor, output 'Plang'.
// - If the number contains 7 as a prime factor, output 'Plong'.
// - If the number does not contain 3, 5, or 7 as a prime factor,
//   just pass the number's digits straight through.
func Convert(n int) string {
	buf := make([]byte, 0, maxLen)
	for i, prime := range primes {
		if n%prime == 0 {
			buf = append(buf, words[i]...)
		}
	}

	if len(buf) == 0 {
		return strconv.Itoa(n)
	}
	return string(buf)
}
