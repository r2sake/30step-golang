package main

import (
	"fmt"
	"github.com/r2sake/30step-golang/components"
	"os"
)

func main() {
	exercisesNo, isOk := components.InputCheckForExercisesCode(os.Args, 4)
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
	fmt.Println("  exercises1 - 問１年齢（整数）、身長（浮動小数点数）、名前（文字列）を表す変数を宣言、値を代入する")
	fmt.Println("  exercises2 - 問２私の名前は[名前]です。年齢は[年齢]歳で、身長は[身長]cmです")
}
func runExercises1() {
	age := 26
	tall := 175.2
	name := "最強エンジニア"
	fmt.Println(tall)
	fmt.Println(age)
	fmt.Println(name)
}

func runExercises2() {
	age := 26
	tall := 175.2
	name := "最強エンジニア"
	fmt.Printf("私の名前は%sです。年齢は%d歳で、身長は%fcmです。", name, age, tall)
}
