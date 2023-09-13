package inter

// Compute gcd of 2 numbers.
// gcd is always positive.
// It is 0 only if both args are 0.
func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b == 0 {
		return a
	}
	if b < 0 {
		b = -b
	}
	return gcd(b, a%b)
}
