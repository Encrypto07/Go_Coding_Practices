package main

import "fmt"

func main() {
	foo()
	boo("sultan", "sheikh")
	v := coo("Ashutosh", " singh")
	fmt.Println(v)

	a, b := goo("Zeeshan ", "Hayat Nagar ", 19)
	fmt.Println(a, b)
}

func foo() {
	fmt.Println("Hello everyone")
}

func boo(first, last string) {
	fmt.Println("Hello", first, last)
}

func coo(Fname, Lname string) string {
	return fmt.Sprint(Fname, Lname, ` say, "Hello"`)
}

func goo(name, address string, age int) (string, bool) {
	p1 := fmt.Sprint(name, address, age)
	p2 := false
	return p1, p2
}
