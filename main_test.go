package main

import "testing"

func TestPort(t *testing.T) {
	if port != 7718 {
		t.Errorf("Port was incorrect, got: %d, want: %d.", port, 7718)
	}
}

func TestBeats(t *testing.T) {
	if beats != 0 {
		t.Errorf("Beats was incorrect, got: %d, want: %d.", beats, 0)
	}
}
