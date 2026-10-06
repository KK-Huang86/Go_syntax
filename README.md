# go-learning

自學 Go 的練習紀錄，邊看邊寫，筆記直接寫在程式碼的註解裡。

## 目錄

### 01-basics

基本語法：`package main`、`import`、函式宣告與回傳值。

### 02-goroutines

- `goroutines1.go`：用 `go` 開第一個 goroutine，觀察 hello / world 印出順序不固定
- `goroutines2.go`：不用 goroutine，單純跑迴圈找 0 ~ 1000000 的質數，大約 21 秒
- `goroutines3.go`：改成每個數字開一個 goroutine，發現 main 結束後沒跑完的 goroutine 會直接被砍掉

## 執行

```bash
go run 02-goroutines/goroutines1.go
```

## 接下來

- 用 `sync.WaitGroup` 等所有 goroutine 跑完
- channel
