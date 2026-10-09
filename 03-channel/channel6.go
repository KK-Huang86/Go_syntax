package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	intChan := make(chan int, 1)
	go generateNumbers(intChan)
	for {
		if num, ok := <-intChan; ok {
			fmt.Printf("isSuccess: %v value: %d\n", ok, num)
			time.Sleep(time.Second)

		} else {
			break
		}
	}
}

/*
重點：
1. 跟 channel5 一樣，goroutine 放 6 個數字後 close；這次 main 改用 num, ok := <-intChan 手動收。
2. ok 是 true：有收到值。
   ok 是 false：channel 已經 close 而且收完了，這時 num 是零值 0，用 break 離開迴圈。
3. 效果跟 for num := range intChan 一樣，range 只是幫你把 ok 的判斷寫好了。
4. 拿掉 close：ok 永遠不會變 false，main 一直等 → deadlock。
5. Println 不吃 %v、%d，要格式化輸出請用 Printf。
*/
