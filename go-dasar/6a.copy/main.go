package slice

import "fmt"

func main() {
	items := []string{"A", "B", "C", "D", "E"}
	i := 2 // Deleting "C"

	// The destination is the part of the slice starting at the deletion point:
	// items[i:] is ["C", "D", "E"] (The view starting at the slot we want to fill)

	// The source is the part of the slice AFTER the element we want to delete:
	// items[i+1:] is ["D", "E"] (The data we want to shift left)

	copy(items[i:], items[i+1:])
	// println(items)
	fmt.Printf("Destination slice: %v\n", items) // Output: Destination slice: [10 20 30]

}

func lenTest (){
	
}