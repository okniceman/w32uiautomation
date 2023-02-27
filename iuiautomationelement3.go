package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElement3 struct {
	IUIAutomationElement2
}

type IUIAutomationElement3Vtbl struct {
	IUIAutomationElement2Vtbl
	Get_CachedIsPeripheral  uintptr
	Get_CurrentIsPeripheral uintptr
	ShowContextMenu         uintptr
}

// IID为8471DF34-AEE0-4A01-A7DE-7DB9AF12C296
var IID_IUIAutomationElement3 = ole.NewGUID("8471DF34-AEE0-4A01-A7DE-7DB9AF12C296")

func (elem3 *IUIAutomationElement3) VTable() *IUIAutomationElement3Vtbl {
	return (*IUIAutomationElement3Vtbl)(unsafe.Pointer(elem3.RawVTable))
}
