// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

// +build !linux

// Package shm provides functions to open and unlink shared memory.
package shm

/*
#cgo LDFLAGS: -lrt

#include <stdlib.h>          // For free
#include <sys/mman.h>        // For shm_*
*/
import "C"

import (
	"os"
	"unsafe"
)

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
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	fd, err := C.shm_open(nameC, C.int(flag), C.mode_t(perm))
	if err != nil {
		return nil, err
	}

	return os.NewFile(uintptr(fd), name), nil
}

// Taken from shm_unlink(3):
// 	The  operation  of shm_unlink() is analogous to unlink(2): it removes a
// 	shared memory object name, and, once all processes  have  unmapped  the
// 	object, de-allocates and destroys the contents of the associated memory
// 	region.  After a successful shm_unlink(),  attempts  to  shm_open()  an
// 	object  with  the same name will fail (unless O_CREAT was specified, in
// 	which case a new, distinct object is created).
func Unlink(name string) error {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	_, err := C.shm_unlink(nameC)
	return err
}
