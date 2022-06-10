package main

import (
        "flag"
	"io"
        "os"
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
                
                // TODO: Create expression reading stream
        } else if *inputFile != "" {
                // TODO: Create file reading stream
        } else {
                panic("No input methods are specified. Choose either -e or -f")
        }
        
        if *outputFile != "" {
                // TODO: Create file writing stream
        } else {
                outputStream = os.Stdout
        }

        handler := &lab2.ComputeHandler{
                Input: inputStream.
                Output: outputStream,
        }

        err := handler.Compute()
        if err {
                panic(err)
        }
}
