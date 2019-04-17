//  @author: cord
//  @create: 2019-04-09 21:16
// 面试题4：二维数组中的查找
// 题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按
// 照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个
// 整数，判断数组中是否含有该整数。
// 思路:
//可以选取右上角和左下角的数字，比如选取右上角，如果比它大，说明排除当前行，如果比它小排除当前列，如果相等即找到该数
//左下角类似

package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
	fmt.Printf("find 7 in matrix{%v}: %t\n", matrix, Find(matrix, 7))
	fmt.Printf("find 19 in matrix{%v}: %t", matrix, Find(matrix, 19))
}

//matrix矩阵 number 要查找的数字
func Find(matrix [][]int, number int) bool {
	var found = false
	rows, columns := len(matrix), len(matrix[0])
	if matrix == nil || rows == 0 || columns == 0 {
		return found
	}
	row, column := 0, columns-1
	for row < rows && column >= 0 {
		if matrix[row][column] == number {
			found = true
			break
		} else if matrix[row][column] > number {
			column--
		} else {
			row++
		}
	}
	return found
}
