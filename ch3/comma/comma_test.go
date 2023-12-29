package comma

import "testing"

func TestComma(t *testing.T) {
	got := comma("12345")
	want := "12,345"

	if got != want {
		t.Errorf("Got %q and Want %q", got, want)
	}
}

func TestBufferComma(t *testing.T) {
	got := buffer_comma([]int{1, 2, 3})
	want := "[1 2 3]"

	if got != want {
		t.Errorf("Got %q and Want %q", got, want)
	}
}
