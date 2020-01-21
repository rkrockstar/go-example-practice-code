// create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER "favSport".package ninja3

package main

import "fmt"

func main() {

	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains")
	case "surfing":
		fmt.Println("go to the beach")
	}

}
