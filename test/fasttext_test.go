package main

import (
	"fmt"
	"testing"

	fasttext "github.com/myagmartseren/fasttext_golang"
)

func TestPrediction(t *testing.T) {
	ft := fasttext.NewFasttext()

	if err := ft.LoadModel("./model.bin"); err != nil {
		fmt.Println("error", err)
		return
	}

	result, err := ft.Predict("Last weekend, I played DiceCTF 2024 Quals with my team, Blue Water. We got first place. Huge thanks to my incredible teammates, and to  @dicegangctf for the nice CTF and challenges.")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("result:", result)
}
