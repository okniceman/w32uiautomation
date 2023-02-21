package w32uiautomation

import (
	"fmt"

	"github.com/go-ole/go-ole"
)

type TextEditChangeType uintptr

const (
	TextEditChangeType_None = iota
	TextEditChangeType_AutoCorrect
	TextEditChangeType_Composition
	TextEditChangeType_CompositionFinalized
	TextEditChangeType_AutoComplete
)

func (t TextEditChangeType) ToString() string {
	switch t {
	case TextEditChangeType_None:
		return "Not related to a specific change type."
	case TextEditChangeType_AutoCorrect:
		return "Change is from an auto-correct action performed by a control."
	case TextEditChangeType_Composition:
		return "Change is from an IME active composition within a control."
	case TextEditChangeType_CompositionFinalized:
		return "Change is from an IME composition going from active to finalized state within a control."
	case TextEditChangeType_AutoComplete:
		return ""
	default:
		panic(fmt.Errorf("Unknown TextEditChangeType: %d", t))
	}
}

type IUIAutomationTextEditTextChangedEventHandler struct {
	ole.IUnknown
	ref int32
}

type IUIAutomationTextEditTextChangedEventHandlerVtbl struct {
	ole.IUnknownVtbl
	HandleTextEditTextChangedEvent uintptr
}

// IID为92FAA680-E704-4156-931A-E32D5BB38F3F
var IID_IUIAutomationTextEditTextChangedEventHandler = &ole.GUID{0x92FAA680, 0xE704, 0x4156, [8]byte{0x93, 0x1A, 0xE3,
	0x2D, 0x5B, 0xB3, 0x8F, 0x3F}}
