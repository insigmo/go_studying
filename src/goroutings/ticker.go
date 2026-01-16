package goroutings

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//
//func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
//	wg := &sync.WaitGroup{}
//	mt := &sync.Mutex{}
//	var nums1, nums2 []int
//	go func() {
//		for i := 0; i < n; i++ {
//			wg.Add(2)
//			go func() {
//				defer wg.Done()
//				mt.Lock()
//				nums1 = append(nums1, <-in1)
//				mt.Unlock()
//			}()
//
//			go func() {
//				defer wg.Done()
//				mt.Lock()
//				nums2 = append(nums2, <-in2)
//				mt.Unlock()
//			}()
//		}
//
//		wg.Wait()
//
//	}()
//
//	go func() {
//		for i := 0; i < n; i++ {
//			go func(index int) {
//				wg.Add(2)
//				var x1, x2 int
//
//				go func() {
//					defer wg.Done()
//					x1 = fn(<-in1)
//				}()
//
//				go func() {
//					defer wg.Done()
//					x2 = fn(<-in2)
//				}()
//
//				wg.Wait()
//				out <- x1 + x2
//			}(i)
//		}
//	}()
//
//}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	res := make([]int, n)
	wg := new(sync.WaitGroup)
	wg.Add(n)

	go func(wg *sync.WaitGroup) {
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2
			go func() {
				defer wg.Done()
				c1 := make(chan int)
				c2 := make(chan int)

				go func() { c1 <- fn(x1) }()
				go func() { c2 <- fn(x2) }()

				res[i] = <-c1 + <-c2
			}()
		}

		wg.Wait()
		for i := 0; i < n; i++ {
			out <- res[i]
		}
	}(wg)
}

//func merge2Channels(
//	fn func(int) int,
//	in1 <-chan int,
//	in2 <-chan int,
//	out chan<- int,
//	n int,
//) {
//	results := make([]int, n)
//	var wg sync.WaitGroup
//	wg.Add(n)
//
//	go func() {
//		for i := 0; i < n; i++ {
//			x1 := <-in1
//			x2 := <-in2
//			go func() {
//				defer wg.Done()
//
//				c1 := make(chan int, 1)
//				c2 := make(chan int, 1)
//
//				go func() { c1 <- fn(x1) }()
//				go func() { c2 <- fn(x2) }()
//
//				results[i] = <-c1 + <-c2
//			}()
//		}
//
//		wg.Wait()
//
//		for i := 0; i < n; i++ {
//			out <- results[i]
//		}
//	}()
//}

func GoMain() {
	const N = 5
	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Millisecond)
		return x * x
	}
	fmt.Println("\nСоздаю каналы, начитаю работу.")
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N)
	for i := 0; i < N; i++ {
		in1 <- i
		in2 <- (i + N)
		fmt.Printf("%v) В канал 1 послано: %v, В канал 2 послано: %v.\n", i+1, i, (i + N))
	}
	fmt.Print("Функция fn с разной паузой возводит в квадрат значения из каналов.\n",
		"Итоговый результат fn(x1) + fn(x2)\n",
		"Барабанная дробь...\n")

	calcFail := false
	timeFail := false
	pass := "OK"
	for i := 0; i < N; i++ {
		c := <-out
		result := i*i + (i+N)*(i+N)
		if c != result {
			calcFail = true
			pass = "Failed."
		}
		fmt.Printf("%v) Результат: %v. Правильный %v. -> %v\n", i+1, c, result, pass)
	}
	if calcFail {
		fmt.Println("Failed. Результат выполнения не соответствует ожидаемому.")
	}
	duration := time.Since(start)
	if duration.Milliseconds() > N {
		timeFail = true
		fmt.Println("Failed. Время превышено.")
	}
	if !calcFail && !timeFail {
		fmt.Println("Passed. OK. Ты сделал это!\nВремя работы: ", duration)
	} else {
		fmt.Println("Пока неверно. Упорства вам не занимать, попробуете еще раз?\nВремя выполнения: ", duration)
	}
}

func Ticker() {
	time.Sleep(2 * time.Second)
	timer := time.After(time.Second)
	<-timer

	ticker := time.Tick(time.Second)
	counter := 0

	for {
		a := <-ticker
		counter++
		fmt.Println(a)

	}
}

func Timer() {
	t := time.NewTimer(time.Second)
	go func() {
		a := <-t.C
		fmt.Println(a)
	}()
	t.Stop()
	t.Reset(time.Second)
	fmt.Println(<-t.C)
}
func GoMain1() {
	arguments := make(chan int)
	done := make(chan struct{})
	result := calculator(arguments, done)
	for i := 0; i < 10; i++ {
		arguments <- 1
	}
	close(done)
	fmt.Println(<-result)
}
func DownloadAccounts() {
	// Download accounts from something where you can download only every 350ms

	ticker := time.NewTicker(350 * time.Millisecond)
	defer ticker.Stop()

	wg := new(sync.WaitGroup)
	defer wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go downloadAccount(i, ticker.C, wg)
	}

}

func downloadAccount(id int, delay <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	t := <-delay
	fmt.Printf("Downloading account #%d ..., time=%s\n", id, t.Format(time.RFC3339Nano))

}

func calculator1(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	out := make(chan int)
	go func(c chan int) {
		defer close(c)
		select {
		case v := <-firstChan:
			c <- v * v
		case a := <-secondChan:
			c <- a * 3
		case <-stopChan:
			return
		}
	}(out)
	return out
}
func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func(c chan int) {
		defer close(c)
		sum := 0
		for {
			select {
			case v := <-arguments:
				sum += v
			case <-done:
				c <- sum
				return
			}
		}
	}(out)
	return out
}
