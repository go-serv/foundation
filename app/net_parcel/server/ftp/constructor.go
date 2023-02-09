package ftp

import (
	"fmt"
	"github.com/mesh-master/foundation/pkg/z/platform"
	"time"
)

func NewUploadProfile(rootDir platform.Pathname, maxSize int64, perms platform.UnixPerms) *uploadProfile {
	prof := new(uploadProfile)
	prof.rootDir = rootDir
	prof.maxFileSize = maxSize
	prof.filePerms = perms
	prof.rootDirPostfixFn = func() platform.Pathname {
		now := time.Now()
		prefix := platform.Pathname("").ComposePath(
			fmt.Sprintf("%d", now.Year()),
			fmt.Sprintf("%.2d", now.Month()),
			fmt.Sprintf("%.2d", now.Day()),
			fmt.Sprintf("%.2d", now.Hour()),
			platform.PathSeparator,
		)
		return prefix
	}
	return prof
}

func WithRootDirPostfix(prof *uploadProfile, fn rootDirPostfixFn) *uploadProfile {
	prof.rootDirPostfixFn = fn
	return prof
}
