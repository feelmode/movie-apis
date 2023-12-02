package rating

import (
	"testing"
)

func TestRound(t *testing.T) {
	result := Round(7.55)
	expected := float32(7.6)
	if expected != result {
		t.Errorf("result was incorrect, got: %v, want: %v.", result, expected)
	}
}
func TestRoundNotUp(t *testing.T) {
	result := Round(7.54)
	expected := float32(7.5)
	if expected != result {
		t.Errorf("result was incorrect, got: %v, want: %v.", result, expected)
	}
}
