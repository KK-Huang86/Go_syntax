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

## 執行

```bash
go run 03-channel/channel1.go
```

## 接下來

- `close` 和 `range` 讀 channel
- `select`
