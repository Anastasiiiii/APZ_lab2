package main

import (
        "flag"
	"io"
        "os"
	"strings"
        lab2 "github.com/Anastasiiiii/APZ_lab2"
)

var (
        inputExpression = flag.String("e", "", "Expression to compute")
        inputFile = flag.String("f", "", "File with expression to compute")
        outputFile = flag.String("o", "", "Optional. File which holds result of the expresstion")
)

func main() {
        flag.Parse()
        
        var inputStream io.Reader
        var outputStream io.Writer

        if *inputExpression != "" {
                if *inputFile != "" {
                        panic("Both input methods are specified. Choose either -e or -f")
                }
                
		inputStream = strings.NewReader(*inputExpression)
        } else if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			panic(err)
		}
		
		defer file.Close()
		inputStream = file
        } else {
                panic("No input methods are specified. Choose either -e or -f")
        }
        
        if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			panic(err)
		}

		defer file.Close()
		outputStream = file
        } else {
                outputStream = os.Stdout
        }

        handler := &lab2.ComputeHandler{
                Input: inputStream.
                Output: outputStream,
        }

        err := handler.Compute()
        if err != nil {
                panic(err)
        }
}
