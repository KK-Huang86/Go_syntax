# go-learning

自學 Go 的練習紀錄，邊看邊寫，筆記直接寫在程式碼的註解裡。

## 目錄

### 01-basics

基本語法：`package main`、`import`、函式宣告與回傳值。

### 02-goroutines

- `goroutines1.go`：用 `go` 開第一個 goroutine，觀察 hello / world 印出順序不固定
- `goroutines2.go`：不用 goroutine，單純跑迴圈找 0 ~ 1000000 的質數，大約 21 秒
- `goroutines3.go`：改成每個數字開一個 goroutine，發現 main 結束後沒跑完的 goroutine 會直接被砍掉
- `waitgroup1.go`：用 `sync.WaitGroup` 等全部 goroutine 跑完，順便踩到 `Add` 數量對不上造成的 deadlock
- `waitgroup2.go`：`defer` 搭配 `wg.Done()`，觀察 100 個任務同時跑、依 Sleep 長短陸續結束

### 03-channel

- `channel1.go`：`make` 給容量、`len` / `cap`、先進先出，還有沒容量的 channel 為什麼會卡住
- `channel2.go`：channel 放 struct、`interface{}`，以及只能送 / 只能收的單向 channel
- `channel3.go`：`close` 之後不能再送（會 panic），但剩下的值還收得到
- `channel4.go`：`for range` 讀 channel，沒 `close` 會 deadlock
- `channel5.go`：一個 goroutine 放、main 收，容量 1 時送的一方會等，不會塞爆
- `channel6.go`：用 `num, ok := <-ch` 判斷 channel 是不是已經關了
- `channel7.go`：用三個 channel 串起來找 10 萬以內的質數，`done` 拿來等所有 goroutine 跑完再 `close`

## 執行

```bash
go run 03-channel/channel7.go
```

同一個資料夾裡每個檔案都有自己的 `main`，所以要一次跑一個檔案，`go run .` 會報錯。

## 接下來

- `select`
- worker pool
