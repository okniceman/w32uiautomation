package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement5 struct {
	ole.IUnknown
}

type IUIAutomationElement5Vtbl struct {
	ole.IUnknownVtbl
}

// IID为98141C1D-0D0E-4175-BBE2-6BFF455842A7
var IID_IUIAutomationElement5 = ole.NewGUID("98141C1D-0D0E-4175-BBE2-6BFF455842A7")
