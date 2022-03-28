package main

import "fmt"

type person struct {
	first_name string
	last_name  string
}

func (p *person) speek() { //whenfunction having a structure it become method or it also called receiver function
	fmt.Println(p.first_name, p.last_name)
}

func main() {

	p0 := struct {
		super_Hero_name     string
		power_of_super_hero string
		weekness            string
	}{
		super_Hero_name:     "Thor",
		power_of_super_hero: "Thunder",
		weekness:            "without hammer thor is week",
	}

	fmt.Println(p0)

	p1 := person{
		first_name: "sultan",
		last_name:  "sheikh",
	}
	p1.speek()
}
