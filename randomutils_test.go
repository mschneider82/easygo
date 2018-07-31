package easygo

import "testing"

func TestRandomIDwithPrefix(t *testing.T) {
	// Test1 with random result
	if len(RandomIDwithPrefix("bla", 10)) != 13 {
		t.Errorf("unexpected length of RandomIDwithPrefix")
	}
	// Test 2 with static result because we only have one Letter to choice
	RandomLetters = []rune("a")
	if RandomIDwithPrefix("bla", 10) != "blaaaaaaaaaaa" {
		t.Errorf("unexpected, should be blaaaaaaaaaaa, got: %v", RandomIDwithPrefix("bla", 10))
	}
}
