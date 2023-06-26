package main

import (
	"fmt"
	"sync"
)

// Golang 的高并发使用原生库来实现的话一般都是通过多 goroutine + select/channel，但是我们看 channel 源码，
// 发现这个东西就是一个队列+一把锁。这也就意味着无法避免多个 goroutine 带来的竞争问题。我之前测试过在多个 goroutine
// 竞争同一个 channel 的时候，性能急剧下降。所以很多高性能的高并发程序如果是用 Golang 来写，很多都会避免使用 channel 来传递数据，
//而是借用类似 disruptor 的 ringbuffer 技术。

//  解释

// Ringbuffer（环形缓冲区）是一种数据结构，用于在高并发环境中高效地传递数据。它是一块预分配的固定大小的内存区域，被视为环形的，可以被多个生产者和消费者并发地读写。
//
//Ringbuffer 的核心思想是使用指针来标记数据的读写位置，而不需要使用锁。这样就能避免由于锁竞争带来的性能损失。通常情况下，生产者会将数据写入环形缓冲区的尾部，而消费者会从头部读取数据。当指针到达缓冲区的尾部时，会回绕到头部，形成一个环。
//
//Ringbuffer 提供了两个指针：读指针和写指针。读指针表示下一个可读取的位置，写指针表示下一个可写入的位置。当读写指针相等时，表示缓冲区为空。当写指针紧邻读指针时，表示缓冲区已满。
//
//使用 Ringbuffer 进行并发数据传递时，生产者可以在不阻塞的情况下将数据写入缓冲区，而消费者可以在不阻塞的情况下从缓冲区读取数据。这使得并发程序可以以非常高效的方式进行数据交换，而不需要使用显式的锁机制。
//
//Ringbuffer 技术在一些高性能的高并发程序中得到了广泛应用，特别是在需要高吞吐量和低延迟的场景中。在 Golang 中，可以使用第三方库或自己实现 Ringbuffer 来优化高并发程序的性能，避免使用原生的通道（channel）带来的竞争问题。

type RingBuffer struct {
	buffer   []int
	size     int
	read     int
	write    int
	mutex    sync.Mutex
	notEmpty *sync.Cond
	notFull  *sync.Cond
}

func NewRingBuffer(size int) *RingBuffer {
	rb := &RingBuffer{
		buffer: make([]int, size),
		size:   size,
	}

	rb.notEmpty = sync.NewCond(&rb.mutex)
	rb.notFull = sync.NewCond(&rb.mutex)

	return rb
}

func (rb *RingBuffer) Put(value int) {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	for rb.isFull() {
		rb.notFull.Wait()
	}

	rb.buffer[rb.write] = value
	rb.write = (rb.write + 1) % rb.size

	rb.notEmpty.Signal()
}

func (rb *RingBuffer) Get() int {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	for rb.isEmpty() {
		rb.notEmpty.Wait()
	}

	value := rb.buffer[rb.read]
	rb.read = (rb.read + 1) % rb.size

	rb.notFull.Signal()

	return value
}

func (rb *RingBuffer) isFull() bool {
	return (rb.write+1)%rb.size == rb.read
}

func (rb *RingBuffer) isEmpty() bool {
	return rb.read == rb.write
}

func main() {
	bufferSize := 10
	rb := NewRingBuffer(bufferSize)

	// 启动生产者协程
	go func() {
		for i := 0; i < 20; i++ {
			rb.Put(i)
			fmt.Println("Produced:", i)
		}
	}()

	// 启动消费者协程
	go func() {
		for i := 0; i < 20; i++ {
			value := rb.Get()
			fmt.Println("Consumed:", value)
		}
	}()

	// 等待生产者和消费者协程完成
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
	}()

	wg.Wait()
}
