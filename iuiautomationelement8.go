package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement8 struct {
	ole.IUnknown
}

type IUIAutomationElement8Vtbl struct {
	ole.IUnknownVtbl
}

// IID为8C60217D-5411-4CDE-BCC0-1CEDA223830C
var IID_IUIAutomationElement8 = ole.NewGUID("8C60217D-5411-4CDE-BCC0-1CEDA223830C")
