package calc

import "testing"

func TestSoma(t *testing.T) {
	if got := Soma(2, 3); got != 5 {
		t.Errorf("Soma(2,3) = %d; want 5", got)
	}
}
