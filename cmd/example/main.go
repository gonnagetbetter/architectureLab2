package main

import (
	"flag"
	"io"
	"os"
	"strings"

	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression from input")
	inputFromFile   = flag.String("f", "", "Expression from file")
	outputToFile    = flag.String("o", "", "Output results to a file")
)

func main() {
	flag.Parse()

	var input io.Reader = nil
	var output = os.Stdout

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	}

	if *inputFromFile != "" {
		f, err := os.Open(*inputFromFile)
		if err != nil {
			os.Stderr.WriteString("Something went wrong. Please check your command")
		}
		defer f.Close()
		input = f
	}

	if *outputToFile != "" {
		f, err := os.Create(*outputToFile)
		if err != nil {
			os.Stderr.WriteString("Something went wrong. Please check your command")
		}
		defer f.Close()
		output = f
	}

	if input == nil {
		os.Stderr.WriteString("Input can't be empty")
		return
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if err != nil {
		println(err)
	}
}
