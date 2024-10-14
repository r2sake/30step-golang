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
	fmt.Println("  exercises1 - 問１名前と自己紹介を表示させるプログラム")
}
func runExercises1() {
	fmt.Println("最強エンジニア")
	fmt.Println("つよつよエンジニアです。よろしくお願いします！")
}
