// using a composite literal
// create a slice of TYPE int
// assign 10 values
// range over the slice and print the values out
// using format printing
// print out the type of the array

package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)

}
