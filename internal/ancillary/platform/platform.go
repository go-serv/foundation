package platform

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
