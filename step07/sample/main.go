package main

import (
	"fmt"
	"github.com/r2sake/30step-golang/components"
	"os"
	"strconv"
)

func main() {
	sampleNo, isOk := components.InputCheckForSampleCode(os.Args, 7)
	if !isOk {
		availableSamples()
	} else {
		switch sampleNo {
		case 1:
			runSample1()
		case 2:
			runSample2()
		case 3:
			runSample3()
		case 4:
			runSample4()
		case 5:
			runSample5()
		default:
			fmt.Printf("Unknown sample: %s\n", os.Args[1])
			availableSamples()
		}
	}
}

func availableSamples() {
	fmt.Println("Available samples:")
	fmt.Println("  sample1 - 基本的なif文")
	fmt.Println("  sample2 - if-else文")
	fmt.Println("  sample3 - if-else if-else文")
	fmt.Println("  sample4 - 初期化付きif")
	fmt.Println("  sample5 - エラーハンドリングでのif文")
}

func runSample1() {
	age := 18
	
	if age >= 18 {
		fmt.Println("成人です")
	}
}
func runSample2() {
	age := 16
	
	if age >= 18 {
		fmt.Println("成人です")
	} else {
		fmt.Println("未成年です")
	}
	
	if !(age >= 18) {
		fmt.Println("未成年です")
	}
}
func runSample3() {
	score := 85
	
	if score >= 90 {
		fmt.Println("優秀な成績です")
	} else if score >= 80 {
		fmt.Println("良い成績です")
	} else if score >= 70 {
		fmt.Println("普通の成績です")
	} else {
		fmt.Println("もう少し頑張りましょう")
	}
}

func runSample4() {
	if num := 10; num < 0 {
		fmt.Printf("%dは負の数です\n", num)
	} else if num == 0 {
		fmt.Println("ゼロです")
	} else {
		fmt.Printf("%dは正の数です\n", num)
	}
	// ここでは構文エラーになります。
	// fmt.Printf("%d", num)
}

func runSample5() {
	if value, err := strconv.Atoi("123"); err != nil {
		fmt.Println("変換エラー:", err)
	} else {
		fmt.Printf("変換成功: %d\n", value)
	}
}
