package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement3 struct {
	ole.IUnknown
}

type IUIAutomationElement3Vtbl struct {
	ole.IUnknownVtbl
}

// IID为8471DF34-AEE0-4A01-A7DE-7DB9AF12C296
var IID_IUIAutomationElement3 = &ole.GUID{0xd22108aa, 0x8ac5, 0x49a5, [8]byte{0x83, 0x7b, 0x37, 0xbb, 0xb3, 0xd7,
	0x59, 0x1e}}
