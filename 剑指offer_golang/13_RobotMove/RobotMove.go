//  @author: cord
//  @create: 2019-05-21 21:17
// 面试题13：机器人的运动范围
// 题目：地上有一个m行n列的方格。一个机器人从坐标(0, 0)的格子开始移动，它
// 每一次可以向左、右、上、下移动一格，但不能进入行坐标和列坐标的数位之和
// 大于k的格子。例如，当k为18时，机器人能够进入方格(35, 37)，因为3+5+3+7=18。
// 但它不能进入方格(35, 38)，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

package main

import "fmt"

func movingCount(threshold int, rows int, cols int) int {
	if threshold < 0 || rows <= 0 || cols <= 0 {
		return 0
	}
	visited := make([]bool, rows*cols)
	for i := 0; i < rows*cols; i++ {
		visited[i] = false
	}
	count := movingCountCore(threshold, rows, cols, 0, 0, visited)
	return count
}

func movingCountCore(threshold int, rows int, cols int, row int, col int, visited []bool) int {
	count := 0
	if check(threshold, rows, cols, row, col, visited) {
		visited[row*cols+col] = true
		count = 1 + movingCountCore(threshold, rows, cols, row-1, col, visited) +
			movingCountCore(threshold, rows, cols, row+1, col, visited) +
			movingCountCore(threshold, rows, cols, row, col-1, visited) +
			movingCountCore(threshold, rows, cols, row, col+1, visited)
	}
	return count
}

func check(threshold int, rows int, cols int, row int, col int, visited []bool) bool {
	if row >= 0 && row < rows && col >= 0 && col < cols && !visited[row*cols+col] && getDigitSum(row)+getDigitSum(col) <= threshold {
		return true
	}
	return false
}

func getDigitSum(number int) int {
	count := 0
	for number > 0 {
		count += number % 10
		number /= 10
	}
	return count
}

func main() {
	fmt.Println(movingCount(5, 10, 10))  //21
	fmt.Println(movingCount(10, 1, 100)) //29

}
