//Create and use an anonymous struct

package main

import "fmt"

func main() {
	p1 := struct { //ananymous struct i.e. it does not contain any name
		first string
		last  string
		age   int
	}{
		first: "sultan",
		last:  "sheikh",
		age:   25,
	}
	fmt.Println(p1)
}
