package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{} // 宣告
	job := 100
	for i := 1; i <= job; i++ {
		wg.Add(1)
		go doTask(wg)
	}
	wg.Wait()
	fmt.Println("All jobs done")
}

func doTask(wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("one task done")
		wg.Done()
	}() //defer must be function call
	number := rand.Intn(10) + 1
	fmt.Println(number)
	time.Sleep(time.Duration(number) * time.Second)

}

/*
重點：
1. main 跑迴圈，每次 wg.Add(1) 再 go doTask(wg)，共開 100 個 goroutine。
2. wg.Wait() 會擋住 main，直到 100 個都呼叫過 wg.Done()（計數器歸零）才往下印 "All jobs done"。
3. defer 是在「doTask 這個函式結束時」才執行，不是整個程式結束時。
   所以每個 goroutine 都是：印 number → Sleep → 函式要結束了 → 跑 defer（印 "one task done"、wg.Done()）。
4. wg.Done() 只是把計數器減 1，不是刪掉 goroutine；goroutine 是 doTask return 後自己結束。
5. 100 個 goroutine 同時跑：一開始 100 個 number 幾乎一起印出來，
   之後 Sleep 1 秒的先印 "one task done"，Sleep 10 秒的最後，總共約 10 秒，不是全部秒數加起來。
6. defer 後面要接函式呼叫，所以匿名函式最後要加 ()。
   把 wg.Done() 放在 defer 裡，就算中途 return 或 panic 也一定會執行，不會卡住 Wait()。
*/
