// Using the operators, wite expressions and assign their values to variables

package main

import "fmt"

func main() {

	n := 4
	a := (n == n)
	b := (n <= 45)
	c := (n >= 45)
	d := (n != 46)
	e := (n < 47)
	f := (n > 48)
	fmt.Println(a, b, c, d, e, f)
}
