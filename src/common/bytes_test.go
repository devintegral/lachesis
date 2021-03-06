package common

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyBytes(t *testing.T) {
	assertO := assert.New(t)

	data1 := []byte{1, 2, 3, 4}
	exp1 := []byte{1, 2, 3, 4}
	res1 := CopyBytes(data1)
	assertO.EqualValues(res1, exp1)
}

func TestLeftPadBytes(t *testing.T) {
	assertO := assert.New(t)

	val1 := []byte{1, 2, 3, 4}
	exp1 := []byte{0, 0, 0, 0, 1, 2, 3, 4}

	res1 := LeftPadBytes(val1, 8)
	res2 := LeftPadBytes(val1, 2)

	assertO.EqualValues(res1, exp1)
	assertO.EqualValues(res2, val1)
}

func TestRightPadBytes(t *testing.T) {
	assertO := assert.New(t)

	val := []byte{1, 2, 3, 4}
	exp := []byte{1, 2, 3, 4, 0, 0, 0, 0}

	resstd := RightPadBytes(val, 8)
	resshrt := RightPadBytes(val, 2)

	assertO.EqualValues(resstd, exp)
	assertO.EqualValues(resshrt, val)
}

func TestFromHex(t *testing.T) {
	input := "0x01"
	expected := []byte{1}
	result := FromHex(input)
	if !bytes.Equal(expected, result) {
		t.Errorf("Expected %x got %x", expected, result)
	}
}

func TestIsHex(t *testing.T) {
	tests := []struct {
		input string
		ok    bool
	}{
		{"", true},
		{"0", false},
		{"00", true},
		{"a9e67e", true},
		{"A9E67E", true},
		{"0xa9e67e", false},
		{"a9e67e001", false},
		{"0xHELLO_MY_NAME_IS_STEVEN_@#$^&*", false},
	}
	for _, test := range tests {
		if ok := isHex(test.input); ok != test.ok {
			t.Errorf("isHex(%q) = %v, want %v", test.input, ok, test.ok)
		}
	}
}

func TestFromHexOddLength(t *testing.T) {
	input := "0x1"
	expected := []byte{1}
	result := FromHex(input)
	if !bytes.Equal(expected, result) {
		t.Errorf("Expected %x got %x", expected, result)
	}
}

func TestNoPrefixShortHexOddLength(t *testing.T) {
	input := "1"
	expected := []byte{1}
	result := FromHex(input)
	if !bytes.Equal(expected, result) {
		t.Errorf("Expected %x got %x", expected, result)
	}
}

func TestHasHexPrefix(t *testing.T) {
	if hasHexPrefix("") {
		t.Errorf("Empty string should not have hex prefix!")
	}
	if hasHexPrefix("0х") || hasHexPrefix("0Х") {
		t.Errorf("Cyrilic should not be in hex prefix!")
	}
}

func TestHex2BytesFixed(t *testing.T) {
	str := "AABBCCDD";
	if !bytes.Equal([]byte{0xDD}, Hex2BytesFixed(str, 1)) {
		t.Errorf("Expected 0xDD found: %v", Hex2BytesFixed(str, 1))
	}
	if !bytes.Equal([]byte{0xAA, 0xBB, 0xCC, 0xDD}, Hex2BytesFixed(str, 4)) {
		t.Errorf("Expected 0x00AABBCCDD found: %v", Hex2BytesFixed(str, 4))
	}
	if !bytes.Equal([]byte{0, 0xAA, 0xBB, 0xCC, 0xDD}, Hex2BytesFixed(str, 5)) {
		t.Errorf("Expected 0x00AABBCCDD found: %v", Hex2BytesFixed(str, 5))
	}
}

func TestToHex(t *testing.T) {
	// TODO FIXME: order of bytes in hex string is not specified
	// need to expand this test to longer byte array once it's specified.
	if "0xaa" != ToHex([]byte{0xaa}) {
		t.Errorf("TestToHex failed, expected 0xaa found %v",
			ToHex([]byte{0xaa}))
	}
}

func TestToHexArray(t *testing.T) {
	strs := ToHexArray([][]byte{{1, 2, 3}, {4, 5}})
	if strs[0] != "0x010203" {
		t.Errorf("First string expected to be '0x010203' but found '%v'", strs[0])
	}
	if strs[1] != "0x0405" {
		t.Errorf("First string expected to be '0x0405' but found '%v'", strs[1])
	}
}
