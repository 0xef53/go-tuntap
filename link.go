package tuntap

import "syscall"

// SetInterfaceUp changes the state of a given interface to UP.
//
// ifName should not exceed 16 bytes.
//
// This is identical to running:  ip link set up dev $ifName
//
// If there is an error, it will be of type *os.SyscallError.
func SetInterfaceUp(ifName string) error {
	var req ifReq
	copy(req.Name[:(syscall.IFNAMSIZ-1)], ifName)
	return ifaceLinkUpIoctl(&req)
}

// SetInterfaceDown changes the state of a given interface to DOWN.
//
// ifName should not exceed 16 bytes.
//
// This is identical to running:  ip link set down dev $ifName
//
// If there is an error, it will be of type *os.SyscallError.
func SetInterfaceDown(ifName string) error {
	var req ifReq
	copy(req.Name[:(syscall.IFNAMSIZ-1)], ifName)
	return ifaceLinkDownIoctl(&req)
}
