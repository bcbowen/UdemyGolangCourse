package dog

import (
	"fmt"
	"testing"
)

// TestYears tests dog.Years
func TestYears(t *testing.T) {
	n := Years(10)
	if n != 70 {
		t.Error("got", n, "want", 70)
	}
}

// TestYearsTwo tests dog.YearsTwo
func TestYearsTwo(t *testing.T) {
	n := YearsTwo(10)
	if n != 70 {
		t.Error("got", n, "want", 70)
	}
}

// ExampleYears gives an example of the expected output from Years
func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

// ExampleYearsTwo gives an example of the expected output from YearsTwo
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

// BenchmarkYears gives an example of the expected output from Years
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(i)
	}
}

// BenchmarkYearsTwo gives an example of the expected output from YearsTwo
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(i)
	}
}
