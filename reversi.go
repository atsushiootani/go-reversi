package main

import (
	"fmt"
)

const FieldWidth = 8
const FieldHeight = 8

var fields [FieldWidth][FieldHeight]int

func main() {
	fmt.Println("aaa")
	printFields(fields)
}

func printFields(fields [FieldWidth][FieldHeight]int) {
	fmt.Print("-")
	for x := 0; x < FieldWidth; x++ {
		for y := 0; y < FieldHeight; y++ {
			fmt.Print(fields[x][y])
		}
	}
}
