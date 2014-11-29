package tuntap

import (
	"os"
	"syscall"
	"unsafe"
)

func ifaceAddIoctl(req *ifReq, uid, gid int, persist bool) error {
	file, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer file.Close()
	req.Flags |= syscall.IFF_TUN_EXCL
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, file.Fd(), uintptr(syscall.TUNSETIFF), uintptr(unsafe.Pointer(req))); err != 0 {
		return os.NewSyscallError("ioctl: TUNSETIFF", err)
	}
	if uid != -1 {
		if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, file.Fd(), uintptr(syscall.TUNSETOWNER), uintptr(uid)); err != 0 {
			return os.NewSyscallError("ioctl: TUNSETOWNER", err)
		}
	}
	if gid != -1 {
		if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, file.Fd(), uintptr(syscall.TUNSETGROUP), uintptr(gid)); err != 0 {
			return os.NewSyscallError("ioctl: TUNSETGROUP", err)
		}
	}
	if persist {
		if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, file.Fd(), uintptr(syscall.TUNSETPERSIST), 1); err != 0 {
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
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, file.Fd(), uintptr(syscall.TUNSETIFF), uintptr(unsafe.Pointer(req))); err != 0 {
		return os.NewSyscallError("ioctl: TUNSETIFF", err)
	}
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, file.Fd(), uintptr(syscall.TUNSETPERSIST), 0); err != 0 {
		return os.NewSyscallError("ioctl: TUNSETPERSIST", err)
	}
	return nil
}

func ifaceFeaturesIoctl() (uint16, error) {
	var features uint16
	file, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, file.Fd(), uintptr(syscall.TUNGETFEATURES), uintptr(unsafe.Pointer(&features))); err != 0 {
		return 0, os.NewSyscallError("ioctl: TUNGETFEATURES", err)
	}
	return features, nil
}
