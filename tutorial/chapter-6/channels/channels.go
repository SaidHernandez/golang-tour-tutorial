package channels

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //send sum to c
}

func BufferedChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 2 error en ejeución: fatal error: all goroutines are asleep - deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func CallSum() {
	fmt.Println("Channel...")
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

}
