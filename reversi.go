package main

import (
	"errors"
	"fmt"
	"strings"
)

const FieldWidth = 8
const FieldHeight = 8

const OUT = -1
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
	for y := 0; y < FieldHeight; y++ {
		fmt.Println(strings.Repeat("-", FieldWidth*4+1))
		for x := 0; x < FieldWidth; x++ {
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

func (field *Field) put(turn int, x, y int) error {
	if !field.isValidPosition(x, y) {
		return errors.New("指定の座標は盤外です")
	}
	if !field.isVoid(x, y) {
		return errors.New("すでに石が置かれています")
	}
	field.blocks[x][y] = turn
	return nil
}

func (field *Field) get(x, y int) int {
	if !field.isValidPosition(x, y) {
		return OUT
	}
	return field.blocks[x][y]
}

// x, yの位置が空いているかどうか
func (field *Field) isVoid(x, y int) bool {
	return field.get(x, y) == VOID
}

func (field *Field) isBlack(x, y int) bool {
	return field.get(x, y) == BLACK
}

func (field *Field) isWhite(x, y int) bool {
	return field.get(x, y) == WHITE
}

// x, yが有効な座標か
func (field *Field) isValidPosition(x, y int) bool {
	return x >= 0 && x < FieldWidth && y >= 0 && y < FieldHeight
}

func main() {
	var field Field
	field.reset()
	field.print()

	turn := BLACK
	for {
		x, y, err := fetchCommandFromInput(&field, turn)
		if err != nil {
			fmt.Println("残念！エラーです")
			continue
		}

		fmt.Printf("%d, %d におきます\n", x+1, y+1)
		field.put(turn, x, y)
		field.print()

		turn = BLACK + WHITE - turn
	}
}

// 入力からコマンドを取得する
func fetchCommandFromInput(field *Field, turn int) (x int, y int, err error) {
	for {
		x, y, err = fetchFromInput(turn)
		switch field.get(x, y) {
		case VOID:
			return
		case BLACK, WHITE:
			fmt.Println("すでに石が置かれています")
		case OUT:
			fmt.Println("1 ~ 8 で指定してください")
		}
	}
}

func fetchFromInput(turn int) (x int, y int, err error) {
	player := "黒"
	if turn == WHITE {
		player = "白"
	}

	fmt.Printf("%sの手盤です。コマンドを入力してください(ex. `1 1`): ", player)
	fmt.Scan(&x)
	fmt.Scan(&y)

	// 入力は 1~8 だが、内部的には 0 ~ 7
	x--
	y--
	return
}
