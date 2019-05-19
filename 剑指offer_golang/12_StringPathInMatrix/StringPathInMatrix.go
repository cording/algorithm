//  @author: cord
//  @create: 2019-05-14 21:14
// 面试题12：矩阵中的路径
// 题目：请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有
// 字符的路径。路径可以从矩阵中任意一格开始，每一步可以在矩阵中向左、右、
// 上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入
// 该格子。例如在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字
// 母用下划线标出）。但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个
// 字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
// A B T G
// C F C S
// J D E H
//回溯法，如果叶节点的状态不满足约束条件，那么回溯到上一个节点再尝试其他的选项
//这里使用一维数组来存储矩阵

package main

import "fmt"

func main() {
	fmt.Println(hasPath([]rune("ABTGCFCSJDEH"), 3, 4, "BFCE"))
	fmt.Println(hasPath([]rune("ABTGCFCSJDEH"), 3, 4, "BFCH"))
	fmt.Println(hasPath([]rune("ABTGCFCSJDEH"), 3, 4, "ACFCEH"))
}

func hasPath(matrix []rune, rows int, cols int, str string) bool {
	if matrix == nil || len(str) == 0 || rows < 1 || cols < 1 {
		return false
	}
	visited := make([]bool, rows*cols)
	pathLength := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if hasPathCore(matrix, rows, cols, row, col, str, pathLength, visited) {
				return true
			}
		}
	}
	return false
}

func hasPathCore(matrix []rune, rows int, cols int, row int, col int, str string, pathLength int, visited []bool) bool {
	//如果找到路径最后一个字符了
	if pathLength == len(str) {
		return true
	}
	hasPath := false
	if row >= 0 && row < rows && col >= 0 && col < cols && matrix[row*cols+col] == []rune(str)[pathLength] && !visited[row*cols+col] {
		pathLength++
		visited[row*cols+col] = true
		hasPath = hasPathCore(matrix, rows, cols, row, col-1, str, pathLength, visited) ||
			hasPathCore(matrix, rows, cols, row, col+1, str, pathLength, visited) ||
			hasPathCore(matrix, rows, cols, row-1, col, str, pathLength, visited) ||
			hasPathCore(matrix, rows, cols, row+1, col, str, pathLength, visited)
		if !hasPath {
			pathLength--
			visited[row*cols+col] = false
		}
	}
	return hasPath

}
