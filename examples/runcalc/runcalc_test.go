package main

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/go-ole/go-ole"
	wa "github.com/lunarforge/w32uiautomation"
)

const (
	calculatorName          = "Calculator"
	clearButtonAutomationId = "clearButton" //"81"
	twoButtonAutomationId   = "num2Button"  //"132"
	threeButtonAutomationId = "num3Button"  //"133"
	plusButtonAutomationId  = "plusButton"  //"93"
	equalButtonAutomationId = "equalButton" //"121"
)

func TestRunCalc(t *testing.T) {
	err := runCalc()
	if err != nil {
		panic(err)
	}

}

func runCalc() error {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	err := exec.Command("calc.exe").Start()
	if err != nil {
		return err
	}

	auto, err := wa.NewUIAutomation()
	if err != nil {
		return err
	}

	root, err := auto.GetRootElement()
	if err != nil {
		return err
	}
	defer root.Release()

	condVal := wa.NewVariantString(calculatorName)
	condition, err := auto.CreatePropertyCondition(wa.UIA_NamePropertyId, condVal)
	if err != nil {
		return err
	}
	calc, err := wa.WaitFindFirst(auto, root, wa.TreeScope_Children, condition)
	if err != nil {
		return err
	}

	calcName, err := calc.Get_CurrentName()
	if err != nil {
		return err
	}
	fmt.Printf("calcName=%v\n", calcName)

	pushButton(auto, calc, clearButtonAutomationId)
	if err != nil {
		return err
	}

	pushButton(auto, calc, twoButtonAutomationId)
	if err != nil {
		return err
	}

	pushButton(auto, calc, plusButtonAutomationId)
	if err != nil {
		return err
	}

	pushButton(auto, calc, threeButtonAutomationId)
	if err != nil {
		return err
	}

	pushButton(auto, calc, equalButtonAutomationId)
	if err != nil {
		return err
	}

	return nil
}

func pushButton(auto *wa.IUIAutomation, calc *wa.IUIAutomationElement, automationId string) error {
	condition, err := auto.CreatePropertyCondition(
		wa.UIA_AutomationIdPropertyId,
		wa.NewVariantString(automationId))
	if err != nil {
		return err
	}

	button, err := wa.WaitFindFirst(auto, calc, wa.TreeScope_Subtree, condition)
	if err != nil {
		return err
	}
	return wa.Invoke(button)
}
