package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomation2 struct {
	ole.IUnknown
	auto *IUIAutomation
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
var IID_IUIAutomation2 = &ole.GUID{0x34723aff, 0x0c9d, 0x49d0, [8]byte{0x98, 0x96, 0x7a, 0xb5, 0x2d, 0xf8, 0xcd, 0x8a}}

func (auto *IUIAutomation2) VTable() *IUIAutomation2Vtbl {
	return (*IUIAutomation2Vtbl)(unsafe.Pointer(auto.RawVTable))
}

// NewUIAutomation2
//
//	@Description: UIAutomation2构造函数
//	@return *IUIAutomation2
//	@return error
func NewUIAutomation2() (*IUIAutomation2, error) {
	result, err := ole.CreateInstance(CLSID_CUIAutomation, IID_IUIAutomation2)
	if err != nil {
		return nil, err
	}

	auto2 := (*IUIAutomation2)(unsafe.Pointer(result))
	auto2.auto, err = NewUIAutomation()

	if err != nil {
		return nil, err
	}

	return auto2, nil
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
