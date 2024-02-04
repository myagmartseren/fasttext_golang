package main

import (
	fasttext "fasttext"
	"fmt"
	"testing"
)

func TestPrediction(t *testing.T) {
	ft := fasttext.NewFasttext()

	ft.LoadModel("./model.bin")

	result := ft.Predict("Some input text")
	fmt.Println("result:", result)
}
