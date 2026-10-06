package main

import (
	"fmt"
	"time"
)

func main() {
	num := 1000000
	start := time.Now()
	for i := 0; i < num; i++ {
		go findPrimes(i)
	}

	end := time.Now()
	fmt.Println(end.Unix()-start.Unix(), "seconds")
	time.Sleep(5 * time.Second)
}

func findPrimes(num int) {
	if num == 1 {
		return
	} else if num < 2 {
		return
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return
			}
		}
		fmt.Println(num)
	}
}

/*
重點：
1. 每次 go findPrimes(i) 開一個新的 goroutine，共一百萬個，分到多核心並行跑，所以輸出順序不固定。
2. main 開完就不管（fire-and-forget），不會追蹤哪些 goroutine 跑完。
3. start ~ end 量到的只是「開完一百萬個 goroutine」的時間，不是全部算完的時間。
4. main 一 return，程式就結束，沒跑完的 goroutine 直接被殺掉。
5. time.Sleep 只是用猜的時間等，可能不夠也可能白等；要精確等全部完成，請用 sync.WaitGroup。
*/
