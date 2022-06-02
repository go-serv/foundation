package platform

const ArchType64 = 32 << (^uint(0) >> 63)

var _ = (ArchType64 == 64)

// Process ID
type ProcId uintptr

type PlatformInterface interface {
}

type mustImplementPlatformInterface struct {
	PlatformInterface
}

type PlatformType struct {
	mustImplementPlatformInterface
}
