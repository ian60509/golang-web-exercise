package main 

import (	
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	var a []int            // nil slice
	b := []int{0, 1, 2, 3} // slice initialized with specified array
	c := make([]int, 4)    // slice of size 4 initialized with zero-valued array of size 4
	d := make([]int, 4, 5) // slice of size 4 initialized with zero-valued array of size 5

	fmt.Printf("a: length: %d, capacity: %d, pointer to underlying array: %p, data: %#v, is nil: %t\n", len(a), cap(a), a, a, a == nil)
	fmt.Printf("b: length: %d, capacity: %d, pointer to underlying array: %p, data: %v, is nil: %t\n", len(b), cap(b), b, b, b == nil)
	fmt.Printf("c: length: %d, capacity: %d, pointer to underlying array: %p, data: %v, is nil: %t\n", len(c), cap(c), c, c, c == nil)
	fmt.Printf("d: length: %d, capacity: %d, pointer to underlying array: %p, data: %v, is nil: %t\n", len(d), cap(d), d, d, d == nil)
}