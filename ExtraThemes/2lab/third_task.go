package main

func MinPizzaCost(s, m, l, cs, cm, cl, x int) int {
	const inf = 100000
	minPrices := make([]int, x+1)
	for i := 0; i <= x; i++ {
		minPrices[i] = inf
	}
	minPrices[0] = 0

	for i := 0; i <= x; i++ {
		if minPrices[i] == inf {
			continue
		}
		if i+s <= x {
			minPrices[i+s] = min(minPrices[i+s], minPrices[i]+cs)
		} else {
			minPrices[x] = min(minPrices[x], minPrices[i]+cs)
		}

		if i+m <= x {
			minPrices[i+m] = min(minPrices[i+m], minPrices[i]+cm)
		} else {
			minPrices[x] = min(minPrices[x], minPrices[i]+cm)
		}

		if i+l <= x {
			minPrices[i+l] = min(minPrices[i+l], minPrices[i]+cl)
		} else {
			minPrices[x] = min(minPrices[x], minPrices[i]+cl)
		}
	}

	return minPrices[x]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
