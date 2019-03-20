package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
	Age  int
}

// func main() {
// 	x := 123
// 	v := reflect.ValueOf(&x)
// 	fmt.Println(v.Elem())
// 	v.Elem().SetInt(999)
// 	fmt.Println(x)
// }

// func Set(o interface{}) {
// 	v := reflect.ValueOf(o)
// 	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
// 		fmt.Println("can not set")
// 		return
// 	}

// 	v = v.Elem()

// 	f := v.FieldByName("Name")
// 	if !f.IsValid() {
// 		fmt.Println("field invalid")
// 	}
// 	if f.Kind() == reflect.String {
// 		f.SetString("Peng")
// 	}
// }

// func main() {
// 	u := User{1, "Tom", 12}
// 	Set(&u)
// 	fmt.Println(u)
// }
func (u User) Hello(name string) {
	fmt.Println("Hello", name, "My name is", u.Name)
}

// func main() {
// 	u := User{1, "Peng", 30}
// 	v := reflect.ValueOf(u)
// 	mv := v.MethodByName("Hello")
// 	args := []reflect.Value{reflect.ValueOf("JOE")}
// 	mv.Call(args)
// }
