// create a person struct
// ● create a func called “changeMe” with a *person as a parameter
// ○ change the value stored at the *person address
// ■ important
// Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 57● to dereference a struct, use (*value).field
// ○ p1.first
// ○ (*p1).first
// ● “As an exception, if the type of x is a named pointer type and (*x).f
// is a valid selector expression denoting a field (but not a method),
// x.f is shorthand for (*x).f.”
// ○ https://golang.org/ref/spec#Selectors
// ● create a value of type person
// ○ print out the value
// ● in func main
// ○ call “changeMe”
// ● in func main
// ○ print out the value
// code: https://play.golang.org/p/AehM662HKS

package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := person{
		Name: "sultan",
		Age:  19,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.Name = "Sheikh"
}
