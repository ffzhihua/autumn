package slice

import (
	"autumn/tools/typeconv"
	"fmt"
	"testing"
)

func TestInSlice(t *testing.T) {
	sl := []string{"A", "b"}
	fmt.Println(InSlice("a", sl))
	if !InSlice("A", sl) {
		t.Error("should be true")
	}
	if InSlice("B", sl) {
		t.Error("should be false")
	}

	arga := make([]interface{}, len(sl))
	for i := range sl {
		arga[i] = sl[i]
	}
	sll := ArrToInterface(sl)
	fmt.Println(SliceRand(sll))
	sint := []int64{1, 2, 3, 5, 6}
	fmt.Println(SliceSum(sint))

	title := "123"
	res, err := typeconv.StrToInt(title)
	fmt.Println(res, err)
}
