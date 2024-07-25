package main

import (
	"brLang/repl"
	"bufio"
	"fmt"
	"os"
	"os/user"
)

func Process(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Bytes())
	}
}

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("welcome to brr language !!", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
