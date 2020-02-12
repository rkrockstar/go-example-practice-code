// using a composite literal
// create an ARRAY which holds 5 values of TYPE int
// assign VALUES to each index position
// Range over the array and print the values out
// using format printing
// print out the TYPE of the array

package main

import "fmt"

func main() {
	a := [5]int{42, 43, 44, 45, 46}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)

}
