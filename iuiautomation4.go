package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomation4 struct {
	IUIAutomation3
}

type IUIAutomation4Vtbl struct {
	IUIAutomation3Vtbl
	AddChangesEventHandler    uintptr
	RemoveChangesEventHandler uintptr
}

// IID为1189C02A-05F8-4319-8E21-E817E3DB2860
var IID_IUIAutomation4 = ole.NewGUID("1189C02A-05F8-4319-8E21-E817E3DB2860")

func NewUIAutomation4() (*IUIAutomation4, error) {
	result, err := ole.CreateInstance(CLSID_CUIAutomation8, IID_IUIAutomation4)
	if err != nil {
		return nil, err
	}

	return (*IUIAutomation4)(unsafe.Pointer(result)), nil
}

func (auto4 *IUIAutomation4) VTable() *IUIAutomation4Vtbl {
	return (*IUIAutomation4Vtbl)(unsafe.Pointer(auto4.RawVTable))

}
