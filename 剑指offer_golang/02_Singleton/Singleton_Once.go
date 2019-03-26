//  @author: cord
//  @create: 2019-03-26 22:20
//  sync/Once 进行唯一一次调用

package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	info string
}

var instance *singleton

var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{"hello word"}
	})
	return instance
}

func main() {
	is := GetInstance()
	fmt.Println(is.info)
}
