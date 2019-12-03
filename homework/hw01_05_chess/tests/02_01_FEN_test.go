package tests

import (
	"github.com/xdire/algrefresh/homework/hw01_05_chess"
	"github.com/xdire/algrefresh/util"
	"strings"
	"testing"
)

func Test_FENParse(t *testing.T) {
	files, err := util.GetDirectoryFiles("02_bits/1_Bitboard_FEN")
	if err != nil {
		t.Fatalf("Cannot get test files from: 01_strings_checks directory")
	}
	testReader(t, files, func(t *testing.T, input string, compareTo string) error {
		bitboard := &hw01_05_chess.BitBoard{}
		inputFormatted := strings.Trim(input," \n\r\t")
		compareFormatted := strings.TrimRight(compareTo, "\r\n")
		err := bitboard.ParseFEN(inputFormatted)
		if err != nil {
			t.Error(err)
		}
		output := bitboard.PrintLineNumberPositions()
		if output != compareFormatted {
			t.Errorf("Failed for input of %s with result of \n[%s] should be \n[%s]", inputFormatted, output, compareFormatted)
		}
		return nil
	})
}