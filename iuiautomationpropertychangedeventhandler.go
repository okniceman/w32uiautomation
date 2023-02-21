package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationPropertyChangedEventHandler struct {
	ole.IUnknown
	ref int32
}

type IUIAutomationPropertyChangedEventHandlerVtbl struct {
	ole.IUnknownVtbl
	HandlePropertyChangedEvent uintptr
}

// IID为40cd37d4-c756-4b0c-8c6f-bddfeeb13b50
var IID_IUIAutomationPropertyChangedEventHandler = &ole.GUID{0x40cd37d4, 0xc756, 0x4b0c, [8]byte{0x8c, 0x6f, 0xbd, 0xdf,
	0xee, 0xb1, 0x3b, 0x50}}

func (pc *IUIAutomationPropertyChangedEventHandler) VTable() *IUIAutomationPropertyChangedEventHandlerVtbl {
	return (*IUIAutomationPropertyChangedEventHandlerVtbl)(unsafe.Pointer(pc.RawVTable))
}

func propertyChangedEventHandler_queryInterface(this *ole.IUnknown, iid *ole.GUID, punk **ole.IUnknown) uint32 {
	*punk = nil
	if ole.IsEqualGUID(iid, ole.IID_IUnknown) {
		*punk = this
	} else if ole.IsEqualGUID(iid, IID_IUIAutomationPropertyChangedEventHandler) {
		*punk = this
	} else {
		return ole.E_NOINTERFACE
	}
	propertyChangedEventHandler_addRef(this)
	return ole.S_OK
}

func propertyChangedEventHandler_addRef(this *ole.IUnknown) int32 {
	pthis := (*IUIAutomationPropertyChangedEventHandler)(unsafe.Pointer(this))
	pthis.ref++
	return pthis.ref
}

func propertyChangedEventHandler_release(this *ole.IUnknown) int32 {
	pthis := (*IUIAutomationPropertyChangedEventHandler)(unsafe.Pointer(this))
	pthis.ref--
	return pthis.ref
}

// NewIUIAutomationPropertyChangedEventHandler
//
//	@Description: 初始化propertyChanged句柄
//	@param handlerFunc 触发后的回调函数
//	@return IUIAutomationFocusChangedEventHandler
func NewIUIAutomationPropertyChangedEventHandler(handlerFunc func(this *IUIAutomationPropertyChangedEventHandler,
	propertyId PROPERTYID, newValue ole.VARIANT) syscall.Handle) IUIAutomationFocusChangedEventHandler {

	lpVtbl := &IUIAutomationPropertyChangedEventHandlerVtbl{
		IUnknownVtbl: ole.IUnknownVtbl{
			QueryInterface: syscall.NewCallback(propertyChangedEventHandler_queryInterface),
			AddRef:         syscall.NewCallback(propertyChangedEventHandler_addRef),
			Release:        syscall.NewCallback(propertyChangedEventHandler_release),
		},
		HandlePropertyChangedEvent: syscall.NewCallback(handlerFunc),
	}

	return IUIAutomationFocusChangedEventHandler{
		IUnknown: ole.IUnknown{RawVTable: (*interface{})(unsafe.Pointer(lpVtbl))},
	}
}
