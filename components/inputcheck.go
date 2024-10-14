package components

import (
	"fmt"
	"os"
	"strconv"
)

func InputCheckForSampleCode(inStr []string, step int) (int, bool) {
	return inputCheck("sample", inStr, step)
}

func InputCheckForExercisesCode(inStr []string, step int) (int, bool) {
	return inputCheck("exercises", inStr, step)
}

func inputCheck(t string, i []string, s int) (int, bool) {
	typeLen := len(t)
	if len(i) < 1 {
		fmt.Println(fmt.Sprintf("Usage: go run step%02d/%s/main.go %s<%s_no>", s, t, t, t))
		return 0, false
	}
	ns := os.Args[1]
	if len(ns) < typeLen+1 || ns[:typeLen] != t {
		fmt.Println(fmt.Sprintf("Invalid %s name: %s\n", t, ns))
		fmt.Println(fmt.Sprintf("Example name should be in the format '%s<number>'", t))
		return 0, false
	}
	no, err := strconv.Atoi(ns[typeLen:])
	if err != nil {
		fmt.Println(fmt.Sprintf("Invalid %s number: %s\n", t, ns[typeLen:]))
		return 0, false
	}
	return no, true
}
