package main

import "fmt"

//配列型

func main(){
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A","B"}
	fmt.Println(arr2)

	arr3 := [3]int{1,2,3}
	fmt.Println(arr3)

	//[...] ドットで要素の数を自動で記述してくれる。
	arr4 := [...]string{"C","D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	//arr2の指定の部分の取り出し
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	//要素は2つしか無いため、2を選ぶと空文字が選択される。
	fmt.Println(arr2[2])
	//引き算も行える。　今回の場合２に空文字　-1なので1の値であるBが表示される
	fmt.Println(arr2[2-1])

	//arr2内の値　2にはセットされていないのでそこにセットする。
	arr2[2] = "C"
	fmt.Println(arr2)

	//要素数が３の中に４つの要素を入れるとエラーに
	//中の要素が同じintでも要素数が異なると別物として考えられる。
	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	// 要素数を調べる関数
	fmt.Println(len(arr1))
}