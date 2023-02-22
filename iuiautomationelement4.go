package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement4 struct {
	ole.IUnknown
}

type IUIAutomationElement4Vtbl struct {
	ole.IUnknownVtbl
}

// IID为3B6E233C-52FB-4063-A4C9-77C075C2A06B
var IID_IUIAutomationElement4 = ole.NewGUID("3B6E233C-52FB-4063-A4C9-77C075C2A06B")
