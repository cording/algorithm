//  @author: cord
//  @create: 2019-05-27 22:10

// 面试题16：数值的整数次方
// 题目：实现函数double Power(double base, int exponent)，求base的exponent
// 次方。不得使用库函数，同时不需要考虑大数问题。

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Power(2, 3))
	fmt.Println(Power(2, -30))
}

func Power(base float64, exponent int) float64 {
	if equal(base, 0.0) && exponent < 0 {
		return 0.0
	}
	if exponent >= 0 {
		return PowerWithUnsignedExponent(base, exponent)
	} else {
		return 1.0 / PowerWithUnsignedExponent(base, -exponent)
	}
}

//func PowerWithUnsignedExponent(base float64, exponent int) float64 {
//	result := 1.0
//	for i:=1;i<=exponent;i++  {
//		result *= base
//	}
//	return result
//}

//递归 x^(2n)=x^n * x^n
func PowerWithUnsignedExponent(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	result := PowerWithUnsignedExponent(base, exponent>>1)
	result *= result
	//如果是奇数则多乘一位
	if exponent&0x1 == 1 {
		result *= base
	}
	return result
}

func equal(a, b float64) bool {
	return math.Abs(a-b) <= 1e-9
}
