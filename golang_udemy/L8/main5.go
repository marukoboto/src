package main

import (
	"fmt"
)

//switch
//型スイッチ

// インターフェース型は全ての引数を格納できる
// ただし、計算などは行うことができない
func anything(a interface{}) {
	// fmt.Println(a)
	//typeで型を判断
	switch v := a.(type){
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func main() {
	anything("aaa")
	anything(1)

	//xにインターフェース型で3が格納されている
	var x interface{} = 3
	//xがint型を確認し、int型でiに代入する（型アサーション）
	i, isInt := x.(int)
	fmt.Println(i, isInt)
	//xは型アサーションを行っていないため、interface型とint型となり、計算が行えない
	// fmt.Println(x + 2)

	//int型なのでfloat型だと復元ができない
	// f := x.(float64)
	// fmt.Println(f)

	//２つ目の返り値を設定するとtrueかfalseかが返ってくる
	//今回の場合　ｘはint型なのにfloat型を指定しているので、エラーが出る
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	// var x interface {} = 3
	//xの形式での分岐式
	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
		//isInt内に入れて、int型にして型アサーション
		//iに数字が入る場合isIntがtrueとなるため、x is Intが出力される。
		//i==int型 &&isint == true の場合i, "x is Int"出力するみたいなイメージ
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't Know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't Know")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}
}
