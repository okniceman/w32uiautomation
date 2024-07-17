package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElementArray struct {
	ole.IUnknown
}

type IUIAutomationElementArrayVtbl struct {
	ole.IUnknownVtbl
	Get_Length uintptr
	GetElement uintptr
}

// IID为14314595-b4bc-4055-95f2-58f2e42c9855
var IID_IUIAutomationElementArray = &ole.GUID{0x14314595, 0xb4bc, 0x4055, [8]byte{0x95, 0xf2, 0x58, 0xf2, 0xe4, 0x2c,
	0x98, 0x55}}

func (arr *IUIAutomationElementArray) VTable() *IUIAutomationElementArrayVtbl {
	return (*IUIAutomationElementArrayVtbl)(unsafe.Pointer(arr.RawVTable))
}

func (arr *IUIAutomationElementArray) GetElement(index int32) (found *IUIAutomationElement, err error) {
	return getElement(arr, index)
}

func (arr *IUIAutomationElementArray) Get_Length() (length int32, err error) {
	return getLength(arr)
}

func getLength(arr *IUIAutomationElementArray) (length int32, err error) {
	hr, _, _ := syscall.SyscallN(arr.VTable().Get_Length,
		uintptr(unsafe.Pointer(arr)),
		uintptr(unsafe.Pointer(&length)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getElement(arr *IUIAutomationElementArray, index int32) (found *IUIAutomationElement, err error) {
	hr, _, _ := syscall.SyscallN(arr.VTable().GetElement,
		uintptr(unsafe.Pointer(arr)),
		uintptr(index),
		uintptr(unsafe.Pointer(&found)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
