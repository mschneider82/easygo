package easygo

import "testing"

func TestStringReverse(t *testing.T) {
	if StringReverse("Normal") != "lamroN" {
		t.Errorf("expected lamroN")
	}
}

func TestStringToXXHash(t *testing.T) {
	if StringToXXHash("easy come easy go") != "767e98630b9a5454" {
		t.Errorf("unexpected hash for string")
	}
}

func TestStringStrip(t *testing.T) {
	var input = "bla \n blub \t\r\n xxx"
	out := StringStrip(input, " \t\n\r")
	if out != "blablubxxx" {
		t.Errorf("expected striped string, got: %v", out)
	}
}

func TestStringToInt(t *testing.T) {
	val, err := StringToInt("123")
	if val != 123 || err != nil {
		t.Errorf("expected number")
	}
}
