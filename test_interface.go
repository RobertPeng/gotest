/*
type i_name interface {
	method1(param_list) return_list
}
*/

package main

import "fmt"

type Igolang interface {
	get()
	post()
}

type STgolang struct {
}

func (this *STgolang) get() {

}

func (this *STgolang) post() {

}

func GolangLTD(i interface{}) {
	_ = i
	fmt.Println(i)
}

// func main() {
// 	var igolang Igolang
// 	igolang = new(STgolang)
// 	_ = igolang
// 	igolang.get()
// 	igolang.post()
// 	GolangLTD(123456)
// 	GolangLTD("asdf")
// }
