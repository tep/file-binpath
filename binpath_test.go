package binpath

import (
	"os"
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	want := os.Args[0]
	if i := strings.LastIndexByte(want, '/'); i != -1 {
		want = want[:i]
	}

	if got := Join(); got != want {
		t.Errorf("got %q; wanted %q", got, want)
	}
}
