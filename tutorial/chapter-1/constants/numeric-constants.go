package constants

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func NeedInt(x int) int { return x*10 + 1 }
func NeedFloat(x float64) float64 {
	return x * 0.1
}
