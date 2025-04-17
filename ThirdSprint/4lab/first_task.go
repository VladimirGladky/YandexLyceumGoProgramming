package main

func StringsGen(lines ...string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, line := range lines {
			out <- line
		}
	}()
	return out
}
