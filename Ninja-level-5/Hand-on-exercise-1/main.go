// Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
// 		● first name
// 		● last name
// 		● favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

package main

import "fmt"

type person struct {
	first         string
	last          string
	favourflavour []string
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		favourflavour: []string{

			"Venila",
			"chocolate",
			"mint",
		},
	}

	p2 := person{
		first: "Mike",
		last:  "Tyson",
		favourflavour: []string{
			"chocolate",
			"vinila",
			"adrak",
		},
	}
	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range p1.favourflavour {
		fmt.Println(i, v)
	}

	for i, v := range p2.favourflavour {
		fmt.Println(i, v)
	}
}
