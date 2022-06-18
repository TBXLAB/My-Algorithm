// Write a function called ThreeFive to print number from 1 to 100 in a new line.
// - if the number is divisible by 3 then print "Three"
// - if the number is divisible by 5 then print "Five"
// - if the number  is divisible by both 3 and 5 then print "ThreeFive"
// - Otherwise just print the number

package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("ThreeFive")
		} else if i%3 == 0 {
			fmt.Println("Three")
		} else if i%5 == 0 {
			fmt.Println("Five")
		} else {
			fmt.Println(i)
		}

	}

}
