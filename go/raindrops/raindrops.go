package raindrops

import (
	"bytes"
	"strconv"
)

var (
	primes = [...]int{3, 5, 7}
	words  = [...]string{"Pling", "Plang", "Plong"}
)

// Convert returns a string, the contents of which depends on the number's prime factors.
//
// - If the number contains 3 as a prime factor, output 'Pling'.
// - If the number contains 5 as a prime factor, output 'Plang'.
// - If the number contains 7 as a prime factor, output 'Plong'.
// - If the number does not contain 3, 5, or 7 as a prime factor,
//   just pass the number's digits straight through.
func Convert(n int) string {
	var buf bytes.Buffer
	for i := 0; i < len(primes); i++ {
		if n%primes[i] == 0 {
			_, _ = buf.WriteString(words[i])
		}
	}

	if buf.Len() == 0 {
		return strconv.Itoa(n)
	}
	return buf.String()
}
