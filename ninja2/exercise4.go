// assign an int to a variable
// print that int in decimal, binary, hex
// shift the bits of that int over 1 position to the left, and assigns that to a variable
// prints that variable in decimal, binary and hex

package main

import "fmt"

func main() {
	n := 42
	fmt.Printf("%d\t%b\t%#x\t", n, n, n)
	m := n << 1

	fmt.Printf("%d\t%b\t%#x", m, m, m)

}
