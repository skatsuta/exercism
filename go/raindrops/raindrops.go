package raindrops

import (
	"math"
	"strconv"
)

var words = []byte("PlingPlangPlong")

// Convert returns a string, the contents of which depends on the number's prime factors.
//
// - If the number contains 3 as a prime factor, output 'Pling'.
// - If the number contains 5 as a prime factor, output 'Plang'.
// - If the number contains 7 as a prime factor, output 'Plong'.
// - If the number does not contain 3, 5, or 7 as a prime factor,
//   just pass the number's digits straight through.
func Convert(n int) string {
	primes := primeFactor(n)

	b := make([]byte, 0, len(words))
	for _, p := range primes {
		switch {
		case p%3 == 0:
			b = append(b, words[0:5]...)
		case p%5 == 0:
			b = append(b, words[5:10]...)
		case p%7 == 0:
			b = append(b, words[10:15]...)
		}
	}
	if len(b) == 0 {
		return strconv.Itoa(n)
	}
	return string(b)
}

func primeFactor(n int) []int {
	ch := make(chan int)
	m := int(math.Sqrt(float64(n)))
	go generate(ch)
	factors := make([]int, 0, m)
	for i := 0; i <= m+1; i++ {
		prime := <-ch
		if n%prime == 0 {
			factors = append(factors, prime)
		}
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	return factors
}

// generate generates a integer sequence into channel.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
