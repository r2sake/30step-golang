package main

import (
	"fmt"
	"github.com/r2sake/30step-golang/components"
	"log"
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
	fmt.Println("  sample1 - 標準出力")
	fmt.Println("  sample2 - フォーマット済み出力")
	fmt.Println("  sample3 - 標準入力")
	fmt.Println("  sample4 - エラーハンドリング")
}

func runSample1() {
	name := "Alice"
	age := 25
	fmt.Println("Name:", name, "Age:", age)
	
}
func runSample2() {
	name := "Bob"
	age := 30
	height := 175.5
	
	fmt.Printf("Name: %s, Age: %d, Height: %.1f cm\n", name, age, height)
	
	pi := 3.14159
	fmt.Printf("Pi is approximately %.2f\n", pi)
}
func runSample3() {
	var name string
	var age int
	
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}

func runSample4() {
	var age int
	
	fmt.Print("Enter your age: ")
	_, err := fmt.Scan(&age)
	if err != nil {
		log.Fatal("Error reading age:", err)
	}
	
	fmt.Printf("You are %d years old.\n", age)
}
