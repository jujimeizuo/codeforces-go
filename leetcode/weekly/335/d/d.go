package main

// https://space.bilibili.com/206214
func waysToReachTarget(target int, types [][]int) int {
	const mod int = 1e9 + 7
	f := make([]int, target+1)
	f[0] = 1
	s := 0
	for _, p := range types {
		count, marks := p[0], p[1]
		s = min(s+count*marks, target)
		for j := s; j > 0; j-- {
			for k := 1; k <= count && k <= j/marks; k++ {
				f[j] += f[j-k*marks]
			}
			f[j] %= mod
		}
	}
	return f[target]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
