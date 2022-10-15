package main

import (
	"fmt"
	"math"
)

func max(vs ...int) int {
	res := math.MinInt
	for _, v := range vs {
		if v > res {
			res = v
		}
	}
	return res
}

func min(vs ...int) int {
	res := math.MaxInt
	for _, v := range vs {
		if v < res {
			res = v
		}
	}
	return res
}

func getScore(first, second string, match, mismatch, gap int) int {
	n := len(first)
	m := len(second)

	tmp := make([]int, m+1)

	for i := 1; i < m+1; i++ {
		tmp[i] = gap * i
	}

	for i := 1; i < n+1; i++ {
		prev := i * gap
		newTmp := make([]int, m+1)
		tmp[0] = gap * (i - 1)
		newTmp[0] = gap * i

		for j := 1; j < m+1; j++ {

			cost := 0
			if first[i-1] == second[j-1] {
				cost = match
			} else {
				cost = mismatch
			}

			substitution := tmp[j-1] + cost
			deleting := tmp[j] + gap
			inserting := prev + gap

			maxCost := max(substitution, deleting, inserting)
			prev = maxCost
			newTmp[j] = maxCost
		}
		copy(tmp, newTmp)
	}

	return tmp[m]
}

func optimizeSolve(first, second string, match, mismatch, gap, k int) (int, error) {
	a := solve(first, second, match, mismatch, gap, k)
	b := solve(first, second, match, mismatch, gap, k+1)
	if a < b {
		return -1, fmt.Errorf("bad input")
	}
	return a, nil
}

func solve(first, second string, match, mismatch, gap, k int) int {
	type Index struct {
		X, Y int
	}
	l1 := len(first) + 1
	l2 := len(second) + 1
	d1 := max(l1-l2, 0)
	d2 := max(l2-l1, 0)

	dp := make(map[Index]int)
	dp[Index{0, 0}] = 0

	for i := 1; i < min(k+d1, l1); i++ {
		dp[Index{i, 0}] = dp[Index{i - 1, 0}] + gap
	}
	for i := 1; i < min(k+d2, l2); i++ {
		dp[Index{0, i}] = dp[Index{0, i - 1}] + gap
	}

	for i := 1; i < l1; i++ {
		for j := max(1, i-k-d1); j < min(l2-1, i+k+d2)+1; j++ {

			cost := 0
			if first[i-1] == second[j-1] {
				cost = match
			} else {
				cost = mismatch
			}

			substitution := math.MinInt
			deleting := math.MinInt
			inserting := math.MinInt

			if val, ok := dp[Index{i - 1, j - 1}]; ok {
				substitution = val + cost
			}
			if val, ok := dp[Index{i - 1, j}]; ok {
				deleting = val + gap
			}
			if val, ok := dp[Index{i, j - 1}]; ok {
				deleting = val + gap
			}

			maxCost := max(substitution, deleting, inserting)
			dp[Index{i, j}] = maxCost
		}
	}

	return dp[Index{len(first), len(second)}]
}

func main() {
	// var first, second string
	// var match, mismatch, gap, k int

	// fmt.Scanf("%s %s %d %d %d %d", &first, &second, &match, &mismatch, &gap, &k)

	fmt.Println(solve("AA", "BBBBBAA", 5, -100, -10, 0))
	fmt.Println(solve("AA", "BBBBBAA", 5, -100, -10, 10))
}
