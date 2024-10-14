package main

import (
	"fmt"
	"github.com/r2sake/30step-golang/components"
	"os"
)

func main() {
	sampleNo, isOk := components.InputCheckForSampleCode(os.Args, 5)
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
		default:
			fmt.Printf("Unknown sample: %s\n", os.Args[1])
			availableSamples()
		}
	}
}

func availableSamples() {
	fmt.Println("Available samples:")
	fmt.Println("  sample1 - 算術演算子")
	fmt.Println("  sample2 - 比較演算子")
	fmt.Println("  sample3 - 論理演算子")
	fmt.Println("  sample4 - 代入演算子")
}

func runSample1() {
	a := 10
	b := 3
	
	fmt.Println("加算:", a+b)
	fmt.Println("減算:", a-b)
	fmt.Println("乗算:", a*b)
	fmt.Println("除算:", a/b)
	fmt.Println("剰余:", a%b)
}

func runSample2() {
	x := 5
	y := 10
	
	fmt.Println("x == y:", x == y)
	fmt.Println("x != y:", x != y)
	fmt.Println("x < y:", x < y)
	fmt.Println("x > y:", x > y)
	fmt.Println("x <= y:", x <= y)
	fmt.Println("x >= y:", x >= y)
}
func runSample3() {
	a := true
	b := false
	
	fmt.Println("a && b:", a && b)
	fmt.Println("a || b:", a || b)
	fmt.Println("!a:", !a)
}
func runSample4() {
	x := 10
	
	x += 5
	fmt.Println("x += 5:", x)
	
	x -= 3
	fmt.Println("x -= 3:", x)
	
	x *= 2
	fmt.Println("x *= 2:", x)
	
	x /= 4
	fmt.Println("x /= 4:", x)
	
	x %= 3
	fmt.Println("x %= 3:", x)
}
