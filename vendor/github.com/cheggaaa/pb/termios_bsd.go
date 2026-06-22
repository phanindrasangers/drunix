//go:build (darwin || freebsd || netbsd || openbsd || dragonfly) && !appengine
// +build darwin freebsd netbsd openbsd dragonfly
// +build !appengine

package pb

import "syscall"

const (
	ioctlReadTermios  = syscall.TIOCGETA
	ioctlWriteTermios = syscall.TIOCSETA
)
