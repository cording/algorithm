// @author: cord
// @create: 2019-03-26 21:59
// 赋值运算符函数:
// 函数结束前返回自身引用，允许连续赋值
// 如果两个对象相同或值相等,不进行操作直接返回

package main

import (
	"fmt"
	"strings"
)

type MyString struct {
	data string
}

func (ms *MyString) assign(m MyString) *MyString {
	if *ms == m || strings.EqualFold(ms.data, m.data) {
		return ms
	} else {
		ms.data = m.data
		return ms
	}
}

func (ms *MyString) print() *MyString {
	fmt.Println(ms.data)
	return ms
}

func main() {
	ms1 := MyString{"a"}
	ms2 := &MyString{"b"}
	ms3 := MyString{"c"}
	ms2.assign(ms1).print().assign(ms3).print()
}

//运行结果:
// a
// c
