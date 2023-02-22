package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement2 struct {
	ole.IUnknown
}

type IUIAutomationElement2Vtbl struct {
	ole.IUnknownVtbl
}

// IID为6749c683-f70d-4487-a698-5f79d55290d6
var IID_IUIAutomationElement2 = ole.NewGUID("6749c683-f70d-4487-a698-5f79d55290d6")
