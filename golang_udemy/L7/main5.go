package main

import "fmt"

//関数
//クロージャーの実装
//変数を覚えて利用する仕組み

//Laterが呼び出されたとき、string型で返す
func Later() func(string) string{
	//storeの定義（空文字）
	var store string
	//nextに引数を代入、store内の文字を返す
	//store内に値が保存される
	return func (next string) string {
		s := store// 現在のstoreの値を変数sに保存
        store = next // storeにnextの値を代入
        return s  // 以前のstoreの値（つまり、nextの前の値）を返す
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("My"))
}