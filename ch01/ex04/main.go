package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countsFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countsFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countsFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, countsFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, countsFiles map[string][]string) {
	input := bufio.NewScanner(f)
	fileName := f.Name()
	for input.Scan() {
		inputText := input.Text()
		counts[inputText] += 1
		if !isValueInArray(countsFiles[inputText], fileName) {
			countsFiles[inputText] = append(countsFiles[inputText], fileName)
		}
	}
	//input.Err()からのエラーの可能性は継続して無視している・・・
}

func isValueInArray(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}
