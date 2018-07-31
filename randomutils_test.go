package easygo

import "testing"

func TestRandomIDwithPrefix(t *testing.T) {
	if len(RandomIDwithPrefix("bla", 10)) != 13 {
		t.Errorf("unexpected length of RandomIDwithPrefix")
	}
}
