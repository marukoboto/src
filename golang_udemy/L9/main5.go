package main

import "fmt"

//スライス
//可変長引数

// ... にすることで可変長引数となり、引数の数を指定せずに関数をわたint型を0個以上受け取れる
//rangeで反復処理→nに足していく
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(Sum(1, 2, 3))

	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9,10))

	fmt.Println(Sum())

	//引数の数を指定せずに渡すことが出来る
	//６
	sl := []int{1,2,3}
	fmt.Println(Sum(sl...))


}
