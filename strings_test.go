package strings_test

import (
	"testing"

	"github.com/jcbwlkr/strings"
)

func TestReverse(t *testing.T) {
	in := "Jacob 助步车"
	want := "车步助 bocaJ"
	got := strings.Reverse(in)

	if got != want {
		t.Errorf("Reverse(%q) = %q; want %q", in, got, want)
	}
}
