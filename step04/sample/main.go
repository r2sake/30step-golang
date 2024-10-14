package main

import (
	"fmt"
	"github.com/r2sake/30step-golang/components"
	"os"
)

func main() {
	sampleNo, isOk := components.InputCheckForSampleCode(os.Args, 6)
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
		default:
			fmt.Printf("Unknown sample: %s\n", os.Args[1])
			availableSamples()
		}
	}
}

func availableSamples() {
	fmt.Println("Available samples:")
	fmt.Println("  sample1 - 変数の宣言方法")
	fmt.Println("  sample2 - 基本的なデータ型")
	fmt.Println("  sample3 - Go言語の特徴：ゼロ値")
}

func runSample1() {
	// 1. varキーワードを使用
	var age int
	age = 30
	
	// 2. 短縮変数宣言
	name := "Gopher"
	
	// 3. 型を明示的に指定
	var height float64 = 175.5
	
	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
	fmt.Println("Height:", height)
}

func runSample2() {
	var integerNum int = 42
	var floatNum float64 = 3.14
	var complexNum complex128 = 1 + 2i
	var isTrue bool = true
	var message string = "Hello, Go!"
	
	fmt.Printf("Integer: %d\n", integerNum)
	fmt.Printf("Float: %.2f\n", floatNum)
	fmt.Printf("Complex: %v\n", complexNum)
	fmt.Printf("Boolean: %t\n", isTrue)
	fmt.Printf("String: %s\n", message)
}

func runSample3() {
	var i int
	var f float64
	var b bool
	var s string
	
	fmt.Printf("int: %v\n", i)
	fmt.Printf("float64: %v\n", f)
	fmt.Printf("bool: %v\n", b)
	fmt.Printf("string: %q\n", s)
}
