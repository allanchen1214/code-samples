package main

import (
	"fmt"
	"sync"
)

var (
	counter, counter2, counter3 int
	l                           sync.Mutex
	mylock                      Lock
)

func TestMethod1() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func TestMethod2() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter2++
			l.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter2)
}

type Lock struct {
	ch chan struct{}
}

func NewLock() Lock {
	var l Lock
	l.ch = make(chan struct{}, 1)
	l.ch <- struct{}{}
	return l
}

func (l Lock) Lock() bool {
	b := false
	select {
	case <-l.ch:
		b = true
	default:
	}
	return b
}

func (l Lock) Unlock() {
	l.ch <- struct{}{}
}

func TestMethod3() {
	mylock = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if b := mylock.Lock(); !b {
				return
			}
			counter3++
			mylock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter3)
}

// 基于redis setnx的分布式锁
// 基于ZK的分布式锁
// 基于etcd的分布式锁

func main() {
	TestMethod1() // 不稳定输出
	TestMethod2() // 稳定输出10000
	TestMethod3() // trylock机制，不稳定输出，但同时只能有一个goroutine在执行
}
