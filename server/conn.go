package server

import "syscall"

type fDconn struct {
	Fd int
}

func (f fDconn) Write(b []byte) (int, error) {
	return syscall.Write(f.Fd, b)
}

func (f fDconn) Read(b []byte) (int, error) {
	return syscall.Read(f.Fd, b)
}

func (f fDconn) Close() error {
	return syscall.Close(f.Fd)
}
