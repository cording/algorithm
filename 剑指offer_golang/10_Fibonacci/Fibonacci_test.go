//  @author: cord
//  @create: 2019-05-08 22:36
//测试结果
//BenchmarkFibonacci_Solution2-4   	  300000	    105229 ns/op    循环
//BenchmarkFibonacci_Solution3-4   	  100000	    265674 ns/op    尾递归

package main

import "testing"

func BenchmarkFibonacci_Solution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci_Solution1(i)
	}
}

func BenchmarkFibonacci_Solution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci_Solution2(i)
	}
}

func BenchmarkFibonacci_Solution3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci_Solution3(0, 1, i)
	}
}
