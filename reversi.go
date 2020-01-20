package main

import (
	"fmt"
	"strings"
)

const FieldWidth = 8
const FieldHeight = 8

const VOID = 0
const BLACK = 1
const WHITE = 2

var fields [FieldWidth][FieldHeight]int

func main() {
	resetFields(&fields)
	printFields(&fields)
}

func resetFields(fields *[FieldWidth][FieldHeight]int) {
	for x := 0; x < FieldWidth; x++ {
		for y := 0; y < FieldHeight; y++ {
			fields[x][y] = VOID
		}
	}
	fields[3][3] = BLACK
	fields[4][4] = BLACK
	fields[3][4] = WHITE
	fields[4][3] = WHITE
}

func printFields(fields *[FieldWidth][FieldHeight]int) {
	for x := 0; x < FieldWidth; x++ {
		fmt.Println(strings.Repeat("-", FieldWidth*4+1))
		for y := 0; y < FieldHeight; y++ {
			fmt.Print("|")
			mark := "   "
			if fields[x][y] == BLACK {
				mark = " ● "
			} else if fields[x][y] == WHITE {
				mark = " ○ "
			}

			fmt.Print(mark)
		}
		fmt.Println("|")
	}
	fmt.Println(strings.Repeat("-", FieldWidth*4+1))
}
