package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationNotCondition struct {
	ole.IUnknown
}

type IUIAutomationNotConditionVtbl struct {
	ole.IUnknownVtbl
	GetChild uintptr
}

func (and *IUIAutomationNotCondition) VTable() *IUIAutomationNotConditionVtbl {
	return (*IUIAutomationNotConditionVtbl)(unsafe.Pointer(and.RawVTable))
}

// IID为f528b657-847b-498c-8896-d52b565407a1
var IID_IUIAutomationNotCondition = &ole.GUID{0xf528b657, 0x847b, 0x498c, [8]byte{0x88, 0x96, 0xd5, 0x2b, 0x56,
	0x54, 0x07, 0xa1}}
