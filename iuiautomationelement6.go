package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement5 struct {
	ole.IUnknown
}

type IUIAutomationElement5Vtbl struct {
	ole.IUnknownVtbl
}

// IID为98141C1D-0D0E-4175-BBE2-6BFF455842A7
var IID_IUIAutomationElement5 = &ole.GUID{0xd22108aa, 0x8ac5, 0x49a5, [8]byte{0x83, 0x7b, 0x37, 0xbb, 0xb3, 0xd7,
	0x59, 0x1e}}
