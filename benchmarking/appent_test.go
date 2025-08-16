package main

import "testing"


func AppendSlice(n int) []int {
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, i)
	}
	return nums
}

func BenchmarkAppendSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendSlice(1000)
	}
}