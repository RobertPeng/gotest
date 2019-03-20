package main

func multi(n uint64) (ret uint64) {
	if n > 0 {
		ret = n * multi(n-1)
		return ret
	}
	return 1
}

// func main() {
// 	var i uint64 = 15
// 	ret := multi(i)
// 	fmt.Println(ret)
// }
