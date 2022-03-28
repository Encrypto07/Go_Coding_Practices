//Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	delete(m, `bond_james`)

	for i, v := range m {
		fmt.Println(i, v)
	}
}
