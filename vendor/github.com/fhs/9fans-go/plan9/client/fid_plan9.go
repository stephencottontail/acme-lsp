package client

import (
	"os"
	"syscall"
)

type Fid struct {
	*os.File
}

// Close closes the fid.
//
// This wrapper prevents runtime panic if *Fid is nil.
// See https://github.com/golang/go/issues/18617
func (f *Fid) Close() error {
	if f == nil {
		return os.ErrInvalid
	}
	return f.File.Close()
}

func (f *Fid) Write(b []byte) (n int, err error) {
	if len(b) == 0 {
		// Zero-length writes never happen in the standard library
		// but it's useful for writes to /mnt/acme/N/data when
		// we want to delete the text specified in /mnt/acme/N/addr.
		//
		// This is the workaround suggested here:
		// https://codereview.appspot.com/7406046#msg4
		// (commit https://go.googlesource.com/go/+/722ee1f4797a81916b19e80df479058897c44923)
		return syscall.Write(int(f.File.Fd()), b)
	}
	return f.File.Write(b)
}
