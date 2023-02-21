package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationBoolCondition struct {
	ole.IUnknown
}

type IUIAutomationBoolConditionVtbl struct {
	ole.IUnknownVtbl
	Get_BooleanValue uintptr
}

func (and *IUIAutomationBoolCondition) VTable() *IUIAutomationBoolConditionVtbl {
	return (*IUIAutomationBoolConditionVtbl)(unsafe.Pointer(and.RawVTable))
}

// IID为1b4e1f2e-75eb-4d0b-8952-5a69988e2307
var IID_IUIAutomationBoolCondition = &ole.GUID{0x1b4e1f2e, 0x75eb, 0x4d0b, [8]byte{0x89, 0x52, 0x5a, 0x69, 0x98,
	0x8e, 0x23, 0x07}}

