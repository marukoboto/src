package main

import "fmt"

//スライス
//copy

func main() {
	// //参照型（スライス型）の場合書き換えが起こる
	// sl := []int{100,200}
	// sl2 := sl

	// sl2[0] = 1000
	// fmt.Println(sl)
	
	// //基本型の場合　書き換えは起こらない
	// var i int = 10
	// i2 := i
	// i2 = 100
	// fmt.Println(i, i2)

	//slにint型で1,2,3,4,5が作られる
	sl := []int{1,2,3,4,5}
	//sl2に5つ分初期値（00000）を作る
	sl2 := make([]int, 5, 10)
	fmt.Println(sl2)

	//sl部分をsl2にコピーする
	//n にコピーされた要素数を貼る
	n := copy(sl2, sl)
	
	fmt.Println(n, sl2)

	//s3にint型で1,2,3,4,5が作られる
	s3 := []int{6,7,8,9,10}
	//sl2に4つ分初期値（0000）を作った場合ははみ出す
	sl4 := make([]int, 4, 10)
	fmt.Println(sl4)

	//sl4部分をs3にコピーする
	//t にコピーされた要素数を貼る
	t := copy(sl4, s3)
	
	fmt.Println(t, sl4)
}