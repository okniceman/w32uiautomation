package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/gonutz/w32/v2"
)

type IUIAutomationElement2 struct {
	IUIAutomationElement
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

func (elem2 *IUIAutomationElement2) GetCachedFlowsFrom() (elemArray *IUIAutomationElementArray, err error) {
	return getCachedFlowsFrom(elem2)
}

func (elem2 *IUIAutomationElement2) GetCachedLiveSetting() {

}

func getCachedLiveSetting(elem2 *IUIAutomationElement2) {

}

func (elem2 *IUIAutomationElement2) GetCachedOptimizeForVisualContent() (b w32.BOOL, err error) {
	return 0, nil
}

func getCachedOptimizeForVisualContent(elem2 *IUIAutomationElement2) (b w32.BOOL, err error) {
	return 0, nil
}

func getCachedFlowsFrom(elem2 *IUIAutomationElement2) (elemArray *IUIAutomationElementArray, err error) {
	hr, _, _ := syscall.SyscallN(elem2.VTable().Get_CachedFlowsFrom, uintptr(unsafe.Pointer(elem2)),
		uintptr(unsafe.Pointer(&elemArray)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
