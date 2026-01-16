package goroutings

import (
	"fmt"
	"sync"
)

func WorkGroup() {
	wg := new(sync.WaitGroup)
	inputStream := make(chan string)
	outputStream := make(chan string)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go removeDuplicates(inputStream, outputStream, wg)
	}
	go func() {
		defer close(inputStream)

		for _, r := range "q112334456" {
			inputStream <- string(r)
		}
	}()
	for x := range outputStream {
		fmt.Print(x)
	}
	close(outputStream)
	wg.Wait()
}

func removeDuplicates(c1, c2 chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(c2)
	seen := make(map[string]struct{})
	for v := range c1 {
		_, ok := seen[v]
		if ok {
			continue
		}
		seen[v] = struct{}{}
	}

	for str, _ := range seen {
		c2 <- str
	}
}

func AnonymousGoroutines() {
	c := make(chan struct{})
	go func(c chan struct{}) {
		//go work()
		close(c)
	}(c)
	<-c
}
