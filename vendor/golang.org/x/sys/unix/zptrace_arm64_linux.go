// Code generated by linux/mkall.go generatePtracePair("arm", "arm64"). DO NOT EDIT.

// +build linux
// +build arm64

package unix

import "unsafe"

// PtraceRegsArm is the registers used by arm binaries.
type PtraceRegsArm struct {
	Uregs [18]uint32
}

// PtraceGetRegsArm fetches the registers used by arm binaries.
func PtraceGetRegsArm(pid int, regsout *PtraceRegsArm) error {
	return ptrace(PTRACE_GETREGS, pid, 0, uintptr(unsafe.Pointer(regsout)))
}

// PtraceSetRegsArm sets the registers used by arm binaries.
func PtraceSetRegsArm(pid int, regs *PtraceRegsArm) error {
	return ptrace(PTRACE_SETREGS, pid, 0, uintptr(unsafe.Pointer(regs)))
}

// PtraceRegsArm64 is the registers used by arm64 binaries.
type PtraceRegsArm64 struct {
	Regs   [31]uint64
	Sp     uint64
	Pc     uint64
	Pstate uint64
}

// PtraceGetRegsArm64 fetches the registers used by arm64 binaries.
func PtraceGetRegsArm64(pid int, regsout *PtraceRegsArm64) error {
	return ptrace(PTRACE_GETREGS, pid, 0, uintptr(unsafe.Pointer(regsout)))
}

// PtraceSetRegsArm64 sets the registers used by arm64 binaries.
func PtraceSetRegsArm64(pid int, regs *PtraceRegsArm64) error {
	return ptrace(PTRACE_SETREGS, pid, 0, uintptr(unsafe.Pointer(regs)))
}

// PtraceGetRegSetArm64 fetches the registers used by arm64 binaries.
func PtraceGetRegSetArm64(pid, addr int, regsout *PtraceRegsArm64) error {
	iovec := Iovec{(*byte)(unsafe.Pointer(regsout)), uint64(unsafe.Sizeof(*regsout))}
	return ptrace(PTRACE_GETREGSET, pid, uintptr(addr), uintptr(unsafe.Pointer(&iovec)))
}

// PtraceSetRegSetArm64 sets the registers used by arm64 binaries.
func PtraceSetRegSetArm64(pid, addr int, regs *PtraceRegsArm64) error {
	iovec := Iovec{(*byte)(unsafe.Pointer(regs)), uint64(unsafe.Sizeof(*regs))}
	return ptrace(PTRACE_SETREGSET, pid, uintptr(addr), uintptr(unsafe.Pointer(&iovec)))
}
