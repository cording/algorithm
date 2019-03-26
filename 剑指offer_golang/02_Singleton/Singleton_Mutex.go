//  @author: cord
//  @create: 2019-03-26 22:20
//  sync.Mutex 同步锁

package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	info string
}

var instance *singleton

var mu sync.Mutex

func GetInstance() *singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()

		if instance == nil {
			instance = &singleton{"hello world"}
		}
	}
	return instance
}

func main() {
	is := GetInstance()
	fmt.Println(is.info)
}
