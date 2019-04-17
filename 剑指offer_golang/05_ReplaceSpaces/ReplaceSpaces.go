//  @author: cord
//  @create: 2019-04-11 21:48
// 面试题5：替换空格
// 题目：请实现一个函数，把字符串中的每个空格替换成"%20"。例如输入“We are happy.”，
// 则输出“We%20are%20happy.”。
// 如果在原字符串上进行替换，有可能会覆盖修改在该字符串后面的内存
//解决思路: 先遍历字符串，统计出空格总数，并计算出替换后的字符串长度，分配对应数组
//然后将字符串中的数据复制到数组中，当遇到空格的时候复制新字符串
//此实现为参考strings.Replace实现
package main

import (
	"fmt"
	. "strings"
)

func main() {
	fmt.Println(ReplaceBlank("a b c d e"))
	fmt.Println(ReplaceBlank("abcdef"))
	fmt.Println(ReplaceBlank(""))
}

func ReplaceBlank(s string) string {
	//len()方法应用在string上返回的是其对应的字节数组的长度，如果要字符串本身的长度得len([]rune(s))
	if len(s) == 0 {
		return s
	}
	//统计空格数量
	m := Count(s, " ")
	if m == 0 {
		return s
	}
	//根据空格数量以及替换字符串长度计算替换后长度并分配对应的数组
	t := make([]byte, len(s)+m*(len("%20")-len(" ")))

	start := 0 //字符串s的下标
	w := 0     //字节数组t的下标
	for i := 0; i < m; i++ {
		j := start
		j += Index(s[start:], " ")   //找到空格的位置
		w += copy(t[w:], s[start:j]) //将空格前面的部分复制到数组
		w += copy(t[w:], "%20")      //将替换字符串复制到数组
		start = j + len(" ")         //字符串下标增加空格长度
	}
	return string(t[0:w])
}
