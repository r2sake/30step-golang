package components

import (
	"fmt"
	"os"
	"strconv"
)

func InputCheckForSampleCode(inStr []string, step int) (int, bool) {
	if len(inStr) < 1 {
		fmt.Println(fmt.Sprintf("Usage: go run step%02d/sample/main.go sample<sample_no>", step))
		return 0, false
	}
	sampleName := os.Args[1]
	if len(sampleName) < 7 || sampleName[:6] != "sample" {
		fmt.Printf("Invalid sample name: %s\n", sampleName)
		fmt.Println("Example name should be in the format 'sample<number>'")
		return 0, false
	}
	sampleNo, err := strconv.Atoi(sampleName[6:])
	if err != nil {
		fmt.Printf("Invalid sample number: %s\n", sampleName[6:])
		return 0, false
	}
	return sampleNo, true
}

func InputCheckForExercisesCode(inStr []string, step int) (int, bool) {
	if len(inStr) < 1 {
		fmt.Println(fmt.Sprintf("Usage: go run step%02d/exercises/main.go exercises<exercises_no>", step))
		return 0, false
	}
	exercisesName := os.Args[1]
	if len(exercisesName) < 10 || exercisesName[:9] != "exercises" {
		fmt.Printf("Invalid exercises name: %s\n", exercisesName)
		fmt.Println("Example name should be in the format 'exercises<number>'")
		return 0, false
	}
	exercisesNo, err := strconv.Atoi(exercisesName[9:])
	if err != nil {
		fmt.Printf("Invalid exercises number: %s\n", exercisesName[9:])
		return 0, false
	}
	return exercisesNo, true
}
