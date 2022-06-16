package ftp

import "github.com/go-serv/service/pkg/z/platform"

func NewUploadProfile(basedir platform.Pathname, maxSize int64, perms uint32) *uploadProfile {
	prof := new(uploadProfile)
	prof.baseDir = basedir
	prof.maxFileSize = maxSize
	prof.filePerms = perms
	return prof
}
