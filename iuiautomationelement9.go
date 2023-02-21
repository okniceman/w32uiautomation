package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement9 struct {
	ole.IUnknown
}

type IUIAutomationElement9Vtbl struct {
	ole.IUnknownVtbl
}

// IID为39325fac-039d-440e-a3a3-5eb81a5cecc3
var IID_IUIAutomationElement9 = &ole.GUID{0xd22108aa, 0x8ac5, 0x49a5, [8]byte{0x83, 0x7b, 0x37, 0xbb, 0xb3, 0xd7,
	0x59, 0x1e}}
