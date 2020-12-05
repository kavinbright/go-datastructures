package main

import (
	"fmt"
	"structs/lib"
)

func main() {
	sl := lib.CreateSingleList()
	fmt.Printf("0.sl: %+v \n", sl)

	sl.Append(1)
	sl.Append(2)
	sl.Append(3)
	sl.Range()

	fmt.Printf("1.length: %d, sl: %+v \n", sl.Length(), sl)
	fmt.Printf("2.sl-head: %+v \n", sl.Head())
	fmt.Printf("3.sl-2: %+v \n", sl.Node(1))

	sl.Insert(100, 2)

	fmt.Printf("3.sl: %+v \n", sl)
	sl.Range()
	fmt.Printf("--------------------------------------------------------- \n")

	sl.Delete(2)
	sl.Range()

	fmt.Printf("--------------------------------------------------------- \n")

	sl.Delete(2)

	sl.Range()
}
