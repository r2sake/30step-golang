package main

import (
	"fmt"
	"github.com/r2sake/30step-golang/components"
	"os"
)

func main() {
	exercisesNo, isOk := components.InputCheckForExercisesCode(os.Args, 6)
	if !isOk {
		availableExercises()
	} else {
		switch exercisesNo {
		case 1:
			runExercises1()
		default:
			fmt.Printf("Unknown exercises: %s\n", os.Args[1])
			availableExercises()
		}
	}
}

func availableExercises() {
	fmt.Println("Available exercises:")
	fmt.Println("  exercises1 - サンプルコードを実際に動かしてみよう")
}
func runExercises1() {
	// 変数宣言と初期化
	message := "Hello, Go!"
	
	// 標準出力に表示
	fmt.Println(message)
	
	// スライス（動的配列）の使用
	numbers := []int{2, 4, 6, 8, 10, 12}
	
	// rangeを使ったイテレーション
	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}
}
