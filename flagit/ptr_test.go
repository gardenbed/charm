package flagit

import "time"

func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func intPtr(i int) *int {
	return &i
}

func int8Ptr(i int8) *int8 {
	return &i
}

func int16Ptr(i int16) *int16 {
	return &i
}

func int32Ptr(i int32) *int32 {
	return &i
}

func int64Ptr(i int64) *int64 {
	return &i
}

func uintPtr(u uint) *uint {
	return &u
}

func uint8Ptr(u uint8) *uint8 {
	return &u
}

func uint16Ptr(u uint16) *uint16 {
	return &u
}

func uint32Ptr(u uint32) *uint32 {
	return &u
}

func uint64Ptr(u uint64) *uint64 {
	return &u
}

func float32Ptr(f float32) *float32 {
	return &f
}

func float64Ptr(f float64) *float64 {
	return &f
}

func bytePtr(b byte) *byte {
	return &b
}

func runePtr(r rune) *rune {
	return &r
}

func durationPtr(d time.Duration) *time.Duration {
	return &d
}
