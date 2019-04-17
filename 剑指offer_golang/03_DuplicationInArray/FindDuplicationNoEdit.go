//  @author: cord
//  @create: 2019-04-02 23:20
// 面试题3（二）：不修改数组找出重复的数字
// 题目：在一个长度为n+1的数组里的所有数字都在1到n的范围内，所以数组中至
// 少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入的
// 数组。例如，如果输入长度为8的数组{2, 3, 5, 4, 3, 2, 6, 7}，那么对应的
// 输出是重复的数字2或者3。

//解法的思路是根据值的范围，采用二分法寻找重复的值在分界的哪一边，然后逐步细化定位
//该方法可行的前提是数组内的值均在length-1的范围内

package main

import "fmt"

func getDuplication(numbers []int) int {
	length := len(numbers)
	if numbers == nil || length <= 0 {
		return -1
	}
	start := 1
	end := length - 1
	for {
		if end < start {
			break
		}
		middle := ((end - start) >> 1) + start
		count := countRange(numbers, start, middle)
		if end == start {
			if count > 1 {
				return start
			} else {
				break
			}
		}
		if count > (middle - start + 1) {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return -1
}

func countRange(numbers []int, start int, end int) int {
	if numbers == nil {
		return 0
	}
	count := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] >= start && numbers[i] <= end {
			count++
		}
	}
	return count
}

func main() {
	numbers := []int{2, 3, 5, 4, 3, 2, 6, 7}
	//numbers := []int{2, 3, 5, 9, 6, 1, 0, 9}  //重复数字的值超过length-1
	ret := getDuplication(numbers)
	fmt.Println(ret)
}
