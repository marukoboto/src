package main

import "fmt"

//ラベル付きfor

func main() {
	// Loop:
	// 	for {
	// 		for {
	// 			for {
	// 				fmt.Println("START")
	// 				//Loopをつける事で無駄な処理を飛ばすことが出来る
	// 				break Loop
	// 			}
	// 			fmt.Println("処理をしない")
	// 		}
	// 		fmt.Println("処理をしない")
	// 	}
	// 	fmt.Println("END")

	//Jが3以下でぐるぐる
	// for i := 0; i < 3; i++ {
	// 	for j := 1; j < 3; j++ {
	// 		fmt.Println(i, j, i*j)
	// 	}
	// }

Loop:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop //Loopを使う事で外のループまで飛ぶ
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理をしない")
	}
}
