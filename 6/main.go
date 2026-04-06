package main

func RandomGenerator() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}