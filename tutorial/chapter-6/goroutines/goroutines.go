package goroutines

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func CallSay() {
	fmt.Println("Goroutines...")
	go say("world")
	say("hello")
}
