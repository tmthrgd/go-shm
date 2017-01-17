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

// Taken from shm_open(3):
// 	shm_open() creates and opens a new, or opens an existing, POSIX shared
// 	memory object. A POSIX shared memory object is in effect a handle which
// 	can be used by unrelated processes to mmap(2) the same region of shared
// 	memory. The shm_unlink() function performs the converse operation,
// 	removing an object previously created by shm_open().
// 	
// 	The operation of shm_open() is analogous to that of open(2). name
// 	specifies the shared memory object to be created or opened. For
// 	portable use, a shared memory object should be identified by a name of
// 	the form /somename; that is, a null-terminated string of up to NAME_MAX
// 	(i.e., 255) characters consisting of an initial slash, followed by one
// 	or more characters, none of which are slashes.
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	fileName := name

	for len(name) != 0 && name[0] == '/' {
		name = name[1:]
	}

	if len(name) == 0 {
		return nil, &os.PathError{Op: "open", Path: fileName, Err: unix.EINVAL}
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
		return nil, &os.PathError{Op: "open", Path: fileName, Err: err}
	}

	return os.NewFile(uintptr(fd), fileName), nil
}

// Taken from shm_unlink(3):
// 	The  operation  of shm_unlink() is analogous to unlink(2): it removes a
// 	shared memory object name, and, once all processes  have  unmapped  the
// 	object, de-allocates and destroys the contents of the associated memory
// 	region.  After a successful shm_unlink(),  attempts  to  shm_open()  an
// 	object  with  the same name will fail (unless O_CREAT was specified, in
// 	which case a new, distinct object is created).
func Unlink(name string) error {
	fileName := name

	for len(name) != 0 && name[0] == '/' {
		name = name[1:]
	}

	if len(name) == 0 {
		return &os.PathError{Op: "unlink", Path: fileName, Err: unix.EINVAL}
	}

	if err := unix.Unlink(devShm + name); err != nil {
		return &os.PathError{Op: "unlink", Path: fileName, Err: err}
	}

	return nil
}
