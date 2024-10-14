package main

import (
	"fmt"
	"github.com/r2sake/30step-golang/components"
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
		default:
			fmt.Printf("Unknown sample: %s\n", os.Args[1])
			availableSamples()
		}
	}
}

func availableSamples() {
	fmt.Println("Available samples:")
	fmt.Println("  sample1 - プログラムファイルを作成")
}

func runSample1() {
	fmt.Println("Hello, Go!")
}
