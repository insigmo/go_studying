package io_scanner

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ScanIO() {
	in := bufio.NewScanner(os.Stdin)
	seenVars := make(map[string]bool)

	for in.Scan() {
		txt := in.Text()

		if _, found := seenVars[txt]; found {
			continue
		}
		seenVars[txt] = true
		fmt.Println(txt)
	}
}

func func_1() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		i, _ := strconv.Atoi(in.Text())

		if i < 10 {
			continue
		} else if i > 100 {
			break
		} else {
			fmt.Println(i)
		}
	}
}

func ReadFile() {
	file, err := os.Open("/Users/ruabbb1/git/go_studying/src/1.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
