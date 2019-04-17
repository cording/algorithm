//  @author: cord
//  @create: 2019-03-27 22:54
// 题目：在一个长度为n的数组里的所有数字都在0到n-1的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，
// 也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。例如，如果输入长度为7的数组{2, 3, 1, 0, 2, 5, 3}，
// 那么对应的输出是重复的数字2或者3。
// 将对应数字放在相同数组下标的位置,这有一个前提，数组里的元素不比数组的长度大，不然该算法无法处理

package main

import "fmt"

func duplicate(numbers []int, duplication *int) bool {
	length := len(numbers)
	if numbers == nil || length == 0 {
		return false
	}
	for i := 0; i < length; i++ {
		for {
			if numbers[i] == i {
				break
			}
			if numbers[i] == numbers[numbers[i]] {
				*duplication = numbers[i]
				return true
			}
			//交换numbers[i]和numbers[numbers[i]]
			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
		}
	}
	return false
}

func main() {
	numbers := []int{2, 1, 3, 1, 4}
	//numbers := []int{2, 1, 3, 0, 4}
	var duplication int
	isDuplication := duplicate(numbers, &duplication)

	if isDuplication {
		fmt.Printf("one of the duplicate number is %v\n", numbers[duplication])
	} else {
		fmt.Println("no duplicate number in the array")
	}

}
