package dfa

import (
	"go_compiler/internal/utilities/strUtl"
)

type tableDFA struct{
	Name string
	Alphabet []string
	Reject int
	Accept int
	Table map[string][]int
}

func runDFA(tb tableDFA, s string)bool{
	state := 0
	for i := 0; i < len(s); i++{
		if state == tb.Reject{
			return false
		}
		if !strUtl.CheckHasString(string(s[i]), tb.Alphabet){
			return false
		}

		if val, ok := tb.Table[string(s[i])]; ok {
			state = val[state]
		}else{
			val = tb.Table["others"]
			state = val[state]
		}
	}
	if state == tb.Accept{
		return true
	}
	return false
}


func checkInt(tb tableDFA,s string)bool{
	if tb.Name == "Int"{
		if s == "0"{
			return true
		}
		if len(s) > 2 && s[:2] == "-0"{
			return false
		}
		return runDFA(tb, s)
	}
	return false
}

func ReturnFitDFA(s string)(hasType bool,tkType string){
	tbList := []tableDFA{
		CmtDFATable,
		WspcDFATable,
		StrDFATable,
		FltDFATable,
		IdnDFATable,
	}

	for i := 0; i < len(tbList); i++{
		if runDFA(tbList[i], s){
			return true,tbList[i].Name
		}
	}
	if checkInt(IntDFATable, s){
		return true, IntDFATable.Name
	}
	return false, ""

}