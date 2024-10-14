package main

import (
	"fmt"
	"github.com/r2sake/30step-golang/components"
	"log"
	"os"
	"time"
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
	fmt.Println("  exercises1 - 問１年齢計算するプログラム")
	fmt.Println("  exercises2 - 問２簡単な計算機プログラム")
}
func runExercises1() {
	var year, month, day, age int
	var err error
	fmt.Println("生まれた年を入力してください")
	_, err = fmt.Scan(&year)
	if err != nil {
		log.Fatalln("生まれた年の入力にエラーがあります:", err)
	}
	fmt.Println("生まれた月を入力してください")
	_, err = fmt.Scan(&month)
	if err != nil {
		log.Fatalln("生まれた月の入力にエラーがあります:", err)
	}
	fmt.Println("生まれた日を入力してください")
	_, err = fmt.Scan(&day)
	if err != nil {
		log.Fatalln("生まれた日の入力にエラーがあります:", err)
	}
	birthdate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	curYear, curMonth, curDay := time.Now().Local().Date()
	age = curYear - year
	// 誕生日を超えていない場合は-1する
	if curMonth < birthdate.Month() || (curMonth == birthdate.Month() && curDay < birthdate.Day()) {
		age--
	}
	fmt.Println(fmt.Sprintf("あなたは%d歳です", age))
}

func runExercises2() {
	var num1, num2 int
	var res int
	var operator string
	var err error
	
	fmt.Println("数値１を入力してください")
	_, err = fmt.Scan(&num1)
	if err != nil {
		log.Fatalln("数値１の入力にエラーがあります:", err)
	}
	fmt.Println("数値２を入力してください")
	_, err = fmt.Scan(&num2)
	if err != nil {
		log.Fatalln("数値２の入力にエラーがあります:", err)
	}
	fmt.Println("演算子を入力してください(+,-.*,/)")
	_, err = fmt.Scan(&operator)
	if err != nil {
		log.Fatalln("演算子の入力にエラーがあります:", err)
	}
	
	if num2 == 0 && operator == "/" {
		log.Fatalln("ゼロでの除算はできません")
	}
	
	switch operator {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
	default:
		log.Fatalln("演算子の入力が不正です")
	}
	fmt.Println(fmt.Sprintf("%d %s %d = %d です", num1, operator, num2, res))
}
