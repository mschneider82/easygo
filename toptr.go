package easygo

import ()

// BoolToPtr Converts to a Pointer
func BoolToPtr(b bool) *bool {
	return &b
}

// StringToPtr Converts to a Pointer
func StringToPtr(s string) *string {
	return &s
}

// StringsToPtr Converts to a Pointer
func StringsToPtr(s []string) *[]string {
	return &s
}

// Int64ToPtr Converts to a Pointer
func Int64ToPtr(i int64) *int64 {
	return &i
}

// Int32ToPtr Converts to a Pointer
func Int32ToPtr(i int32) *int32 {
	return &i
}

// IntToPtr Converts to a Pointer
func IntToPtr(i int) *int {
	return &i
}

// IntsToPtr Converts to a Pointer
func IntsToPtr(i []int) *[]int {
	return &i
}

// Float32ToPtr Converts to a Pointer
func Float32ToPtr(f float32) *float32 {
	return &f
}

// Float64ToPtr Converts to a Pointer
func Float64ToPtr(f float64) *float64 {
	return &f
}

// ByteToPtr Converts to a Pointer
func ByteToPtr(b byte) *byte {
	return &b
}

// BytesToPtr Converts to a Pointer
func BytesToPtr(b []byte) *[]byte {
	return &b
}
