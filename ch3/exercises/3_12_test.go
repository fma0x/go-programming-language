// Escreva uma função que informe se duas strings são
// anagramas uma da outra.

package exercises_test

import (
	"testing"

	"github.com/fma0x/ch3/exercises"
)

func TestIsAnagram(t *testing.T) {
	got := exercises.AreAnagram("silent", "listen")
	want := true

	if got != true {
		t.Errorf("Expect %t but got %t", want, got)
	}
}
