package main

import "fmt"

//スライス
//for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// //インデックス番号と字を指定
	// for _, v := range sl {
	// 	fmt.Println( v)
	// }

	//iがsl以下の場合iを増やしていき、
	//iを出力していく
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}