package leecode

import (
	"fmt"
	"math"
	"testing"
)

// 暴力不行
func countPrimes1(n int) (cnt int) {
	if n < 2 {
		return 0
	}

	for i := 2; i < n; i++ {
		sqrti := int(math.Sqrt(float64(i)))
		isPrim := true
		for j := 2; j <= sqrti; j++ {
			if i%j == 0 {
				isPrim = false
				break
			}
		}

		if isPrim {
			fmt.Println(i)
			cnt++
		}

	}

	return
}

func TestC(t *testing.T) {
	countPrimes1(10)
}
