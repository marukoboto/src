package main

import "fmt"

//スライス
//宣言、操作

func main() {
	//整数型のスライス　カッコ内に要素数を指定しない
	var sl []int
	fmt.Println(sl)

	//数値を格納したスライスを作成
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	//要素数分のデータ型の初期値
	//要素数分のint型の初期値　　
	sl4 := make([]int, 5)
	fmt.Println(sl4)

	//sl2に1000の代入
	sl2[0] = 1000
	fmt.Println(sl2)

	//スライスの作成
	sl5 := []int{1,2,3,4,5}
	fmt.Println(sl5)
	
	//0番目　1が出力
	fmt.Println(sl5[0])

	//2~4番目を出力
	fmt.Println(sl5[2:4])

	//2番目までを出力
	fmt.Println(sl5[:2])

	//2番目以降を出力
	fmt.Println(sl5[2:])

	//配列全部を摘出
	fmt.Println(sl5[:])

	//配列の最後のみ摘出
	fmt.Println(sl5[len(sl5)-1])

	//配列の最初と最後のみ摘出
	//配列として出てくるため、カッコが付く
	fmt.Println(sl5[1 : len(sl5)-1])
}