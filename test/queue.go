package main

import (
	"fmt"
	"time"
)

type FixBuffer struct {
	chann       chan interface{}    // 缓冲管道
	consumer    func([]interface{}) // 消费者方法
	consumerNum int32               // 消费者数量
}

func main() {

	loopTimes := 10
	for i := 1; i <= loopTimes; i++ {
		test := &testObj{
			Num:  int32(i),
			Data: fmt.Sprintf("I am data:%d", i),
		}

		NewQueue <- test
	}

	time.Sleep(2 * time.Second)
	// test := &testObj{

	// }
	// NewQueue <-
}

func Init(capactity int32, consumerNum int32, consumer func([]interface{})) *FixBuffer {
	b := &FixBuffer{
		consumer:    consumer,
		consumerNum: consumerNum,
	}

	if capactity > 0 {
		b.chann = make(chan interface{}, capactity)
	} else {
		b.chann = make(chan interface{})
	}

	return b
}

type testObj struct {
	Num  int32  `json:"num"`
	Data string `json:"data"`
}

var (
	// 处理队列，缓存大小为1即可
	NewQueue = make(chan *testObj, 1)
)

func init() {
	go doNewQueue()
}

func doNewQueue() {
	// 阻塞循环
	for {
		o := <-NewQueue
		fmt.Println(">>>>sleep 1s")
		time.Sleep(1 * time.Second)
		// go Do(o.Num, o.Data)
		go Do(o.Num, o.Data)
	}
}

func Do(num int32, data string) (err error) {
	// time.Sleep(1 * time.Second)
	fmt.Printf("==num:%d, data:%s do...\n", num, data)
	return
}
