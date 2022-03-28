// Take the code from the previous exercise, then store the values of type person in a map with the
// key of last name. Access each value in the map. Print out the values, ranging over the slice.

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
	fmt.Println(p2)
	m := map[string]person{
		p1.first: p1,
		p1.last:  p1,
	}

	for i, v := range m {
		fmt.Println(i, v)
	}
}
