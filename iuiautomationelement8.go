package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement8 struct {
	ole.IUnknown
}

type IUIAutomationElement8Vtbl struct {
	ole.IUnknownVtbl
}

// IID为8C60217D-5411-4CDE-BCC0-1CEDA223830C
var IID_IUIAutomationElement8 = &ole.GUID{0xd22108aa, 0x8ac5, 0x49a5, [8]byte{0x83, 0x7b, 0x37, 0xbb, 0xb3, 0xd7,
	0x59, 0x1e}}
