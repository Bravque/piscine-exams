// Write a function that returns the number of elements in the linked list that's
// passed to it.

package main

import "fmt"

// Define the Node structure
type Node struct {
	value int
	next  *Node
}

// Function to count the number of elements in the linked list
func countElements(head *Node) int {
	count := 0
	current := head

	// Traverse the list and count nodes
	for current != nil {
		count++
		current = current.next
	}

	return count
}

func main() {
	// Create a simple linked list
	node1 := &Node{value: 10}
	node2 := &Node{value: 20}
	node3 := &Node{value: 30}
	node1.next = node2
	node2.next = node3

	// Call the countElements function
	result := countElements(node1)
	fmt.Println("Number of elements in the linked list:", result)
}
