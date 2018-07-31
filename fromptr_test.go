package easygo

import "testing"

func TestPtrToString(t *testing.T) {
	s := "some string"
	if PtrToString(&s) != "some string" {
		t.Errorf("expected a string")
	}
}

func TestPtrToStrings(t *testing.T) {
    s := []string{}"some string"}
	if PtrToString(&s) != "some string" {
		t.Errorf("expected a string")
	}
}

func TestPtrToInt(t *testing.T) {

}

func TestPtrToInts(t *testing.T) {

}

func TestPtrToInt32(t *testing.T) {

}

func TestPtrToInt64(t *testing.T) {

}

func TestPtrToBool(t *testing.T) {

}

func TestPtrToFloat64(t *testing.T) {

}

func TestPtrToByte(t *testing.T) {

}

func TestPtrToBytes(t *testing.T) {

}
