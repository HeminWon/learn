package test

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// for i := 0; i < 10; i++ {
		// 	data := <-ch
		// 	fmt.Println(data)
		// }

		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChanel(t *testing.T) {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, wg)
	wg.Add(1)
	dataReceiver(ch, wg)
	wg.Add(1)
	dataReceiver(ch, wg)
	wg.Wait()
}
