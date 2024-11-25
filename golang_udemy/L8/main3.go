package main

import "fmt"

//for
//繰り返し処理

func main() {
	// i := 0
	// for {
	// 	i++
	// 	if i ==3{
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }

	// point := 0
	// for point < 10{
	// 	fmt.Println(point)
	// 	point++
	// }

	// for i := 0; i < 10; i++{
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	if i == 6{
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// arr := [3]int{1,2,3}
	// for i := 0; i < len(arr); i++ { //lenは配列の長さを取得する
	// 	fmt.Println(arr[i])
	// }

	// arr := [3]int{1,2,3}

	// // //rangeはインデックス番号と値を取り出す。
	// // //インデックス番号を消す場合アンダーバー
	// // for _, v := range arr {
	// // 	fmt.Println(v)
	// // }

	// //インデックス番号のみの場合片方のみ記載
	// for v := range arr {
	// 	fmt.Println(v)
	// }

	// //スライス型
	// sl := []string{"python", "PHP", "GO"}
	// for i, v := range sl{
	// 	fmt.Println(i,v)
	// }

	//キーが文字列型、数値がint型となる
	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// //キーがint型、数値が文字列型となる
	// n := map[int]string{100: "momo", 200: "kaki"}

	// for k, v := range n {
	// 	fmt.Println(k, v)
	// }

	//アンダーバーで数値のみ(momo,kaki)
	//1つだけでkeyのみ(100,200)
	n := map[int]string{100: "momo", 200: "kaki"}

	for v := range n {
		fmt.Println(v)
	}
}
