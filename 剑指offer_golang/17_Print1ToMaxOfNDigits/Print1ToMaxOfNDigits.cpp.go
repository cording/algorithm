//  @author: cord
//  @create: 2019-06-04 21:34
// 面试题17：打印1到最大的n位数
// 题目：输入数字n，按顺序打印出从1最大的n位十进制数。比如输入3，则
// 打印出1、2、3一直到最大的3位数即999。
//思路：这个问题的关键在于没有给出n的范围，如果n超出了能表示的最大范围，这个时候先求最大值然后循环输出的
//的方式就行不通。所以只能通过数组来解决这个问题
//这个问题本质上可以看成是把数字的每一位从0到9的全排列，而全排列可以用递归解决

package main

import "fmt"

func main() {
	//Print1ToMaxOfDigits(1)
	Print1ToMaxOfDigits(3)
	//Print1ToMaxOfDigits(1)
}

func Print1ToMaxOfDigits(n int) {
	if n <= 0 {
		return
	}
	number := make([]int, n)
	for i := 0; i < 10; i++ {
		number[0] = i
		Print1ToMaxOfDigitsRecursively(number, n, 0)
	}
}

func Print1ToMaxOfDigitsRecursively(number []int, length int, index int) {
	if index == length-1 {
		PrintNumber(number)
		return
	}
	for i := 0; i < 10; i++ {
		number[index+1] = i
		Print1ToMaxOfDigitsRecursively(number, length, index+1)
	}
}

func PrintNumber(number []int) {
	isBeginning := true
	nLength := len(number)

	for i := 0; i < nLength; i++ {
		//找到第一个不是0的
		if isBeginning && number[i] != 0 {
			isBeginning = false
		}
		if !isBeginning {
			fmt.Printf("%d", number[i])
		}
	}
	fmt.Println()
}
