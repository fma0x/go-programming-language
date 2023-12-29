package basename

import "testing"

func TestBasename(t *testing.T) {
	got := basename("a/b/c.go")
	want := "c"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestNewBasename(t *testing.T) {
	got := newbasename("a/f/c.go")
	want := "c"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
