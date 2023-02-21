package w32uiautomation

import "github.com/go-ole/go-ole"

type IUIAutomation3 struct {
	ole.IUnknown
	auto2 *IUIAutomation2
}

type IUIAutomation3Vtbl struct {
	ole.IUnknownVtbl
	AddTextEditTextChangedEventHandler    uintptr
	RemoveTextEditTextChangedEventHandler uintptr
}

// IID为73D768DA-9B51-4B89-936E-C209290973E7
var IID_IUIAutomation3 = &ole.GUID{0x73D768DA, 0x9B51, 0x4B89, [8]byte{0x93, 0x6E, 0xC2, 0x09, 0x29, 0x09, 0x73, 0xE7}}
