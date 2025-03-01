package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
    n := len(a)
    
    // If no numbers are passed, return an empty slice
    if len(nbrs) == 0 {
        return nil
    }
    
    start := nbrs[0]
    end := n // Default end is the length of the slice
    
    // If there are two numbers, set end
    if len(nbrs) > 1 {
        end = nbrs[1]
    }
    
    // Handle negative indices
    if start < 0 {
        start = n + start
    }
    if end < 0 {
        end = n + end
    }
    
    // Ensure start and end are within bounds
    if start < 0 {
        start = 0
    }
    if start > n {
        start = n
    }
    if end < 0 {
        end = 0
    }
    if end > n {
        end = n
    }
    
    // Return the slice from start to end
    if start < end {
        return a[start:end]
    }
    
    // If start > end, return an empty slice
    return nil
}
func main() {
    a := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Printf("%#v\n", Slice(a, 1))           // ["algorithm", "ascii", "package", "golang"]
    fmt.Printf("%#v\n", Slice(a, 2, 4))        // ["ascii", "package"]
    fmt.Printf("%#v\n", Slice(a, -3))          // ["ascii", "package", "golang"]
    fmt.Printf("%#v\n", Slice(a, -2, -1))      // ["package"]
    fmt.Printf("%#v\n", Slice(a, 2, 0))        // nil
}