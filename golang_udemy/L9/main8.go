package main

import "fmt"

//channel
//複数のゴルーチン間でのデータの受け渡しをする為に設計されたデータ構造。
//宣言、操作

func main() {
	//nilのチャネル
	var ch1 chan int 

	
	// //受信専用
	// var ch2 <- chan int

	// //送信専用
	// var ch3	chan <- int
	
	
	ch1 = make(chan int)

	ch2 := make(chan int)

	//容量を知る
	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	//バッファサイズ（容量）の大きさを指定
	//先入れ先出し
	//5以上になるとデッドロックが発生し、送信側がブロックされる。
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	//データを送る
	ch3 <- 1
	//送った要素数を確認
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	//3件追加したので３
	fmt.Println("len", len(ch3))


	i := <-ch3
	fmt.Println(i)
	//1件取り出したので２
	fmt.Println("len", len(ch3))

	i2 := <-ch3
	fmt.Println(i2)
	//もう一件取り出したので１
	fmt.Println("len", len(ch3))

	//この書き方でも行える
	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))


	ch3 <- 1
	fmt.Println("len", len(ch3))
	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6

}