package iteration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 7)
	expected := "aaaaaaa"

	assert.Equal(t, actual, expected)
}

func ExampleRepeat() {
	repeated := Repeat("x", 5)
	fmt.Println(repeated)
	// Output: xxxxx
}

// To run benchmarks: `go test -bench=.`
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 12)
	}
}
