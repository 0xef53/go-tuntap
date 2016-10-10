package tuntap

import "syscall"

// AddTapInterface creates and configures a new tap interface with the given name and IFF flags.
//
// ifName should not exceed 16 bytes.
//
// If persist is true, tap interface will be configured as persistent.
//
// If uid or gid is not -1 then ioctl TUNSETOWNER or TUNSETGROUP will be called respectively.
//
// This is identical to running:  ip tuntap add mode tap dev $ifName
//
// If there is an error, it will be of type *os.SyscallError.
func AddTapInterface(ifName string, uid, gid int, flags uint16, persist bool) error {
	var req ifReq
	copy(req.Name[:(syscall.IFNAMSIZ-1)], ifName)
	req.Flags = syscall.IFF_TAP | flags
	return ifaceAddIoctl(&req, uid, gid, persist)
}

// DelTapInterface destroys an existing tap interface.
//
// This is identical to running:  ip tuntap del mode tap dev $ifName
//
// If there is an error, it will be of type *os.SyscallError.
func DelTapInterface(ifName string) error {
	var req ifReq
	copy(req.Name[:(syscall.IFNAMSIZ-1)], ifName)
	req.Flags = syscall.IFF_TAP
	return ifaceDelIoctl(&req)
}

// GetFeatures returns all valid IFF flags supported by your kernel.
//
// If there is an error, it will be of type *os.SyscallError.
//
// TUNGETFEATURES was introduced in kernel version 2.6.27.
func GetFeatures() (uint16, error) {
	features, err := ifaceFeaturesIoctl()
	if err != nil {
		return 0, err
	}
	return features, nil
}
