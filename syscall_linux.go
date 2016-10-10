package tuntap

import (
	"os"
	"syscall"
	"unsafe"
)

func ioctl(a1, a2, a3 uintptr) error {
	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, a1, a2, a3); errno != 0 {
		return errno
	}

	return nil
}

func ifaceAddIoctl(req *ifReq, uid, gid int, persist bool) error {
	file, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	req.Flags |= syscall.IFF_TUN_EXCL

	if err := ioctl(file.Fd(), uintptr(syscall.TUNSETIFF), uintptr(unsafe.Pointer(req))); err != nil {
		return os.NewSyscallError("ioctl: TUNSETIFF", err)
	}

	if uid != -1 {
		if err := ioctl(file.Fd(), uintptr(syscall.TUNSETOWNER), uintptr(uid)); err != nil {
			return os.NewSyscallError("ioctl: TUNSETOWNER", err)
		}
	}

	if gid != -1 {
		if err := ioctl(file.Fd(), uintptr(syscall.TUNSETGROUP), uintptr(gid)); err != nil {
			return os.NewSyscallError("ioctl: TUNSETGROUP", err)
		}
	}

	if persist {
		if err := ioctl(file.Fd(), uintptr(syscall.TUNSETPERSIST), 1); err != nil {
			return os.NewSyscallError("ioctl: TUNSETPERSIST", err)
		}
	}

	return nil
}

func ifaceDelIoctl(req *ifReq) error {
	file, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := ioctl(file.Fd(), uintptr(syscall.TUNSETIFF), uintptr(unsafe.Pointer(req))); err != nil {
		return os.NewSyscallError("ioctl: TUNSETIFF", err)
	}

	if err := ioctl(file.Fd(), uintptr(syscall.TUNSETPERSIST), 0); err != nil {
		return os.NewSyscallError("ioctl: TUNSETPERSIST", err)
	}

	return nil
}

func ifaceFeaturesIoctl() (uint16, error) {
	file, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var features uint16

	if err := ioctl(file.Fd(), uintptr(syscall.TUNGETFEATURES), uintptr(unsafe.Pointer(&features))); err != nil {
		return 0, os.NewSyscallError("ioctl: TUNGETFEATURES", err)
	}

	return features, nil
}
