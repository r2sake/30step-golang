package main

import (
	"fmt"
	"github.com/r2sake/30step-golang/components"
	"log"
	"os"
	"time"
	"unicode"
)

func main() {
	exercisesNo, isOk := components.InputCheckForExercisesCode(os.Args, 8)
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
	fmt.Println("  exercises1 - 問１映画館の料金計算プログラム")
	fmt.Println("  exercises2 - 問２パスワードの強度チェックプログラム")
}
func runExercises1() {
	var age int
	var isAcademy, hasCard bool
	var discountRate float64
	var err error
	discountRate = 1.0
	base := 1800.0
	price := 0.0
	fmt.Println("年齢を入力してください")
	_, err = fmt.Scan(&age)
	if err != nil {
		log.Fatalln("年齢の入力にエラーがあります:", err)
	}
	fmt.Println("学生ですか？")
	_, err = fmt.Scan(&isAcademy)
	if err != nil {
		log.Fatalln("学生判定の入力にエラーがあります:", err)
	}
	fmt.Println("メンバーカードを持っていますか？")
	_, err = fmt.Scan(&hasCard)
	if err != nil {
		log.Fatalln("メンバーカード判定の入力にエラーがあります:", err)
	}
	
	w := time.Now().Weekday()
	if age < 13 {
		discountRate = 0.5
	} else if age >= 65 {
		discountRate = 0.7
	} else if isAcademy && age <= 25 {
		discountRate = 0.8
	}
	if hasCard {
		discountRate -= 0.05
	}
	price = base * discountRate
	if !(w == time.Saturday || w == time.Sunday) {
		price -= 100
	}
	fmt.Println(fmt.Sprintf("本日の映画料金は、%f円です", price))
}

func runExercises2() {
	var pw string
	var err error
	fmt.Println("パスワードを入力してください")
	_, err = fmt.Scan(&pw)
	if err != nil {
		log.Fatalln("PWの入力にエラーがあります:", err)
	}
	
	if len(pw) < 8 {
		fmt.Println("パスワードは8文字以上入力してください")
	}
	hasUpper := false
	for _, r := range pw {
		if unicode.IsUpper(r) {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		fmt.Println("パスワードは少なくとも1つの大文字を含めてください")
	}
	
	hasLower := false
	for _, r := range pw {
		if unicode.IsLower(r) {
			hasLower = true
			break
		}
	}
	if !hasLower {
		fmt.Println("パスワードは少なくとも1つの小文字を含めてください")
	}
	
	hasNumber := false
	for _, r := range pw {
		if '0' <= r && r <= '9' {
			hasNumber = true
			break
		}
	}
	if !hasNumber {
		fmt.Println("パスワードは少なくとも1つの数字を含めてください")
	}
	
	hasSpecial := false
	specialChar := "!@#$%^&*"
	for _, r := range pw {
		for _, s := range specialChar {
			if r == s {
				hasSpecial = true
				break
			}
		}
	}
	if !hasSpecial {
		fmt.Println("パスワードは少なくとも1つの特殊文字を含めてください")
	}
}
