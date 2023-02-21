package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement6 struct {
	ole.IUnknown
}

type IUIAutomationElement6Vtbl struct {
	ole.IUnknownVtbl
}

// IID为4780d450-8bca-4977-afa5-a4a517f555e3
var IID_IUIAutomationElement6 = &ole.GUID{0xd22108aa, 0x8ac5, 0x49a5, [8]byte{0x83, 0x7b, 0x37, 0xbb, 0xb3, 0xd7,
	0x59, 0x1e}}
