package w32uiautomation

import (
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
