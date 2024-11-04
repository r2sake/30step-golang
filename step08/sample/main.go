package main

import (
	"fmt"
	"github.com/r2sake/30step-golang/components"
	"os"
)

func main() {
	sampleNo, isOk := components.InputCheckForSampleCode(os.Args, 8)
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
		case 6:
			runSample6()
		default:
			fmt.Printf("Unknown sample: %s\n", os.Args[1])
			availableSamples()
		}
	}
}

func availableSamples() {
	fmt.Println("Available samples:")
	fmt.Println("  sample1 - 論理積 AND演算子（&&）")
	fmt.Println("  sample2 - 論理和 OR演算子（||）")
	fmt.Println("  sample3 - 否定 NOT演算子")
	fmt.Println("  sample4 - 比較演算子との組み合わせ")
	fmt.Println("  sample5 - 複数条件の組み合わせ")
	fmt.Println("  sample6 - 短絡評価（Short-circuit evaluation）")
}

func runSample1() {
	age := 25
	income := 300000
	
	if age >= 18 && income >= 250000 {
		fmt.Println("クレジットカードの申し込み資格があります")
	} else {
		fmt.Println("申し込み条件を満たしていません")
	}
}
func runSample2() {
	isStudent := true
	isElderly := false
	
	if isStudent || isElderly {
		fmt.Println("割引が適用されます")
	} else {
		fmt.Println("通常料金です")
	}
}
func runSample3() {
	isClosed := false
	
	if !isClosed {
		fmt.Println("お店は営業中です")
	} else {
		fmt.Println("お店は閉店しています")
	}
}

func runSample4() {
	score := 85
	
	if score >= 80 {
		fmt.Println("優秀な成績です")
	} else if score >= 60 {
		fmt.Println("合格です")
	} else {
		fmt.Println("不合格です")
	}
}

func runSample5() {
	age := 20
	isStudent := true
	hasMembership := false
	
	if (age < 18 || isStudent) && !hasMembership {
		fmt.Println("学生割引が適用されます")
	} else if hasMembership {
		fmt.Println("会員割引が適用されます")
	} else {
		fmt.Println("通常料金です")
	}
}

func runSample6() {
	age := 15
	income := 300000
	// 比較が行われないというのをわかりやすくするために、checkIncomeという変数型関数を定義。
	checkIncome := func(i int) bool {
		fmt.Println("incomeをチェックしています")
		return i >= 200000
	}
	// ageが18未満なので、checkIncomeは実行されない
	if age >= 18 && checkIncome(income) {
		fmt.Println("条件を満たしています")
	} else {
		fmt.Println("条件を満たしていません")
	}
}
