package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElement4 struct {
	IUIAutomationElement3
}

type IUIAutomationElement4Vtbl struct {
	IUIAutomationElement3Vtbl
	Get_CachedAnnotationObjects  uintptr
	Get_CachedAnnotationTypes    uintptr
	Get_CachedLevel              uintptr
	Get_CachedPositionInSet      uintptr
	Get_CachedSizeOfSet          uintptr
	Get_CurrentAnnotationObjects uintptr
	Get_CurrentAnnotationTypes   uintptr
	Get_CurrentLevel             uintptr
	Get_CurrentPositionInSet     uintptr
	Get_CurrentSizeOfSet         uintptr
}

// IID为3B6E233C-52FB-4063-A4C9-77C075C2A06B
var IID_IUIAutomationElement4 = ole.NewGUID("3B6E233C-52FB-4063-A4C9-77C075C2A06B")

func (elem4 *IUIAutomationElement4) VTable() *IUIAutomationElement4Vtbl {
	return (*IUIAutomationElement4Vtbl)(unsafe.Pointer(elem4.RawVTable))
}
