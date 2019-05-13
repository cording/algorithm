//  @author: cord
//  @create: 2019-05-13 21:36
// 面试题11：旋转数组的最小数字
// 题目：把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如数组
// {3, 4, 5, 1, 2}为{1, 2, 3, 4, 5}的一个旋转，该数组的最小值为1。

package main

import (
	"fmt"
)

func main() {
	array1 := []int{3, 4, 5, 1, 2}
	fmt.Printf("array1 min %d\n", Min(array1))
	array2 := []int{3, 4, 5, 1, 1, 2}
	fmt.Printf("array2 min %d\n", Min(array2))
	array3 := []int{3, 4, 5, 1, 2, 2}
	fmt.Printf("array3 min %d\n", Min(array3))
	array4 := []int{1, 0, 1, 1, 1}
	fmt.Printf("array4 min %d\n", Min(array4))
	array5 := []int{1, 2, 3, 4, 5}
	fmt.Printf("array5 min %d\n", Min(array5))
	array6 := []int{3}
	fmt.Printf("array6 min %d\n", Min(array6))
	var array7 []int = nil
	fmt.Printf("array7 min %d\n", Min(array7))
}

func Min(numbers []int) int {
	if numbers == nil {
		panic("invalid parameters")
	}
	index1, index2 := 0, len(numbers)-1
	indexMid := index1
	for numbers[index1] >= numbers[index2] {
		// 如果index1和index2指向相邻的两个数，
		// 则index1指向第一个递增子数组的最后一个数字，
		// index2指向第二个子数组的第一个数字，也就是数组中的最小数字
		if index2-index1 == 1 {
			indexMid = index2
			break
		}
		// 如果下标为index1、index2和indexMid指向的三个数字相等，
		// 则只能顺序查找
		indexMid = (index1 + index2) / 2
		if numbers[index1] == numbers[index2] && numbers[index2] == numbers[indexMid] {
			return MinInOrder(numbers, index1, index2)
		}
		//缩小查找范围
		//如果index1到indexMid仍然增序则从indexMid开始
		if numbers[indexMid] >= numbers[index1] {
			index1 = indexMid
		} else if numbers[indexMid] <= numbers[index2] {
			index2 = indexMid
		}
	}
	return numbers[indexMid]
}

//顺序查找最小值
func MinInOrder(numbers []int, index1 int, index2 int) int {
	result := numbers[index1]
	for i := index1 + 1; i <= index2; i++ {
		if numbers[i] < result {
			result = numbers[i]
		}
	}
	return result
}
