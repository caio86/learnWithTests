package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat a string 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"

		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("Do", 2)
	fmt.Println(result)
	// Output: DoDo
}
