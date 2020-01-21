// Print every rune code point of the uppercase alphabet three times.package ninja3
// output should look like this
// U+0041 'A'

package main

import "fmt"

func main() {

	for i := 65; i <= 90; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%#U", i)
			fmt.Println("")
		}
		fmt.Println("")
	}

}
