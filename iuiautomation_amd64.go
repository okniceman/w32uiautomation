//go:build amd64
// +build amd64

package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func createPropertyCondition(aut *IUIAutomation, propertyId PROPERTYID, value ole.VARIANT) (*IUIAutomationCondition, error) {
	var newCondition *IUIAutomationCondition
	hr, _, _ := syscall.SyscallN(
		aut.VTable().CreatePropertyCondition,
		uintptr(unsafe.Pointer(aut)),
		uintptr(propertyId),
		uintptr(unsafe.Pointer(&value)),
		uintptr(unsafe.Pointer(&newCondition)))
	if hr != 0 {
		return nil, ole.NewError(hr)
	}
	return newCondition, nil
}
