// package main

// import "fmt"

// func main() {
// 	c := make(chan bool)
// 	go func() {
// 		fmt.Println("GOOD---")
// 		c <- true
// 		close(c)
// 	}()
// 	// <-c
// 	close(c)
// 	for v := range c {
// 		fmt.Println(v)
// 	}
// }

package main

import "fmt"

// func main() {
// 	origin, wait := make(chan int), make(chan struct{})
// 	Processor(origin, wait)
// 	for num := 2; num < 10; num++ {
// 		origin <- num
// 	}
// 	close(origin)
// 	<-wait
// }

func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			fmt.Println("haha", num)
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}
