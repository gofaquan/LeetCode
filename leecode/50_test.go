package leecode

import (
	"fmt"
	"testing"
)

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return qM(x, n)
	} else {
		return 1.0 / qM(x, -n)
	}
}

func qM(x float64, n int) float64 {
	ans := 1.0

	xx := x
	for n > 0 {
		if n%2 == 1 {
			ans *= xx
		}

		xx *= xx

		n /= 2
	}

	return ans
}

func TestPow(t *testing.T) {
	fmt.Println(myPow(2, 10))
}
