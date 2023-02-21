package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationPropertyCondition struct {
	ole.IUnknown
}

type IUIAutomationPropertyConditionVtbl struct {
	ole.IUnknownVtbl
	Get_PropertyConditionFlags uintptr
	Get_PropertyId             uintptr
	Get_PropertyValue          uintptr
}

func (and *IUIAutomationPropertyCondition) VTable() *IUIAutomationPropertyConditionVtbl {
	return (*IUIAutomationPropertyConditionVtbl)(unsafe.Pointer(and.RawVTable))
}

// IID为99ebf2cb-5578-4267-9ad4-afd6ea77e94b
var IID_IUIAutomationPropertyCondition = &ole.GUID{0x99ebf2cb, 0x5578, 0x4267, [8]byte{0x9a, 0xd4, 0xaf, 0xd6, 0xea,
	0x77, 0xe9, 0x4b}}
