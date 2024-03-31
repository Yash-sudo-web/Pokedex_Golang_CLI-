package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		print("PokeDex>>")
		scanner.Scan()
		input := scanner.Text()
		fmt.Println(input)
	}
}

func textClean(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
