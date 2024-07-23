package main

import (
	"fmt"
	// "github.com/thegangtechnology/go-tutorial/demo"
	"github.com/thegangtechnology/go-tutorial/exercise"
)

func main() {
	// This is the main entry point of your program
	fmt.Println("Hello, World!")
	res, _ := exercise.CalculateTopThree("exercise/wordcount.txt")
	fmt.Println(res)
	fmt.Println("Test edited Go file")
	//
	// demo.hello()
}
