// create a program that uses a switch statement with noswitch expression specified

package main

func main() {
	i := 1

	switch {
	case false:
		print("this should not print")
	default:
		print(i)
	}
}
