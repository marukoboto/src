package main

import (
	"fmt"
	"strconv"
)

//if
//条件分岐
//エラーハンドリング

func main() {
	var s string = "100"

	i, err := strconv.Atoi(s) 
	//stringをintに変換　実行とエラー2つを返す。　エラーが無い場合はnilを返す
	if err != nil{
		fmt.Println(err) 
	} else {
		fmt.Printf("i = %T\n", i)
	}

}