package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/gonutz/w32/v2"
)

type IUIAutomationElement9 struct {
	IUIAutomationElement8
}

type IUIAutomationElement9Vtbl struct {
	IUIAutomationElement8Vtbl
	Get_CachedIsDialog  uintptr
	Get_CurrentIsDialog uintptr
}

// IID为39325fac-039d-440e-a3a3-5eb81a5cecc3
var IID_IUIAutomationElement9 = ole.NewGUID("39325fac-039d-440e-a3a3-5eb81a5cecc3")

func (ele9 *IUIAutomationElement9) VTable() *IUIAutomationElement9Vtbl {
	return (*IUIAutomationElement9Vtbl)(unsafe.Pointer(ele9.RawVTable))
}

func (ele9 *IUIAutomationElement9) GetCachedIsDialog() (b w32.BOOL, err error) {
	return getCachedIsDialog(ele9)
}

func (ele9 *IUIAutomationElement9) GetCurrentIsDialog() (b w32.BOOL, err error) {
	return getCurrentIsDialog(ele9)
}

func getCurrentIsDialog(ele9 *IUIAutomationElement9) (b w32.BOOL, err error) {
	hr, _, _ := syscall.SyscallN(ele9.VTable().Get_CurrentIsDialog, uintptr(unsafe.Pointer(ele9)),
		uintptr(b))

	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getCachedIsDialog(ele9 *IUIAutomationElement9) (b w32.BOOL, err error) {

	hr, _, _ := syscall.SyscallN(ele9.VTable().Get_CachedIsDialog, uintptr(unsafe.Pointer(ele9)), uintptr(b))

	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
