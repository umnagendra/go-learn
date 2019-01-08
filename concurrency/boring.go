package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dchest/uniuri"
)

func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 20; i++ {
		select {
		case j := <-joe:
			fmt.Println(j)
		case a := <-ann:
			fmt.Println(a)
		}
		fmt.Println("--------")
		// fmt.Println(<-joe)
		// fmt.Println(<-ann)
	}
	fmt.Println("You're boring. I'm leaving.")
}

// generator function = function that returns a channel
func boring(who string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%d: %s says '%s'", i, who, uniuri.New())
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		}
	}()
	return c
}
