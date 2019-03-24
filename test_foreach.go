package main

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// wrong
	// for _, stu := range stus {
	// 	m[stu.Name] = &stu
	// }

	// right
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}

	for k, v := range m {
		println(k, "=>", v.Name, v.Age)
	}
}

// func main() {
// 	pase_student()
// }
