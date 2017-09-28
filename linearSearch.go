// Linear Search Algorithm is a simple algorithm
// Compare TheNumToSearch with each item in the list sequentially

package main

import "fmt"

func linearsearch(n int, items []int) (bool, int) {
	for i := 0; i < len(items); i++ {
		if n == items[i] {
			return true, i + 1
		}
	}
	return false, 0
}

func main() {
	items := []int{6, 2, 9, 7, 3, 1, 8, 10, 4, 5}
	presence, position := linearsearch(10, items)
	if !presence {
		fmt.Println("The number you searched for is not present")
		return
	}
	fmt.Printf("The number is present at location %v\n", position)
}
