package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationCondition struct {
	ole.IUnknown
}

type IUIAutomationConditionVtbl struct {
	ole.IUnknownVtbl
}

func (v *IUIAutomationCondition) VTable() *IUIAutomationConditionVtbl {
	return (*IUIAutomationConditionVtbl)(unsafe.Pointer(v.RawVTable))
}

// IID为352ffba8-0973-437c-a61f-f64cafd81df9
var IID_IUIAutomationCondition = &ole.GUID{0x352ffba8, 0x0973, 0x437c, [8]byte{0xa6, 0x1f, 0xf6, 0x4c, 0xaf, 0xd8, 0x1d, 0xf9}}
