package main

import (
	"fmt"
	"github.com/bxcodec/go-clean-arch/article"
	_ "github.com/bxcodec/go-clean-arch/docs"
	"os"
)

func main() {
	// Prepare Repository


	// Build service Layer
	svc := article.NewService()
	dir, _:= os.Getwd()

	err := svc.Merge(dir + "/pdf/input/file1.pdf", dir + "/pdf/input/file2.pdf")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Merge Successful.")
	}

	err = svc.Split(dir + "/pdf/input/2_pageinput.pdf", 1)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Split Successful.")
	}

	err = svc.Compressed(dir + "/pdf/input/file2.pdf")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Compressed Successful.")
	}
}
