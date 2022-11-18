package leecode

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMulitp(t *testing.T) {
	multiply("123", "456")
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n := len(num2)
	var tmpNum []string

	for i := n - 1; i >= 0; i-- {
		tmpNum = append(tmpNum, mult(int(num2[i]-'0'), num1))
	}
	fmt.Println(tmpNum)

	if len(tmpNum) == 1 {
		return tmpNum[0]
	}

	op := "0"
	r := tmpNum[0]
	for i := 1; i < len(tmpNum); i++ {
		r = add(r, tmpNum[i]+op)
		fmt.Println(r)
		op += "0"
	}

	return r
}

func mult(x int, str string) (ans string) {
	t := 0
	var s string
	for i := len(str) - 1; i >= 0; i-- {
		t += x * int(str[i]-'0')
		s += strconv.Itoa(t % 10)
		t /= 10
	}

	if t != 0 {
		s += strconv.Itoa(t)
	}
	//fmt.Println(s)

	for i := len(s) - 1; i >= 0; i-- {
		ans += string(s[i])
	}

	return
}

func add(str1, str2 string) (ans string) {
	if len(str2) > len(str1) {
		return add(str2, str1)
	}

	var s string
	n := len(str2)
	t := 0
	str1 = reverse(str1)
	str2 = reverse(str2)

	for i := 0; i < len(str1); i++ {
		if i < n {
			t += int(str1[i]-'0') + int(str2[i]-'0')
		} else {
			t += int(str1[i] - '0')
		}
		s += strconv.Itoa(t % 10)
		t /= 10
	}

	if t != 0 {
		s += strconv.Itoa(t)
	}
	for i := len(s) - 1; i >= 0; i-- {
		ans += string(s[i])
	}
	return
}

func reverse(str string) string {
	var result string
	length := len(str)
	for i := 0; i < length; i++ {
		result = result + fmt.Sprintf("%c", str[length-i-1])
	}
	return result
}
func TestMult(t *testing.T) {
	fmt.Println(mult(3, "480"))
}

func TestAdd(t *testing.T) {
	fmt.Println(add("123", "12345"))
}
