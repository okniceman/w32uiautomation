package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomation5 struct {
	ole.IUnknown
	*IUIAutomation4
}

type IUIAutomation5Vtbl struct {
	ole.IUnknownVtbl
}

// IID为25F700C8-D816-4057-A9DC-3CBDEE77E256
var IID_IUIAutomation5 = ole.NewGUID("25F700C8-D816-4057-A9DC-3CBDEE77E256")

func NewUIAutomation5() (*IUIAutomation5, error) {
	result, err := ole.CreateInstance(CLSID_CUIAutomation8, IID_IUIAutomation5)
	if err != nil {
		return nil, err
	}

	auto5 := (*IUIAutomation5)(unsafe.Pointer(result))

	auto5.IUIAutomation4, err = NewUIAutomation4()
	if err != nil {
		return nil, err
	}
	return auto5, nil
}

func (auto5 *IUIAutomation5) VTable() *IUIAutomation5Vtbl {
	return (*IUIAutomation5Vtbl)(unsafe.Pointer(auto5.RawVTable))
}
