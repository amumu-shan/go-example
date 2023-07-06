package routine

import (
	"fmt"
	"sync"
	"time"
)

func Start() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	ch := make(chan interface{})
	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, &wg, done, ch)
	}
	for i := 0; i < 10; i++ { // 向 ch 中发送数据，关闭 goroutine
		ch <- i
	}
	close(done)
	wg.Wait()
	fmt.Println("all done!")
}
func doIt(worked int, wg *sync.WaitGroup, done <-chan struct{}, ch <-chan interface{}) {
	fmt.Printf("[%v] is runing\n", worked)
	defer wg.Done()
	for {
		select {
		case m := <-ch:
			fmt.Printf("[%v] m => %v\n", worked, m)
			time.Sleep(1 * time.Second)
		case <-done:
			fmt.Printf("[%v] is done\n", worked)
			return
		}

	}
}

func UnCacheChan() {
	ch := make(chan string, 2)
	go func() {
		for m := range ch {
			fmt.Println("Processed:", m)
			time.Sleep(2 * time.Second)
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2"
	ch <- "cmd.3"
	ch <- "cmd.4"
	ch <- "cmd.5"
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func SingleChan() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
