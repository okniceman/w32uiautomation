package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement4 struct {
	ole.IUnknown
}

type IUIAutomationElement4Vtbl struct {
	ole.IUnknownVtbl
}

// IID为3B6E233C-52FB-4063-A4C9-77C075C2A06B
var IID_IUIAutomationElement4 = &ole.GUID{0xd22108aa, 0x8ac5, 0x49a5, [8]byte{0x83, 0x7b, 0x37, 0xbb, 0xb3, 0xd7,
	0x59, 0x1e}}
