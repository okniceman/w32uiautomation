package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement6 struct {
	ole.IUnknown
}

type IUIAutomationElement6Vtbl struct {
	ole.IUnknownVtbl
}

// IID为4780d450-8bca-4977-afa5-a4a517f555e3
var IID_IUIAutomationElement6 = ole.NewGUID("4780d450-8bca-4977-afa5-a4a517f555e3")
