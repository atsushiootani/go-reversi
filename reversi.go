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

type Field struct {
	blocks [FieldWidth][FieldHeight]int
}

func (field *Field) reset() {
	for x := 0; x < FieldWidth; x++ {
		for y := 0; y < FieldHeight; y++ {
			field.blocks[x][y] = VOID
		}
	}
	field.blocks[3][3] = BLACK
	field.blocks[4][4] = BLACK
	field.blocks[3][4] = WHITE
	field.blocks[4][3] = WHITE
}

func (field *Field) print() {
	for x := 0; x < FieldWidth; x++ {
		fmt.Println(strings.Repeat("-", FieldWidth*4+1))
		for y := 0; y < FieldHeight; y++ {
			fmt.Print("|")
			mark := "   "
			if field.blocks[x][y] == BLACK {
				mark = " ● "
			} else if field.blocks[x][y] == WHITE {
				mark = " ○ "
			}

			fmt.Print(mark)
		}
		fmt.Println("|")
	}
	fmt.Println(strings.Repeat("-", FieldWidth*4+1))
}

func main() {
	var field Field
	field.reset()
	field.print()
}
