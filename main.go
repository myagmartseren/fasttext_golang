package main

import (
	fasttext "fasttext/wrapper"
	"fmt"
)

func main() {
	// Create FastText instance
	ft := fasttext.NewFasttext()

	// Load pre-trained model
	ft.LoadModel("./model.bin")

	// Perform prediction
	result := ft.Predict("Some input text")
	fmt.Println("result:", result)
	var input string
	for {
		fmt.Print("Input:")
		fmt.Scan(&input)
		fmt.Println(ft.Predict(input))
	}

	// // Add your assertion based on the expected result
	// expectedResult := "Expected result"
	// if result != expectedResult {
	// 	fmt.Errorf("Expected %s, but got %s", expectedResult, result)
	// }

	// Destroy resources (not implemented in the provided C++ code)
	// ft.Destroy()
}
