package printints

import "testing"

func TestIntsToString(t *testing.T) {
	got := intsToString([]int{1, 2, 3})
	want := "[1, 2, 3]"

	if got != want {
		t.Errorf("Expect %q but got %q", want, got)
	}
}
