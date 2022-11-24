package app

import (
	"fmt"
	"reflect"
	"testing"
)

type normalizeIdTest struct {
	arg, expected string
}

var normalizeIdTests = []normalizeIdTest{
	{"ac%3Ae401ae6a54335205f612e6ba6de9bc2a:1634056979:v1", "ac:e401ae6a54335205f612e6ba6de9bc2a"},
	{"ac%3a2d042299c53b12d825e896e4c8a08921:1645994832:v1", "ac:2d042299c53b12d825e896e4c8a08921"},
	{"ac%3a1edc95876661c2a9e8d8104eaa8bcc81:1643488273:v1", "ac:1edc95876661c2a9e8d8104eaa8bcc81"},
	{"ac%3Ac664609abbeb679e4e3d188fbaf7a30f:1642710733:v1", "ac:c664609abbeb679e4e3d188fbaf7a30f"},
	{"6e44715e-3235-4fec-904c-8985d25b3f0b:1651666557:v1", "6e44715e-3235-4fec-904c-8985d25b3f0b"},
	{"285f534e-fbfa-4453-b735-ce0b4730d4c8:1651122514:v1", "285f534e-fbfa-4453-b735-ce0b4730d4c8"},
	{"ac%3A360d1b796f0fa80bb5a749545fa4471e:1651579756:v1", "ac:360d1b796f0fa80bb5a749545fa4471e"},
}

func TestNormalizeIdType(t *testing.T){
	for _, test := range normalizeIdTests{
		if output := normalizeId(test.arg); reflect.TypeOf(output).Kind() != reflect.TypeOf(test.expected).Kind() {
			t.Errorf("Output type %q not equal to expected type %q", output, test.expected)
		}
	}
}

func TestNormalizeIdValue(t *testing.T){
	for _, test := range normalizeIdTests{
		if output := normalizeId(test.arg); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkNormalizeId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		normalizeId("ac%3Ae401ae6a54335205f612e6ba6de9bc2a:1634056979:v1")
	}
}

func ExampleNormalizedId() {
	fmt.Println(normalizeId("ac%3Ae401ae6a54335205f612e6ba6de9bc2a:1634056979:v1"))
	//Output: ac:e401ae6a54335205f612e6ba6de9bc2a
}