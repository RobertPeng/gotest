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
