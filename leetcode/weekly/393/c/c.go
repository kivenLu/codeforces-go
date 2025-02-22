package main

import (
	"math/bits"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func findKthSmallest(coins []int, k int) int64 {
	subsetLcm := make([]int, 1<<len(coins))
	subsetLcm[0] = 1
	for i, x := range coins {
		bit := 1 << i
		for mask, l := range subsetLcm[:bit] {
			subsetLcm[bit|mask] = lcm(l, x)
		}
	}
	for i := range subsetLcm {
		if bits.OnesCount(uint(i))%2 == 0 {
			subsetLcm[i] *= -1
		}
	}

	ans := sort.Search(slices.Min(coins)*k, func(m int) bool {
		cnt := 0
		for _, l := range subsetLcm[1:] {
			cnt += m / l
		}
		return cnt >= k
	})
	return int64(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
