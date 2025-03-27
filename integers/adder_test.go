package integers

import (
	"fmt"
	"testing"
)

func numAssert(t testing.TB, a, b int){
	t.Helper()
	if a!=b {
		t.Fatalf("Expected %d, got %d", a, b)
	}
}

func Add(a, b int) int {
	return a+b
}

func TestAdder(t *testing.T){
	t.Run("Sum", func(t *testing.T) {
		numAssert(t, Add(1, 2), 3)
	})
}

func ExampleAdd(){
	sum := Add(1, 3)
	fmt.Println(sum)
	// Output: 4
}