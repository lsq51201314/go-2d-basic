package engine

import "unsafe"

const (
	SM_CXSCREEN = uintptr(0)
	SM_CYSCREEN = uintptr(1)
)

const (
	FLOAT_LENGTH = int32(unsafe.Sizeof(float32(0)))
	INT_LENGTH   = int32(unsafe.Sizeof(int32(0)))
)
