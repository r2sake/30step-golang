package main

import (
	"fmt"
	"github.com/r2sake/30step-golang/components"
	"os"
)

func main() {
	sampleNo, isOk := components.InputCheckForSampleCode(os.Args, 1)
	if !isOk {
		availableSamples()
	} else {
		switch sampleNo {
		case 1:
			runSample1()
		default:
			fmt.Printf("Unknown sample: %s\n", os.Args[1])
			availableSamples()
		}
	}
}

func availableSamples() {
	fmt.Println("Available samples:")
	fmt.Println("  sample1 - Go言語の簡単なコード例")
}

func runSample1() {
	// 変数宣言と初期化
	message := "Hello, Go!"
	
	// 標準出力に表示
	fmt.Println(message)
	
	// スライス（動的配列）の使用
	numbers := []int{1, 2, 3, 4, 5}
	
	// rangeを使ったイテレーション
	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}
}
