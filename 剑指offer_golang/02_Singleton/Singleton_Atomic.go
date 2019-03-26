//  @author: cord
//  @create: 2019-03-26 22:20
//  sync/atomic 原子操作

package main

import (
	"fmt"
	"sync/atomic"
)

type singleton struct {
	info string
}

var instance *singleton

var initialized uint32

func GetInstance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	if initialized == 0 {
		instance = &singleton{"hello world"}
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func main() {
	is := GetInstance()
	fmt.Println(is.info)
}
