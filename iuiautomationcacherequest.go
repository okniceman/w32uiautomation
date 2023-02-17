package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationCacheRequest struct {
	ole.IUnknown
}

type IUIAutomationCacheRequestVtbl struct {
	ole.IUnknownVtbl
	AddPattern                uintptr
	AddProperty               uintptr
	Clone                     uintptr
	Get_AutomationElementMode uintptr
	Get_TreeFilter            uintptr
	Get_TreeScope             uintptr
	Put_AutomationElementMode uintptr
	Put_TreeFilter            uintptr
	Put_TreeScope             uintptr
}

var IID_IUIAutomationCacheRequest = &ole.GUID{0xb32a92b5, 0xbc25, 0x4078, [8]byte{0x9c, 0x08, 0xd7, 0xee, 0x95, 0xc4, 0x8e, 0x03}}

func (cacheRequet *IUIAutomationCacheRequest) VTable() *IUIAutomationCacheRequestVtbl {
	return (*IUIAutomationCacheRequestVtbl)(unsafe.Pointer(cacheRequet.RawVTable))
}

// AddPattern
//
//	@Description: Adds a control pattern to the cache request.
//	@receiver cacheRequet
//	@param patternId
//	@return err
func (cacheRequet *IUIAutomationCacheRequest) AddPattern(patternId PATTERNID) (err error) {
	return addPattern(cacheRequet, patternId)
}

// AddProperty
//
//	@Description: Adds a property to the cache request.
//	@receiver cacheRequet
//	@param propertyId
//	@return err
func (cacheRequet *IUIAutomationCacheRequest) AddProperty(propertyId PROPERTYID) (err error) {
	return addProperty(cacheRequet, propertyId)
}

// GetTreeScope
//
//	@Description: Specifies the scope of caching.
//	@receiver cacheRequet
//	@return scope
//	@return err
func (cacheRequet *IUIAutomationCacheRequest) GetTreeScope() (scope TreeScope, err error) {
	return getTreeScope(cacheRequet)
}

// PutTreeScope
//
//	@Description: Specifies the scope of caching.
//	@receiver cacheRequet
//	@param scope
//	@return err
func (cacheRequet *IUIAutomationCacheRequest) PutTreeScope(scope TreeScope) (err error) {
	return putTreeScope(cacheRequet, scope)
}

func putTreeScope(cacheRequest *IUIAutomationCacheRequest, scope TreeScope) (err error) {
	hr, _, _ := syscall.SyscallN(cacheRequest.VTable().Put_TreeScope, uintptr(unsafe.Pointer(cacheRequest)), uintptr(scope))

	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getTreeScope(cacheRequest *IUIAutomationCacheRequest) (scope TreeScope, err error) {
	hr, _, _ := syscall.SyscallN(cacheRequest.VTable().Get_TreeScope, uintptr(unsafe.Pointer(cacheRequest)), uintptr(scope))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func addPattern(cacheRequest *IUIAutomationCacheRequest, patternId PATTERNID) (err error) {
	hr, _, _ := syscall.SyscallN(cacheRequest.VTable().AddPattern, uintptr(unsafe.Pointer(cacheRequest)),
		uintptr(patternId))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func addProperty(cacheRequet *IUIAutomationCacheRequest, propertyId PROPERTYID) (err error) {
	hr, _, _ := syscall.SyscallN(cacheRequet.VTable().AddProperty, uintptr(unsafe.Pointer(cacheRequet)), uintptr(propertyId))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
