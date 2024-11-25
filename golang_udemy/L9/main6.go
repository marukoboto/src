package main

import "fmt"

//マップ
//箱みたいなもの→名前と数字を一緒にリスト化するイメージ


func main() {
	//map内にキーとバリュー（中身を持つ）
	var m = map[string]int{"A": 100, "B": 200}

	fmt.Println(m)

	//暗黙的宣言
	m2 := map[string] int {"A": 100, "B":200}

	fmt.Println(m2)

	//改行しても行末にコロンをつければ作れる
	m3 := map[int]string {
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	//make関数
	//データ型の指定
	m4 := make(map[int]string)
	fmt.Println(m4)

	//文字列型
	//それぞれ追加される　map[1:JAPAN 2:USA]
	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	//値の取り出し
	fmt.Println(m["A"])
	fmt.Println(m4[2])
	//登録されていない値→初期値（空欄）が出力される
	fmt.Println(m4[3])

	//エラーハンドリング
	//返り値を2つ受け取れるため、エラーハンドリングが行える
	_, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}
	// fmt.Println(s, ok)

	//値の更新
	m4[2] = "US"
	fmt.Println(m4)

	m4[3] = "CHINA"
	fmt.Println(m4)

	//値の削除
	//第一引数削除したいmap 第二引数　キー
	delete(m4,3)
	fmt.Println(m4)

	//len関数で要素数の確認
	fmt.Println(len(m4))

	//floatで小数点以下も行える
	m5 := map[int]float64{
		1: 3.14,
		2: 3.33,
	}
	fmt.Println(m5)

	m6 := make(map[int]float64)
		m6[1]= 3.141592
		m6[2]= 3.3333334
	fmt.Println(m6)
}