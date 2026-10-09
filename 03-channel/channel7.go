package main

import "fmt"

func main() {

	size := 100000
	numbers := make(chan int, size)
	primeNumbers := make(chan int, size)
	done := make(chan bool, size)

	go putNumber(size, numbers)
	go findPrimes(numbers, primeNumbers, done)
	for num := range primeNumbers {
		fmt.Println(num)
	}

}

func putNumber(size int, numbers chan int) {
	for i := 1; i <= size; i++ {
		numbers <- i
	}
	close(numbers)
}

func findPrimes(numbers chan int, primeNumbers chan int, done chan bool) {
	count := 0 // 記錄開了幾個 checkPrimes 協程
	for number := range numbers {
		go checkPrimes(number, primeNumbers, done)
		count++
	}
	for i := 0; i < count; i++ { // 開幾個就收幾次 done，等全部檢查完
		<-done
	}
	close(primeNumbers)
}

func checkPrimes(number int, primeNumbers chan int, done chan bool) {
	defer func() {
		done <- true
	}()
	if number < 2 {
		return
	} else {
		for i := 2; i < number; i++ {
			if number%i == 0 {
				return
			}
		}
		primeNumbers <- number
	}

}

/*
流程：
1. main 建立 3 個 channel：numbers（要檢查的數字）、primeNumbers（找到的質數）、done（完成通知）。
2. putNumber（1 個協程）：把 1 ~ size 依序放進 numbers，放完 close(numbers)。
3. findPrimes（1 個協程）：range numbers 每拿到一個數字，就開一個 checkPrimes 協程，共開 size 個。
4. checkPrimes：是質數就放進 primeNumbers；不管是不是質數，結束時 defer 都會送一個 true 到 done。
5. findPrimes 收滿 count 個 done，代表全部檢查完，才 close(primeNumbers)。
6. main 用 range primeNumbers 印出質數，primeNumbers 被 close 後迴圈結束，程式結束。

重點：
- done 的作用跟 WaitGroup 一樣：等所有 checkPrimes 跑完。
- 沒等 done 就 close(primeNumbers)，還在跑的 checkPrimes 會送到已關閉的 channel → panic。
- 原本用 cap(done) 當等待次數，容量一改就對不上，所以改成自己數 count。
- 印出的質數不一定照大小排序，因為各個 checkPrimes 誰先跑完不一定。
*/
