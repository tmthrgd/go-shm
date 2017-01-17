// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

// Package shm provides functions to open and unlink shared memory.
package shm

import (
	"os"

	"golang.org/x/sys/unix"
)

const devShm = "/dev/shm/"

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	fileName := name

	for len(name) != 0 && name[0] == '/' {
		name = name[1:]
	}

	if len(name) == 0 {
		return nil, unix.EINVAL
	}

	o := uint32(perm.Perm())
	if perm&os.ModeSetuid != 0 {
		o |= unix.S_ISUID
	}
	if perm&os.ModeSetgid != 0 {
		o |= unix.S_ISGID
	}
	if perm&os.ModeSticky != 0 {
		o |= unix.S_ISVTX
	}

	fd, err := unix.Open(devShm+name, flag|unix.O_CLOEXEC, o)
	if err != nil {
		return nil, err
	}

	return os.NewFile(uintptr(fd), fileName), nil
}

func Unlink(name string) error {
	for len(name) != 0 && name[0] == '/' {
		name = name[1:]
	}

	if len(name) == 0 {
		return unix.EINVAL
	}

	return unix.Unlink(devShm + name)
}
