package hw01_05_chess

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	// Whites
	WPawn uint8 = iota
	WKnight
	WBishop
	WRook
	WQueen
	WKing
	// Blacks
	BPawn
	BKnight
	BBishop
	BRook
	BQueen
	BKing
)

var resolveMap = map[rune]uint8 {
	// Whites
	'P': WPawn,
	'N': WKnight,
	'B': WBishop,
	'R': WRook,
	'Q': WQueen,
	'K': WKing,
	// Blacks
	'p': BPawn,
	'n': BKnight,
	'b': BBishop,
	'r': BRook,
	'q': BQueen,
	'k': BKing,
}

var resolveOrder = [12]rune{'P', 'N', 'B', 'R', 'Q', 'K', 'p', 'n', 'b', 'r', 'q', 'k'}

type BitBoard struct {
	board [12]uint64
}

//func NewBitBoard() *BitBoard {
//	return &BitBoard{board: [12]uint64{}}
//}

func (bb *BitBoard) DecodeFENCharacter(chr rune) (uint8, error) {
	if val, ok := resolveMap[chr]; ok {
		return val, nil
	}
	return 255, errors.New("not found")
}

func (bb *BitBoard) PlaceFENCharacter(chr rune) (uint8, error) {
	if val, ok := resolveMap[chr]; ok {
		return val, nil
	}
	return 255, errors.New("not found")
}
/**
	For FEN encoding following is true:
		- ROWS are coming top-down, meaning first row is 8, last is 1
		- Every FIRST ROW POSITION is: ROW * 8 (Cells)
 */
func (bb *BitBoard) ParseFEN(str string) error {
	rows := strings.Split(str, "/")
	for row, chars := range rows {
		pos := uint8(0)
		for _, chr := range []rune(chars) {
			numVal, err := strconv.ParseUint(string(chr), 10, 0)
			if err == nil {
				pos += uint8(numVal)
				continue
			}
			figure, err := bb.DecodeFENCharacter(chr)
			if err != nil {
				return fmt.Errorf("wrong FEN format at row %d and character %c at pos %d", row, chr, pos)
			}
			// As the rows coming top-down we need to subtract first row from max row before
			// adding the array positional value
			bb.board[figure] |= 1 << uint64((7 - uint8(row)) * 8 + pos)
			pos++
		}
	}

	return nil
}

func (bb *BitBoard) PrintLineNumberPositions() string  {
	out := ""
	for i, _ := range resolveOrder {
		out += fmt.Sprintf("%d\r\n", bb.board[i])
	}
	return strings.TrimRight(out, "\r\n")
}