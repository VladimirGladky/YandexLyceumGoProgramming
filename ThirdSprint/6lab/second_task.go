package main

func AtomicSwap(a *int32, b *int32) {
	t := *a
	*a = *b
	*b = t
}
