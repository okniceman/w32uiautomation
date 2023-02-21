package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationAndCondition struct {
	ole.IUnknown
}

type IUIAutomationAndConditionVtbl struct {
	ole.IUnknownVtbl
	Get_ChildCount           uintptr
	GetChildren              uintptr
	GetChildrenAsNativeArray uintptr
}

func (and *IUIAutomationAndCondition) VTable() *IUIAutomationAndConditionVtbl {
	return (*IUIAutomationAndConditionVtbl)(unsafe.Pointer(and.RawVTable))
}

// IID为a7d0af36-b912-45fe-9855-091ddc174aec
var IID_IUIAutomationAndCondition = &ole.GUID{0xa7d0af36, 0xb912, 0x45fe, [8]byte{0x98, 0x55, 0x09, 0x1d, 0xdc, 0x17,
	0x4a, 0xec}}
