package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomation5 struct {
	IUIAutomation4
}

type IUIAutomation5Vtbl struct {
	IUIAutomation4Vtbl
	AddNotificationEventHandler    uintptr
	RemoveNotificationEventHandler uintptr
}

// IID为25F700C8-D816-4057-A9DC-3CBDEE77E256
var IID_IUIAutomation5 = ole.NewGUID("25F700C8-D816-4057-A9DC-3CBDEE77E256")

func NewUIAutomation5() (*IUIAutomation5, error) {
	result, err := ole.CreateInstance(CLSID_CUIAutomation8, IID_IUIAutomation5)
	if err != nil {
		return nil, err
	}

	return (*IUIAutomation5)(unsafe.Pointer(result)), nil
}

func (auto5 *IUIAutomation5) VTable() *IUIAutomation5Vtbl {
	return (*IUIAutomation5Vtbl)(unsafe.Pointer(auto5.RawVTable))

}
