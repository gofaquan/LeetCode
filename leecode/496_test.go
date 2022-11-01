package leecode

import (
	"fmt"
	"testing"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var result []int

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {

			flag := true
			if nums1[i] == nums2[j] { // 找到了
				// 从 j+1 开始寻找
				for k := j + 1; k < len(nums2); k++ {
					if nums2[k] > nums2[j] {
						flag = false
						result = append(result, nums2[k])
						break
					}
				}
				if flag {
					result = append(result, -1)
				}

			}
		}
	}

	return result
}

var n1 = []int{4, 1, 2}
var n2 = []int{15, 6, 4, 1, 2, 5, 8, 9}

func Test(t *testing.T) {
	fmt.Println(nextGreaterElement(n1, n2))
}
