package base64

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	data := []string{
		"root", "root$123",
		"mysql", "mysql$123",
		"oracle", "oracle$123",
		"postgres", "postgres$123",
	}

	for _, d := range data {
		encStr := Encode(d)
		fmt.Printf("%v -> %v\n", d, encStr)
	}
}

func TestDecode(t *testing.T) {
	data := []string{
		"cm9vdA==", "cm9vdCQxMjM=",
		"bXlzcWw=", "bXlzcWwkMTIz",
		"b3JhY2xl", "b3JhY2xlJDEyMw==",
		"cG9zdGdyZXM=", "cG9zdGdyZXMkMTIz",
	}

	for _, d := range data {
		decStr, err := Decode(d)
		if err != nil {
			t.Errorf("failed to decode string %v, error: %v", d, err)
			continue
		}

		fmt.Printf("%v -> %v\n", d, decStr)
	}
}
