package main

import "fmt"

func main() {
	// test1()

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	fmt.Println("====s1: ", s[:len(s)/2])
	go test2(s[:len(s)/2], c)
	// fmt.Println("===c1: ", c)
	fmt.Println("====s2: ", s[len(s)/2:])
	go test2(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}

func test1() {
	c := make(chan int)
	// defer close(c)
	go func() {
		c <- 3 + 4
		// close(c)
		c <- 5 + 4
		defer close(c)
	}()

	i := <-c
	fmt.Println(i)
	// c <- 5 + 4
	// j := <-c
	// fmt.Println(j)

	x, ok := <-c
	if !ok {
		fmt.Println("get channel 1 not ok")
	}

	fmt.Println(x)

	x, ok = <-c
	if !ok {
		fmt.Println("get channel 2 not ok")
	}

	fmt.Println(x)
}

func test2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		fmt.Println("v: ", v)
		sum += v
	}

	c <- sum
}
