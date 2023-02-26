package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElement2 struct {
	*IUIAutomationElement
}

type IUIAutomationElement2Vtbl struct {
	IUIAutomationElementVtbl
	Get_CachedFlowsFrom                 uintptr
	Get_CachedLiveSetting               uintptr
	Get_CachedOptimizeForVisualContent  uintptr
	Get_CurrentFlowsFrom                uintptr
	Get_CurrentLiveSetting              uintptr
	Get_CurrentOptimizeForVisualContent uintptr
}

// IID为6749c683-f70d-4487-a698-5f79d55290d6
var IID_IUIAutomationElement2 = ole.NewGUID("6749c683-f70d-4487-a698-5f79d55290d6")

func (elem2 *IUIAutomationElement2) VTable() *IUIAutomationElement2Vtbl {
	return (*IUIAutomationElement2Vtbl)(unsafe.Pointer(elem2.RawVTable))
}
