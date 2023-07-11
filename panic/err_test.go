package panic

import (
	"errors"
	"fmt"
	"testing"
)

func TestErr(t *testing.T) {
	err := errors.New("Not found error")
	fmt.Printf("error: %v", err)
}

func TestProtect(t *testing.T) {
	protect(func() {
		panic("aaa")
	})
}
