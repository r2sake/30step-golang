package main

import (
	"fmt"
	"github.com/r2sake/30step-golang/components"
	"os"
)

func main() {
	exercisesNo, isOk := components.InputCheckForExercisesCode(os.Args, 3)
	if !isOk {
		availableExercises()
	} else {
		switch exercisesNo {
		case 1:
			runExercises1()
		case 2:
			runExercises2()
		case 3:
			runExercises3()
		default:
			fmt.Printf("Unknown exercises: %s\n", os.Args[1])
			availableExercises()
		}
	}
}

func availableExercises() {
	fmt.Println("Available exercises:")
	fmt.Println("  exercises1 - 問１Hello, World!の代わりに、あなたの名前を表示するプログラム")
	fmt.Println("  exercises2 - 問２fmt.Printlnの代わりにfmt.Printfを使用して出力するプログラム")
	fmt.Println("  exercises3 - 問３複数のfmt.Printlnを使用するプログラム")
}
func runExercises1() {
	fmt.Println("最強エンジニア")
}

func runExercises2() {
	fmt.Printf("現在の年: %d", 2024)
}

func runExercises3() {
	fmt.Println("*****")
	fmt.Println("Hello")
	fmt.Println("World")
	fmt.Println("*****")
}
