package lbs_des

import (
	"fmt"
	"testing"
)

func TestDesLbs_StrEnc(t *testing.T) {
	des := NewDesLbs()

	result := des.StrEnc("sdfsdfsdfs", "aaa", "bbb", "ccc")

	fmt.Println(result)
}
