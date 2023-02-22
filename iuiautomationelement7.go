package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement7 struct {
	ole.IUnknown
}

type IUIAutomationElement7Vtbl struct {
	ole.IUnknownVtbl
}

// IID为204e8572-cfc3-4c11-b0c8-7da7420750b7
var IID_IUIAutomationElement7 = ole.NewGUID("204e8572-cfc3-4c11-b0c8-7da7420750b7")
