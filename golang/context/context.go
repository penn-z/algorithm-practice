package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	/*
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		// go handle(ctx, 500*time.Millisecond)
		go handle(ctx, 2000*time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println("main", ctx.Err())
		}
	*/

	test1()

}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)

	}
}

func test1() {
	var keyA string = "keyA"
	ctx := context.Background()
	ctxA := context.WithValue(ctx, keyA, "valA")

	var keyC string = "keyA"
	ctxC := context.WithValue(ctxA, keyC, "valC")

	println("keyA 1 => ", ctxA.Value(keyA).(string))
	println("keyA => ", ctxC.Value(keyA).(string))
	println("keyC => ", ctxC.Value(keyC).(string))
}
