
package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(intSlice)
	fmt.Println(intSlice) // [1 2 3 4 5 6]

sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
	fmt.Println( intSlice)


	multipleWords := []string{"Abel","Caesar" , "Baal", "Daniel"}
fmt.Println("Printing the biblical names without sorting", multipleWords)


sort.StringSlice(multipleWords).Sort()
fmt.Println("Printing the biblical names after sorting in ascending order", multipleWords)
 sort.Sort(sort.Reverse(sort.StringSlice(multipleWords)))

	fmt.Println("Now printing the biblical names in reverse order ",multipleWords )


}
