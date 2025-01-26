package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

type Programmer struct {
	Name string
	Domain string
	Languages string
}


func main() {
	a := add(2, 3)
	b := &a
	c := Programmer{Name: "Raikou 320", Domain: "Websites", Languages: "HTML, CSS, JS, Python, TS"}
	d := &c
	if *b == 5 {
		fmt.Println("Test 1 Passed")
	}
	*b = 10
	if a == 10 {
		fmt.Println("Test 2 Passed")
	}
	for i := 0; i < 3; i ++ {
		fmt.Println(fmt.Sprintf("%d %s", i, d))
	}
}