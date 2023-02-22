package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement3 struct {
	ole.IUnknown
}

type IUIAutomationElement3Vtbl struct {
	ole.IUnknownVtbl
}

// IID为8471DF34-AEE0-4A01-A7DE-7DB9AF12C296
var IID_IUIAutomationElement3 = ole.NewGUID("8471DF34-AEE0-4A01-A7DE-7DB9AF12C296")
