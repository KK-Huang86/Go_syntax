package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	num := 1000000
	start := time.Now().Unix()
	for i := 0; i < num; i++ {
		wg.Add(1) // 每開一個協程就 +1，數量會對得上
		go findPrime(i, wg)
	}
	wg.Wait() // block main 主程式的協程運作，等其他協程都處理完
	end := time.Now().Unix()
	fmt.Println(end-start, "seconds")

}

func findPrime(num int, wg *sync.WaitGroup) {
	defer wg.Done() //該協程結束後會呼叫 waitgroup 減少協程開的數量
	if num < 2 {    // 0 和 1 都不是質數
		return
	} else if num == 2 {
		fmt.Println(num)
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
1. WaitGroup 是一個計數器：Add(n) 加 n，Done() 減 1，Wait() 會卡住直到歸零。
2. 每開一個 goroutine 前 wg.Add(1)，goroutine 結束時 defer wg.Done()，數量才會對得上。
3. 一開始寫 wg.Add(num) 但迴圈從 1 開始，只開了 999999 個，計數器剩 1 永遠不歸零
   → Wait() 一直等 → fatal error: all goroutines are asleep - deadlock!
4. 跟 goroutines3.go 比：不用 time.Sleep 猜時間，main 會等全部算完才結束，印出的質數數量每次都一樣。
5. 量到的時間才是「全部算完」的時間。
*/
