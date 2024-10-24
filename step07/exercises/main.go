package main

import (
	"fmt"
	"github.com/r2sake/30step-golang/components"
	"log"
	"os"
)

func main() {
	exercisesNo, isOk := components.InputCheckForExercisesCode(os.Args, 7)
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
	fmt.Println("  exercises1 - 問１入力された年齢を次の通り分類するプログラム")
	fmt.Println("  exercises2 - 問２簡単な成績判定プログラム")
}
func runExercises1() {
	var age int
	var err error
	fmt.Println("年齢を入力してください")
	_, err = fmt.Scan(&age)
	if err != nil {
		log.Fatalln("年齢の入力にエラーがあります:", err)
	}
	if age < 0 {
		fmt.Println("年齢の値が不正です")
	}
	if age < 13 {
		fmt.Println("子供")
	} else if age < 17 {
		fmt.Println("青少年")
	} else if age < 65 {
		fmt.Println("成人")
	} else {
		fmt.Println("老人")
	}
}

func runExercises2() {
	var score int
	var err error
	fmt.Println("点数を入力してください")
	_, err = fmt.Scan(&score)
	if err != nil {
		log.Fatalln("点数の入力にエラーがあります:", err)
	}
	if score < 0 || score > 100 {
		fmt.Println("点数の値が不正です")
	}
	if score > 89 {
		fmt.Println("A")
	} else if score > 79 {
		fmt.Println("B")
	} else if score > 69 {
		fmt.Println("C")
	} else if score > 59 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}
