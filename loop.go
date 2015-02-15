// Loops project loop.go
package main

import (
	"fmt"
)

func main() {
	mySlice := []int{178, 58, 5, 0, 285, 310}
	for i, currentEntry := range mySlice {
		fmt.Println(i, " - ", currentEntry)
	}

	myArray := []int{2, 3, 1, 5, 3, 4}
	key := 3
	for i := 0; i < 6; i++ {
		//var curr = myArray[i]
		if key == myArray[i] {
			fmt.Println("Key is first shown at index:", i)
			break
		}
	}

}
