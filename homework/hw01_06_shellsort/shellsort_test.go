package hw01_06_shellsort

import (
	"github.com/xdire/algrefresh/sorts"
	"reflect"
	"testing"
)

func Test_ShellSortGap2(t *testing.T) {
	t.Run("ShellSort Half Gap", func(t *testing.T) {
		data := []int64{1, 3, 7, 0, 12, 2, 5}
		sorts.ShellInt(data)
		if !reflect.DeepEqual(data, []int64{0, 1, 2, 3, 5, 7, 12}) {
			t.Error("ShellSort expected result not right")
		}
	})
}