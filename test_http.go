// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"
// )

// func sayHelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	fmt.Println(r.Form)
// 	fmt.Println("path:", r.URL.Path)
// 	fmt.Println("scheme:", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Hello Peng!")
// }

// func main() {
// 	http.HandleFunc("/", sayHelloName)
// 	err := http.ListenAndServe(":9090", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe:", err)
// 	}
// }

// package main

// import (
// 	"io"
// 	"net/http"
// )

// type a struct{}

// func (*a) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	path := r.URL.String()
// 	switch path {
// 	case "/":
// 		io.WriteString(w, "<h1>root</h1><a href=\"abc\">abc</a>")
// 	case "/abc":
// 		io.WriteString(w, "<h1>abc</h1><a href=\"/\">root</a>")
// 	}
// }
// func main() {
// 	http.ListenAndServe(":8080", &a{})
// }

// package main

// import (
// 	"io"
// 	"net/http"
// )

// type b struct{}

// func (*b) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "hello")
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.Handle("/h", &b{})
// 	http.ListenAndServe(":8080", mux)
// }

package main

import (
	"io"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "say hello")
}

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/h", func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "hello")
// 	})
// 	mux.HandleFunc("/w", func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "world")
// 	})
// 	mux.HandleFunc("/sh", sayHello)
// 	http.ListenAndServe(":8080", mux)
// }
