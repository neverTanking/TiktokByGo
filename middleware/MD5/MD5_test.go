package MD5

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	s := "TY12345"
	a := "KJLKKK450"
	b := "ty12345"
	fmt.Println(Encode(s))
	fmt.Println(Encode(a))
	fmt.Println(Encode(b))
}

func TestCheck(t *testing.T) {
	a := "876hjkP"
	b := "87hjm11"
	c := "87hjm11"
	fmt.Println(Check(a, Encode(b)))
	fmt.Println(Check(b, Encode(c)))
}
