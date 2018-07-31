package easygo

// PtrToString converts *string to string, if nil a empty string is returned
func PtrToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// PtrToStrings converts *[]string to []string, if nil a empty string slice is returned
func PtrToStrings(s *[]string) []string {
	if s != nil {
		return *s
	}
	var x []string
	return x
}

// PtrToInt converts *int to int, if nil 0 is returned
func PtrToInt(i *int) int {
	if i != nil {
		return *i
	}
	return 0
}

// PtrToInts converts *int to int, if nil 0 is returned
func PtrToInts(i *[]int) []int {
	if i != nil {
		return *i
	}
	return []int{}
}

// PtrToInt32 converts *int32 to int32, if nil 0 is returned
func PtrToInt32(i *int32) int32 {
	if i != nil {
		return *i
	}
	return 0
}

// PtrToInt64 converts *int64 to int64, if nil 0 is returned
func PtrToInt64(i *int64) int64 {
	if i != nil {
		return *i
	}
	return 0
}

// PtrToBool converts *bool to bool, if nil false is returned
func PtrToBool(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

// PtrToFloat64 converts *float64 to float64, if nil 0 is returned
func PtrToFloat64(f *float64) float64 {
	if f != nil {
		return *f
	}
	return float64(0)
}

// PtrToByte converts *byte to byte, if nil a empty byte is returned
func PtrToByte(b *byte) byte {
	if b != nil {
		return *b
	}
	var x byte
	return x
}

// PtrToBytes converts *[]byte to []byte, if nil a empty []byte is returned
func PtrToBytes(b *[]byte) []byte {
	if b != nil {
		return *b
	}
	var x []byte
	return x
}
