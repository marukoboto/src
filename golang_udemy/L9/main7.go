package main

import "fmt"

//マップ
//for

func main () {
	m := map[string]int {
		"Apple": 100,
		"Banana": 200,
	}

	//アンダーバーでキーだけ、アンダーバーだけも行える。
	for _,v:= range m {
		fmt.Println(v)
	}
}