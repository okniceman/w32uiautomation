package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement9 struct {
	ole.IUnknown
}

type IUIAutomationElement9Vtbl struct {
	ole.IUnknownVtbl
}

// IID为39325fac-039d-440e-a3a3-5eb81a5cecc3
var IID_IUIAutomationElement9 = ole.NewGUID("39325fac-039d-440e-a3a3-5eb81a5cecc3")
