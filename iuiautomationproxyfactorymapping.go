package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationProxyFactoryMapping struct {
	ole.IUnknown
}

type IUIAutomationProxyFactoryMappingVtbl struct {
	ole.IUnknownVtbl
	ClearTable          uintptr
	Get_Count           uintptr
	GetEntry            uintptr
	GetTable            uintptr
	InsertEntries       uintptr
	InsertEntry         uintptr
	RemoveEntry         uintptr
	RestoreDefaultTable uintptr
	SetTable            uintptr
}

func (mapping *IUIAutomationProxyFactoryMapping) VTable() *IUIAutomationProxyFactoryMappingVtbl {
	return (*IUIAutomationProxyFactoryMappingVtbl)(unsafe.Pointer(mapping.RawVTable))
}
