package strsim

import (
	"fmt"
	"testing"
)

func Test_EditDistance_Compare(t *testing.T) {
	e := &edit{}
	v := e.Compare("love", "lolpe")
	fmt.Printf("v = %f\n", v)
}
