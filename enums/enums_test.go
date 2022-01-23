package enums

import "testing"

func TestSrc(t *testing.T) {
	for _, valid := range validSrc {
		if !valid.IsValid() {
			t.Fatalf("Seemingly valid source is invalid")
		}
	}
	invalid := Src("uniswap")
	if invalid.IsValid() {
		t.Fatalf("Seemingly invalid source is valid")
	}
}
