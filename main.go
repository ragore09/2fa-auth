package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("secrets")
	Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		fmt.Printf("%s: %s\n", words[0], GetTOTP(words[1]))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
