//  @author: cord
//  @create: 2019-05-22 21:55
// 面试题14：剪绳子
// 题目：给你一根长度为n绳子，请把绳子剪成m段（m、n都是整数，n>1并且m≥1）。
// 每段的绳子的长度记为k[0]、k[1]、……、k[m]。k[0]*k[1]*…*k[m]可能的最大乘
// 积是多少？例如当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此
// 时得到最大的乘积18。

package main

import (
	"fmt"
	"math"
)

//动态规划
//f(n) = max(f(i)*f(n-i)) 递归公式，递归有很多重复子问题，可以考虑用循环从下往上计算
//从小到大依次求出子问题的最优解
func maxProductAfterCuttingSolution1(length int) int {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}
	//product数组里面存的是子问题的最优解，考虑的是部分绳子的情况，与上面的不一样
	//长度小于等于3最好肯定是一刀不剪了
	//product := [length+1]int{}
	product := make([]int, length+1)
	product[0] = 0
	product[1] = 1
	product[2] = 2
	product[3] = 3

	var max int
	//从小往大计算分段的最大值
	for i := 4; i <= length; i++ {
		max = 0
		for j := 1; j <= i/2; j++ { //这里取中位数，因为由于j,i-j, 后面一半是重复的
			if product[j]*product[i-j] > max {
				max = product[j] * product[i-j]
			}
		}
		product[i] = max
	}
	return product[length]
}

//贪婪算法
//当length >= 5 的时候尽可能剪长度为3的绳子，当剩下绳子为4的时候剪成2+2的绳子
func maxProductAfterCuttingSolution2(length int) int {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}
	timesOf3 := length / 3
	ret := length % 3
	last := 1
	if ret == 1 { //余数为1说明最后一段为4，为4的时候需要剪成2*2，所以要回退一段
		timesOf3 -= 1
		last = 4
	} else if ret == 2 {
		last = 2
	}
	return int(math.Pow(3, float64(timesOf3)) * float64(last))
}

func main() {
	//动态规划
	fmt.Println(maxProductAfterCuttingSolution1(4)) //4
	fmt.Println(maxProductAfterCuttingSolution1(5)) //6
	fmt.Println(maxProductAfterCuttingSolution1(6)) //9
	fmt.Println(maxProductAfterCuttingSolution1(7)) //12
	//贪婪算法
	fmt.Println(maxProductAfterCuttingSolution2(4)) //4
	fmt.Println(maxProductAfterCuttingSolution2(5)) //6
	fmt.Println(maxProductAfterCuttingSolution2(6)) //9
	fmt.Println(maxProductAfterCuttingSolution2(7)) //12
}
