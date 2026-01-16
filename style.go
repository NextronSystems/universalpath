package universalpath

import (
	"github.com/NextronSystems/universalpath/unix"
	"github.com/NextronSystems/universalpath/windows"
)

type Style int

const (
	Unix Style = iota
	Windows
)

func (p Style) Clean(s string) string {
	if p == Unix {
		return unix.Clean(s)
	} else {
		return windows.Clean(s)
	}
}

func (p Style) Split(s string) (string, string) {
	if p == Unix {
		return unix.Split(s)
	} else {
		return windows.Split(s)
	}
}

func (p Style) Base(s string) string {
	if p == Unix {
		return unix.Base(s)
	} else {
		return windows.Base(s)
	}
}

func (p Style) Ext(s string) string {
	if p == Unix {
		return unix.Ext(s)
	} else {
		return windows.Ext(s)
	}
}

func (p Style) Dir(s string) string {
	if p == Unix {
		return unix.Dir(s)
	} else {
		return windows.Dir(s)
	}
}

func (p Style) Join(s ...string) string {
	if p == Unix {
		return unix.Join(s...)
	} else {
		return windows.Join(s...)
	}
}

func (p Style) Separator() string {
	if p == Unix {
		return "/"
	} else {
		return `\`
	}
}

func (p Style) SeparatorByte() byte {
	if p == Unix {
		return '/'
	} else {
		return '\\'
	}
}

func (p Style) IsAbs(s string) bool {
	if p == Unix {
		return unix.IsAbs(s)
	} else {
		return windows.IsAbs(s)
	}
}
