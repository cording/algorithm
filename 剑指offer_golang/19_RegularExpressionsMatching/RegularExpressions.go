//  @author: cord
//  @create: 2019-06-11 21:39
// 面试题19：正则表达式匹配
// 题目：请实现一个函数用来匹配包含'.'和'*'的正则表达式。模式中的字符'.'
// 表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题
// 中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"
// 和"ab*ac*a"匹配，但与"aa.a"及"ab*a"均不匹配。

package main

import (
	"fmt"
	"strings"
)

func main() {
	Test("Test01", "", "", true)
	Test("Test02", "", ".*", true)
	Test("Test03", "", ".", false)
	Test("Test04", "", "c*", true)
	Test("Test05", "a", ".*", true)
	Test("Test06", "a", "a.", false)
	Test("Test07", "a", "", false)
	Test("Test08", "a", ".", true)
	Test("Test09", "a", "ab*", true)
	Test("Test10", "a", "ab*a", false)
	Test("Test11", "aa", "aa", true)
	Test("Test12", "aa", "a*", true)
	Test("Test13", "aa", ".*", true)
	Test("Test14", "aa", ".", false)
	Test("Test15", "ab", ".*", true)
	Test("Test16", "ab", ".*", true)
	Test("Test17", "aaa", "aa*", true)
	Test("Test18", "aaa", "aa.a", false)
	Test("Test19", "aaa", "a.a", true)
	Test("Test20", "aaa", ".a", false)
	Test("Test21", "aaa", "a*a", true)
	Test("Test22", "aaa", "ab*a", false)
	Test("Test23", "aaa", "ab*ac*a", true)
	Test("Test24", "aaa", "ab*a*c*a", true)
	Test("Test25", "aaa", ".*", true)
	Test("Test26", "aab", "c*a*b", true)
	Test("Test27", "aaca", "ab*a*c*a", true)
	Test("Test28", "aaba", "ab*a*c*a", false)
	Test("Test29", "bbbba", ".*a*a", true)
	Test("Test30", "bcbbabab", ".*a*a", false)
	Test("Test31", "ab", "a*a", false)
}

func Test(testName string, str string, pattern string, expected bool) {
	if testName != "" {
		fmt.Printf("%s begins:", testName)
	}
	if match(str, pattern) == expected {
		fmt.Printf("Passed.\n")
	} else {
		fmt.Printf("FAILED.\n")
	}

}

func match(str string, pattern string) bool {
	//处理匹配字符串为空时的情况，因为空字符串转换的slice也是空slice,一旦取元素则会越界
	if str == "" {
		if pattern == "" || (len(pattern) == 2 && strings.Index(pattern, "*") == 1) {
			return true
		} else {
			return false
		}
	}
	return matchCore([]rune(str), 0, []rune(pattern), 0)
}

func matchCore(str []rune, sIndex int, pattern []rune, pIndex int) bool {
	//字符串和模式均遍历完,或指向下一位,均越界一位
	if sIndex == len(str) && pIndex == len(pattern) {
		return true
	}
	//字符串未遍历完，但是模式未匹配完
	if sIndex != len(str)-1 && pIndex == len(pattern)-1 {
		return false
	}
	//下标越界
	if sIndex >= len(str) || pIndex >= len(pattern) {
		return false
	}

	//如果下一个字符是*
	if pIndex+1 < len(pattern) && pattern[pIndex+1] == rune('*') {
		if (pattern[pIndex] == rune('.') && sIndex <= len(str)-1) || str[sIndex] == pattern[pIndex] {
			//if (pattern[pIndex] == rune('.') && !checkSliceEnd(str, sIndex)) || str[sIndex] == pattern[pIndex]{
			//	进入有穷自动机下一个状态
			return matchCore(str, sIndex+1, pattern, pIndex+2) ||
				//继续留在自动机的当前位置
				matchCore(str, sIndex+1, pattern, pIndex) ||
				//略过一个'*'
				matchCore(str, sIndex, pattern, pIndex+2)
		} else {
			//略过一个'*'
			return matchCore(str, sIndex, pattern, pIndex+2)
		}
	}

	//如果下下一个字符是*
	if pIndex+2 < len(pattern) && pattern[pIndex+2] == rune('*') {
		if str[sIndex] == pattern[pIndex] {
			return matchCore(str, sIndex+1, pattern, pIndex+3) ||
				matchCore(str, sIndex+1, pattern, pIndex+1)
		}
	}

	if str[sIndex] == pattern[pIndex] || (pattern[pIndex] == rune('.') && sIndex < len(str)) {
		return matchCore(str, sIndex+1, pattern, pIndex+1)
	}

	return false
}

//
//func checkSliceEnd(str []rune, index int)  bool {
//	if (len(str) == 0 && index == 0) || (len(str) > 0 && index == len(str)-1 )  {
//		return true
//	}
//	return false
//}
