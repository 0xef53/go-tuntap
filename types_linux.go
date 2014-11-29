package tuntap

import "syscall"

type ifReq struct {
	Name  [syscall.IFNAMSIZ]byte
	Flags uint16
}
