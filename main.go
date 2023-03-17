package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 生产者
func producerCopy(data string, channel chan<- string) {
	for {
		channel <- fmt.Sprintf("%s: %v", data, rand.Int31())
		time.Sleep(time.Second)
	}
}

// 消费者
func customerCopy(channel <-chan string) {
	for {
		message := <-channel
		fmt.Println(message)
	}
}

func main() {
	//// 创建一个字符串类型的通道
	//channel := make(chan string)
	//// 创建producer()函数的并发goroutine
	//go producerCopy("cat", channel)
	//go producerCopy("dog", channel)
	//// 数据消费函数
	//customerCopy(channel)

}

// 多线程n++ 问题
func manyN() {
	wg := sync.WaitGroup{}
	locker := sync.Mutex{}
	n := 0
	for i := 0; i < 10; i++ {
		wg.Add(1) // 开一个进程加1
		go func() {
			defer wg.Done() // 结束一个
			defer locker.Unlock()
			locker.Lock()
			n++
		}()
	}
	wg.Wait()
	fmt.Println(n)
}
