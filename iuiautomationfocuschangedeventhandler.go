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

// IID为c270f6b5-5c69-4290-9745-7a7f97169468
var IID_IUIAutomationFocusChangedEventHandler = &ole.GUID{0xc270f6b5, 0x5c69, 0x4290, [8]byte{0x97, 0x45, 0x7a, 0x7f, 0x97,
	0x16, 0x94, 0x68}}

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
