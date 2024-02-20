package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLineByLine(path string) []string {
	var content []string

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Can't open file")
		return content
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
	}

	return content

}
