package iteration

import (
	"strings"
	"testing"
)

func assertString(t testing.TB, a, b string){
	t.Helper()
	if(a!=b){
		t.Errorf("Got %q, want %q", a, b)
	}
}

// func Repeat(a string) string {
// 	ret:=""
// 	for i := 0; i < 5; i++ {
// 		ret+=a
// 	}
// 	return ret
// }

func Repeat(a string, n int) string {
	var ret strings.Builder
	for i := 0; i < n; i++ {
		ret.WriteString(a)
	}
	return ret.String()
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func TestIter(t *testing.T){
	assertString(t, Repeat("a", 5), "aaaaa")
}