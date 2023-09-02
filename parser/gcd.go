package parser

// Compute Gcd of 2 numbers.
// Gcd is always positive.
// It is 0 only if both args are 0.
func Gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b == 0 {
		return a
	}
	if b < 0 {
		b = -b
	}
	return Gcd(b, a%b)
}
