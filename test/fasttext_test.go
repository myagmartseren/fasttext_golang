package main

import (
	fasttext "fasttext"
	"fmt"
	"testing"
)

func TestPrediction(t *testing.T) {
	ft := fasttext.NewFasttext()

	if err := ft.LoadModel("./model.bin"); err != nil {
		fmt.Println("error", err)
		return
	}

	result, err := ft.Predict("Some input text")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("result:", result)
}
