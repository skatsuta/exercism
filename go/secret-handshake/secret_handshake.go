package secret

var (
	ints       = []int{1, 2, 4, 8}
	handshakes = map[int]string{
		1: "wink",
		2: "double blink",
		4: "close your eyes",
		8: "jump",
	}
)

// Handshake returns events.
func Handshake(n int) []string {
	if n < 0 {
		return nil
	}

	reverse := false
	if n&(1<<4) == 1<<4 {
		reverse = true
	}

	l := uint64(len(handshakes))
	h := make([]string, 0, l)
	for i := range ints {
		e := ints[i]
		if reverse {
			e = ints[int(l)-1-i]
		}
		if v, exists := handshakes[n&e]; exists {
			h = append(h, v)
		}
	}
	return h
}
