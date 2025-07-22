package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rd <file-path>")
		return
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file :%s", err)
	}
	defer file.Close()

	read := bufio.NewScanner(file)
	for read.Scan() {
		line := read.Text()
		fmt.Println(line)
	}
	if err := read.Err(); err != nil {
		log.Fatalf("error reading file %s", err)
	}
}
