package main

import "fmt"

//switch
//式スイッチ

//データの型を合わせる
func main() {
	// n := 5
	// switch n{
	// case 1,2:
	// 	fmt.Println("1 or 2")
	// case 3,4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don't Know")
	// }

	//スイッチ文内のみ参照可能
	// switch n := 2; n{
	// case 1,2:
	// 	fmt.Println("1 or 2")
	// case 3,4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don't Know")
	// }

	n := 6
	switch {
	case n > 0 && n < 4:
		fmt.Println("0 < n < 4")
	case n > 3 && n < 7:
		fmt.Println("3 < n < 7")
	}

	//列挙型と演算子の混在はエラーとなる
	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	case n > 3 && n < 6:
		fmt.Println("3 < n < 6")
	default:
		fmt.Println("I don't Know")
	}
}
