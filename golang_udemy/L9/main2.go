package main

import "fmt"

//スライス
//appned make len cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	sl = append(sl, 300)
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	sl2 := make([]int, 5)
	fmt.Println(sl2)
	
	//要素数を調べる
	fmt.Println(len(sl2))

	//cap(キャパシティー)で容量を算出する
	fmt.Println(cap(sl2))

	//第2引数が要素数、第３引数が容量
	sl3 := make([]int, 5, 10)

	//要素数を調べる
	fmt.Println(len(sl3))

	//cap(キャパシティー)で容量を算出する
	fmt.Println(cap(sl3))

	//slに7つ要素を追加する
	sl3 = append(sl3, 1,2,3,4,5,6,7)

	//要素数を調べる
	fmt.Println(len(sl3))

	//cap(キャパシティー)で容量を算出する
	//7つ追加したので、容量が20に自動的に増加する
	//自動的に増加した場合メモリの消費が倍になる
	fmt.Println(cap(sl3))
}
