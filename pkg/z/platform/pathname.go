package platform

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

func (p Pathname) MultiExt() string {
	filename := p.Filename()
	parts := strings.Split(filename.String(), ".")
	l := len(parts)
	switch l {
	case 0:
		return ""
	case 2:
		return "." + parts[1]
	default:
		return "." + parts[l-2] + "." + parts[l-1]
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
