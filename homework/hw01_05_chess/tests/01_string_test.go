package tests

import (
	"github.com/xdire/algrefresh/util"
	"strconv"
	"strings"
	"testing"
)

func Test_StringCounters(t *testing.T) {
	files, err := util.GetDirectoryFiles("01_strings_checks")
	if err != nil {
		t.Fatalf("Cannot get test files from: 01_strings_checks directory")
	}
	testReader(t, files, func(t *testing.T, input string, compareTo string) error {
		length := len(strings.Trim(input," \n\r\t"))
		desired, err := strconv.ParseInt(compareTo, 10, 0)
		if err != nil {
			t.Errorf("Parameter to compare is wrongly defined for " +
				"String Counters, should be integer, while it is: %+v, error: %+v", compareTo, err)
		}
		if int64(length) != desired{
			t.Errorf("Failed for length of %d of input [%s] should be %d", length, input, desired)
		}
		return nil
	})
}