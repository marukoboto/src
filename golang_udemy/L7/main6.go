package main

import "fmt"

//関数
//ジェネレーターの実装

//integersで呼び出され、funcで返す。
func integers() func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}


func main(){
	ints := integers() //ここ（intsに返ってきた値が保存されている。）
	fmt.Println(ints()) //1
	fmt.Println(ints())	//2
	fmt.Println(ints())	//3

	otherints := integers()
	fmt.Println(otherints()) //1
	fmt.Println(otherints()) //2
	fmt.Println(otherints()) //3

}