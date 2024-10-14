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
		case 2:
			runExercises2()
		default:
			fmt.Printf("Unknown exercises: %s\n", os.Args[1])
			availableExercises()
		}
	}
}

func availableExercises() {
	fmt.Println("Available exercises:")
	fmt.Println("  exercises1 - 問１華氏を摂氏に変換するプログラム")
	fmt.Println("  exercises2 - 問２与えられた年が閏年かどうか判定するプログラム")
}
func runExercises1() {
	f := 77
	fmt.Println("f=", f, " c:", (f-32)*5/9)
}

func runExercises2() {
	year := 2024
	r4 := year % 4
	r100 := year % 100
	r400 := year % 400
	if r4 == 0 {
		if r100 != 0 && r400 == 0 {
			fmt.Println("閏年です")
			return
		}
	}
	fmt.Println("閏年ではありません")
}
