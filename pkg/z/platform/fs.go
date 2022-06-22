package platform

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type (
	Pathname  string
	UnixPerms uint32
)

var (
	PathSeparator = string(os.PathSeparator)
)

func (p Pathname) String() string {
	return string(p)
}

func (p Pathname) Normalize() Pathname {
	out := p.String()
	if os.PathSeparator == '\\' && os.PathListSeparator == ';' {
		out = strings.ReplaceAll(out, "/", PathSeparator)
	} else { // assume a Unix platform
		out = strings.ReplaceAll(out, "\\", PathSeparator)
	}
	return Pathname(out)
}

func (p Pathname) IsFilename() bool {
	return !strings.ContainsRune(p.String(), os.PathSeparator)
}

func (p Pathname) IsRelPath() bool {
	return !p.IsFilename() && !p.IsAbsPath()
}

func (p Pathname) IsAbsPath() bool {
	if os.PathSeparator == '\\' && os.PathListSeparator == ';' {
		match, _ := regexp.MatchString("^[0-9a-zA-Z\\s_-]+:", p.String())
		return match
	} else {
		var first rune
		for _, c := range p.String() {
			first = c
			break
		}
		return first == os.PathSeparator
	}
}

func (p Pathname) Dirname() Pathname {
	dirname := filepath.Dir(p.String())
	return Pathname(dirname)
}

func (p Pathname) Filename() Pathname {
	filename := filepath.Base(p.String())
	return Pathname(filename)
}

func (p Pathname) IsCanonical() bool {
	return true
}

func (p Pathname) FileExists() bool {
	_, err := os.Stat(p.String())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (p Pathname) DirExists() bool {
	info, err := os.Stat(p.String())
	if err != nil {
		return false
	} else {
		return info.IsDir()
	}
}

func (p Pathname) Ext() string {
	return filepath.Ext(p.String())
}

func (p Pathname) ComposePath(parts ...string) Pathname {
	var v string
	path := strings.TrimRight(p.String(), PathSeparator) + PathSeparator
	for i := 0; i < len(parts); i++ {
		if parts[i] == PathSeparator {
			path += PathSeparator
			continue
		}
		v = strings.TrimRight(parts[i], PathSeparator)
		v = strings.TrimLeft(v, PathSeparator)
		if i < len(parts)-1 {
			v += PathSeparator
		}
		path += v
	}
	return Pathname(path)
}

type FileDescriptor struct {
	*os.File
}

type FilesystemInterface interface {
	OpenFile(p Pathname, flags int, mode os.FileMode) (FileDescriptor, error)
	Write(fd *os.File, data []byte) error
	WriteAt(fd *os.File, off int64, data []byte) error
	CloseFile(FileDescriptor)
	CreateZeroFile(p Pathname, size int64, perm UnixPerms) (*os.File, error)
	//LockFile(*FileDescriptor) error
	//UnlockFile(*FileDescriptor)
	CreateDir(p Pathname, perms UnixPerms) error
	DirectoryExists(Pathname) bool
	RmDir(Pathname) error
	AvailableDiskSpace(Pathname) uint64
}
