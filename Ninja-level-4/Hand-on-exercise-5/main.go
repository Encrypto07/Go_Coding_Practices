// To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise,
// follow these steps:
// 		● start with this slice
// 				○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// 		● use APPEND & SLICING to get these values here which you should ASSIGN to avariable “y” and then print:
//				○ [42, 43, 44, 48, 49, 50, 51]

package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	xi = append(xi[:3], xi[6:]...)

	fmt.Println(xi)
}
