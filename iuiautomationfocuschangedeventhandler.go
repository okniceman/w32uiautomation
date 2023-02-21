package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationFocusChangedEventHandler struct {
	ole.IUnknown
	ref int32
}

type IUIAutomationFocusChangedEventHandlerVtbl struct {
	ole.IUnknownVtbl
	HandleFocusChangedEvent uintptr
}

var IID_IUIAutomationFocusChangedEventHandler = &ole.GUID{}

func (fceh *IUIAutomationFocusChangedEventHandler) VTable() *IUIAutomationFocusChangedEventHandlerVtbl {
	return (*IUIAutomationFocusChangedEventHandlerVtbl)(unsafe.Pointer(fceh.RawVTable))
}

func focusChangedEventHandler_addRef(this *ole.IUnknown) int32 {
	pthis := (*IUIAutomationFocusChangedEventHandler)(unsafe.Pointer(this))
	pthis.ref++
	return pthis.ref
}

func focusChangedEventHandler_release(this *ole.IUnknown) int32 {
	pthis := (*IUIAutomationFocusChangedEventHandler)(unsafe.Pointer(this))
	pthis.ref--
	return pthis.ref
}

func focusChangedEventHandler_queryInterface(this *ole.IUnknown, iid *ole.GUID, punk **ole.IUnknown) uint32 {
	*punk = nil
	if ole.IsEqualGUID(iid, ole.IID_IUnknown) {
		*punk = this
	} else if ole.IsEqualGUID(iid, IID_IUIAutomationFocusChangedEventHandler) {
		*punk = this
	} else {
		return ole.E_NOINTERFACE
	}

	// addRef
	focusChangedEventHandler_addRef(this)
	return ole.S_OK
}

func NewIUIAutomationFocusChangedEventHandler(handlerFunc func(this *IUIAutomationFocusChangedEventHandler,
	sendor *IUIAutomationElement) syscall.Handle) IUIAutomationFocusChangedEventHandler {

	lpVtbl := &IUIAutomationFocusChangedEventHandlerVtbl{
		IUnknownVtbl: ole.IUnknownVtbl{
			QueryInterface: syscall.NewCallback(focusChangedEventHandler_queryInterface),
			AddRef:         syscall.NewCallback(focusChangedEventHandler_addRef),
			Release:        syscall.NewCallback(focusChangedEventHandler_release),
		},
		HandleFocusChangedEvent: syscall.NewCallback(handlerFunc),
	}

	return IUIAutomationFocusChangedEventHandler{
		IUnknown: ole.IUnknown{RawVTable: (*interface{})(unsafe.Pointer(lpVtbl))},
	}
}
