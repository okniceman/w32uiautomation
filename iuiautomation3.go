package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomation3 struct {
	*IUIAutomation2
}

type IUIAutomation3Vtbl struct {
	*IUIAutomation2Vtbl
	AddTextEditTextChangedEventHandler    uintptr
	RemoveTextEditTextChangedEventHandler uintptr
}

// IID为73D768DA-9B51-4B89-936E-C209290973E7
var IID_IUIAutomation3 = ole.NewGUID("73D768DA-9B51-4B89-936E-C209290973E7")

func NewUIAutomation3() (*IUIAutomation3, error) {
	result, err := ole.CreateInstance(CLSID_CUIAutomation8, IID_IUIAutomation3)
	if err != nil {
		return nil, err
	}

	auto3 := (*IUIAutomation3)(unsafe.Pointer(result))
	auto3.IUIAutomation2, err = NewUIAutomation2()
	if err != nil {
		return nil, err
	}

	return auto3, nil
}

func (auto3 *IUIAutomation3) VTable() *IUIAutomation3Vtbl {
	// return (*IUIAutomation3Vtbl)(unsafe.Pointer(auto3.RawVTable))
	return nil
}
