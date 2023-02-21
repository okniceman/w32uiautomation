package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement2 struct {
	ole.IUnknown
}

type IUIAutomationElement2Vtbl struct {
	ole.IUnknownVtbl
}

// IID为6749c683-f70d-4487-a698-5f79d55290d6
var IID_IUIAutomationElement2 = &ole.GUID{0xd22108aa, 0x8ac5, 0x49a5, [8]byte{0x83, 0x7b, 0x37, 0xbb, 0xb3, 0xd7,
	0x59, 0x1e}}
