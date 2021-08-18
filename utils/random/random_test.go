package random

import (
	"fmt"
	"testing"
)

func TestAlphanumeric(t *testing.T) {
	fmt.Println(Alphanumeric(10))
}

func TestAlphabetic(t *testing.T) {
	fmt.Println(Alphabetic(10))
}

func TestNumeric(t *testing.T) {
	fmt.Println(Numeric(10))
}

func TestAscii(t *testing.T) {
	fmt.Println(Ascii(10))
}
