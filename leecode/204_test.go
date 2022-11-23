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

// 埃式筛法
func countPrimes2(n int) (cnt int) {
	st := make([]bool, n+1)
	primes := make([]int, n+1)

	for i := 2; i < n; i++ {
		if !st[i] {
			primes[cnt] = n
			cnt++

			for j := 2 * i; j < n; j += i {
				st[j] = true
			}
		}

	}

	return
}

func TestC2(t *testing.T) {
	println(countPrimes2(2))
}

// 线性筛法
func countPrimes(n int) (cnt int) {
	st := make([]bool, n+1)
	primes := make([]int, n+1)

	for i := 2; i < n; i++ {
		if !st[i] {
			primes[cnt] = i
			cnt++
		}

		for j := 0; primes[j] <= n/i; j++ {
			st[primes[j]*i] = true
			if i%primes[j] == 0 {
				break
			}
		}

	}

	return
}

func TestC12(t *testing.T) {
	println(countPrimes(10))
}
