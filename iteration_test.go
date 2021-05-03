package integers

import (
	"fmt"
	"testing"
)

func Repeat(char string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += char
	}
	return
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeatedChar := Repeat("a")
	fmt.Println(repeatedChar)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
