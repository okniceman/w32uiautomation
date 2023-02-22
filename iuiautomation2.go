package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomation2 struct {
	ole.IUnknown
	*IUIAutomation
}

type IUIAutomation2Vtbl struct {
	ole.IUnknownVtbl
	Get_AutoSetFocus       uintptr
	Get_ConnectionTimeout  uintptr
	Get_TransactionTimeout uintptr
	Put_AutoSetFocus       uintptr
	Put_ConnectionTimeout  uintptr
	Put_TransactionTimeout uintptr
}

// IID为34723aff-0c9d-49d0-9896-7ab52df8cd8a
var IID_IUIAutomation2 = ole.NewGUID("34723aff-0c9d-49d0-9896-7ab52df8cd8a")

// CUIAutomation8为e22ad333-b25f-460c-83d0-0581107395c9,win8之后的版本使用此字段
var CLSID_CUIAutomation8 = ole.NewGUID("e22ad333-b25f-460c-83d0-0581107395c9")

// NewUIAutomation2
//
//	@Description: UIAutomation2构造函数
//	@return *IUIAutomation2
//	@return error
func NewUIAutomation2() (*IUIAutomation2, error) {
	result, err := ole.CreateInstance(CLSID_CUIAutomation8, IID_IUIAutomation2)
	if err != nil {
		return nil, err
	}

	auto2 := (*IUIAutomation2)(unsafe.Pointer(result))
	auto2.IUIAutomation, err = NewUIAutomation()

	if err != nil {
		return nil, err
	}
	return auto2, nil
}

func (auto2 *IUIAutomation2) VTable() *IUIAutomation2Vtbl {
	return (*IUIAutomation2Vtbl)(unsafe.Pointer(auto2.RawVTable))
}

// GetAutoSetFocus
//
//	@Description: Specifies whether calls to UI Automation control pattern methods automatically set focus to the target element.
//	@receiver auto2
//	@return b
//	@return err
func (auto2 *IUIAutomation2) GetAutoSetFocus() (b bool, err error) {
	return getAutoSetFocus(auto2)
}

// PutAutoSetFocus
//
//	@Description: Specifies whether calls to UI Automation control pattern methods automatically set focus to the target element.
//	@receiver auto
//	@param b
//	@return err
func (auto2 *IUIAutomation2) PutAutoSetFocus(b bool) (err error) {
	return putAutoSetFocus(auto2, &b)
}

func putAutoSetFocus(auto2 *IUIAutomation2, b *bool) (err error) {
	hr, _, _ := syscall.SyscallN(auto2.VTable().Put_AutoSetFocus, uintptr(unsafe.Pointer(auto2)),
		uintptr(unsafe.Pointer(b)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getAutoSetFocus(auto2 *IUIAutomation2) (b bool, err error) {
	hr, _, _ := syscall.SyscallN(auto2.VTable().Get_AutoSetFocus, uintptr(unsafe.Pointer(auto2)),
		uintptr(unsafe.Pointer(&b)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
