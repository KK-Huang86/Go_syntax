package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 主協程與 分支協程之間的運作情況

func generateNumbers(intChan chan int) {
	count := 0
	for {
		if count != 6 {
			num := rand.Intn(10) + 1
			intChan <- num
			fmt.Println("put number:", num)
			count++
		} else {
			close(intChan)
			break
		}

	}
}

func main() {
	intChan := make(chan int, 10)
	go generateNumbers(intChan)
	for num := range intChan {
		fmt.Println("get number:", num)
		time.Sleep(time.Second)

	}
}

/*
重點：
1. goroutine 負責放 6 個數字，放完 close；main 用 range 收，兩邊同時跑。
2. channel 滿了，送的一方會等 main 收走，不會塞爆。
3. range 收到 channel 被 close 且收完才結束。

deadlock：所有 goroutine 都卡住，沒人能往下跑。
- 拿掉 close：main 一直等第 7 個，沒人送 → deadlock。
- 在 main 送資料但 channel 沒空間又沒人收，或收資料但沒人送 → deadlock。
*/
