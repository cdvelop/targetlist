package targetlist

func (l targetList) ResetAdapterView() {

	l.Log("** ResetAdapterView targetList")

	l.CallJsFunctionObject("resetTargetList", true)

	// err := l.CallFunction("resetargetListx")
	// if err != nil {
	// 	return err
	// }

}

// func (l targetList) callJsFunc(func_name string, enable bool) {

// 	err := l.CallFunction(func_name, map[string]interface{}{
// 		"enable": enable,
// 		"query":  f.theme.QuerySelectorObject(f.Object.ModuleName, f.Object.ObjectName),
// 	})
// 	if err != nil {
// 		f.Log(err)
// 	}

// }
