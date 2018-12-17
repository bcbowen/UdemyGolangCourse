package word

import (
	"fmt"
	"testing"
)

const testString = "you dipped a chip, you took a bite, you dipped it again"

//TestCount tests word.Count()
func TestCount(t *testing.T) {
	result := Count(testString)
	expected := 12
	if result != expected {
		t.Error("want", expected, "got", result)
	}

}

// ExampleCount shows word.Count
func ExampleCount() {
	fmt.Println(Count(testString))
	// Output:
	// 12

}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(testString)
	}
}

func TestUseCount(t *testing.T) {
	result := UseCount(testString)
	count := result["you"]
	if count != 3 {
		t.Error("want", 3, "got", count)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(testString)
	}
}
