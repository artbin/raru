// +build !linux

package raru

import "syscall"

func Setuid(uid int) (err error) {
	return syscall.Setuid(uid)
}

func Setgid(gid int) (err error) {
	return syscall.Setgid(gid)
}
