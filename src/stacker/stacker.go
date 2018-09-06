// stacker.go
// Perform operations on a stack
package main

import (
	"fmt"
	"os"
	"stacker/stack"
)

var hayStack stack.Stack

func pushDefaultElements() {
	hayStack.Push(19)
	hayStack.Push("hello")
	hayStack.Push([]int{1, 3, 9})
	hayStack.Push(-23.457462)
	hayStack.Push(true)
	hayStack.Push([]string{"heya", "not", "cool"})
}

func pushCmdLineElements() {
	if len(os.Args) > 1 {
		for param := range os.Args {
			// we don't want to push os.Args[0] itself
			if param == 0 {
				continue
			}
			fmt.Printf("Pushing %s into stack ...\n", os.Args[param])
			hayStack.Push(os.Args[param])
		}
	}
}

func pushElements() {
	pushDefaultElements()
	pushCmdLineElements()
}

func printStackProperties() {
	fmt.Printf("STACK LENGTH = %d\n", hayStack.Len())
	fmt.Printf("STACK CAPACITY = %d\n", hayStack.Cap())
}

func printStackContents() {
	fmt.Println("---- START ----")
	for {
		item, err := hayStack.Pop()
		if err == nil {
			fmt.Println(item)
		} else {
			fmt.Println("---- DONE ----")
			break
		}
	}
}

func main() {
	pushElements()
	printStackProperties()
	printStackContents()
}
