package main

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	defer close(prime_nums)
	for i := 2; i <= N; i++ {
		select {
		case <-stop:
			return
		default:
			if IsPrime(i) {
				prime_nums <- i
			}
		}
	}
}

func IsPrime(x int) bool {
	if x <= 1 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
