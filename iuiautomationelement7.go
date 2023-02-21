package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement7 struct {
	ole.IUnknown
}

type IUIAutomationElement7Vtbl struct {
	ole.IUnknownVtbl
}

// IID为204e8572-cfc3-4c11-b0c8-7da7420750b7
var IID_IUIAutomationElement7 = &ole.GUID{0xd22108aa, 0x8ac5, 0x49a5, [8]byte{0x83, 0x7b, 0x37, 0xbb, 0xb3, 0xd7,
	0x59, 0x1e}}
