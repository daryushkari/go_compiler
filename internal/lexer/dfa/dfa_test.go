package dfa

import (
	"fmt"
	"testing"
)

type tokenTest struct {
	Word  string
	Token string
}

func cmtData()([]tokenTest){
	testList := []tokenTest{
		{Word: "//sdfsdfsderwr", Token: "False"},
		{Word: "//sdfsdfsderwr/", Token: "False"},
		{Word: "/sd//fsdfsderwr//", Token: "False"},
		{Word: "///sdfsdfsderwr/", Token: "False"},
		{Word: "//;//", Token: "False"},
		{Word: "//", Token: "False"},
		{Word: "/z//z/", Token: "False"},
		{Word: "/ ///", Token: "False"},
		{Word: "///", Token: "False"},
		{Word: "sdf   	sdfdf//", Token: "False"},
		{Word: "/sdf   	sdfdf//", Token: "False"},
		{Word: "", Token: "False"},

		{Word: "//    \n \t //", Token: "Comment"},
		{Word: "//  sdf  \nsdf \t sdfdf//", Token: "Comment"},
		{Word: "//  sdfdf dfsdf ffff' ' ' df //", Token: "Comment"},
		{Word: "////", Token: "Comment"},
		{Word: "//asdfsdfsdf//", Token: "Comment"},
		{Word: "//////////////", Token: "Comment"},
	}

	return testList
}


func wspcData()([]tokenTest){
	testList := []tokenTest{
		tokenTest{Word: "a ", Token: "False"},
		tokenTest{Word: " ; ", Token: "False"},
		tokenTest{Word: " \t \n h \n", Token: "False"},
		tokenTest{Word: "", Token: "False"},
		tokenTest{Word: "\ty\t", Token: "False"},
		tokenTest{Word: "\t\tt", Token: "False"},
		tokenTest{Word: "", Token: "False"},

		tokenTest{Word: "\t", Token: "WhiteSpace"},
		tokenTest{Word: "\n", Token: "WhiteSpace"},
		tokenTest{Word: " ", Token: "WhiteSpace"},
		tokenTest{Word: "\t \n \n \t     ", Token: "WhiteSpace"},
		tokenTest{Word: "  \t \t \n ", Token: "WhiteSpace"},
		tokenTest{Word: " ", Token: "WhiteSpace"},
		tokenTest{Word: "    ", Token: "WhiteSpace"},
		tokenTest{Word: "\t\t", Token: "WhiteSpace"},
		tokenTest{Word: "\n\n", Token: "WhiteSpace"},
	}

	return testList
}


func strData()([]tokenTest){
	testList := []tokenTest{
		tokenTest{Word: " qw", Token: "False"},
		tokenTest{Word: "_wewe", Token: "False"},
		tokenTest{Word: "'\"", Token: "False"},
		tokenTest{Word: "\"", Token: "False"},
		tokenTest{Word: "'''", Token: "False"},
		tokenTest{Word: "''sdsd", Token: "False"},
		tokenTest{Word: "\"\"dfg", Token: "False"},
		tokenTest{Word: "\" \" ", Token: "False"},
		tokenTest{Word: "\"\"\"", Token: "False"},
		tokenTest{Word: "''''", Token: "False"},
		tokenTest{Word: "';'", Token: "False"},
		tokenTest{Word: "", Token: "False"},

		tokenTest{Word: "'sdfd dfdf'", Token: "String"},
		tokenTest{Word: "\" hello world \"", Token: "String"},
		tokenTest{Word: "\"''''\"", Token: "String"},
		tokenTest{Word: "'\"sdfsdf'", Token: "String"},
		tokenTest{Word: "' sdf sdf sdf'", Token: "String"},
		tokenTest{Word: "'456hug'", Token: "String"},
		tokenTest{Word: "''", Token: "String"},
		tokenTest{Word: "\"\"", Token: "String"},
	}

	return testList
}


func fltData()([]tokenTest){
	testList := []tokenTest{
		tokenTest{Word: "123123", Token: "False"},
		tokenTest{Word: "123e", Token: "False"},
		tokenTest{Word: "123.", Token: "False"},
		tokenTest{Word: "00.12", Token: "False"},
		tokenTest{Word: "-.12", Token: "False"},
		tokenTest{Word: "12-12", Token: "False"},
		tokenTest{Word: ".123", Token: "False"},
		tokenTest{Word: "-123", Token: "False"},
		tokenTest{Word: "123.11 ", Token: "False"},
		tokenTest{Word: "123.11-", Token: "False"},
		tokenTest{Word: "", Token: "False"},

		tokenTest{Word: "-123.123", Token: "Float"},
		tokenTest{Word: "123.123", Token: "Float"},
		tokenTest{Word: "1.0", Token: "Float"},
		tokenTest{Word: "-0.1", Token: "Float"},
		tokenTest{Word: "1.00", Token: "Float"},
		tokenTest{Word: "-1.0", Token: "Float"},
		tokenTest{Word: "123.123", Token: "Float"},
		tokenTest{Word: "123.0", Token: "Float"},
		tokenTest{Word: "0.0", Token: "Float"},
	}

	return testList
}


func intData()([]tokenTest){
	testList := []tokenTest{
		tokenTest{Word: "123 123", Token: "False"},
		tokenTest{Word: "--123", Token: "False"},
		tokenTest{Word: "001", Token: "False"},
		tokenTest{Word: "-001", Token: "False"},
		tokenTest{Word: "-01", Token: "False"},
		tokenTest{Word: "0123", Token: "False"},
		tokenTest{Word: "12m12", Token: "False"},
		tokenTest{Word: "120-", Token: "False"},
		tokenTest{Word: "", Token: "False"},

		tokenTest{Word: "-120", Token: "Int"},
		tokenTest{Word: "120", Token: "Int"},
		tokenTest{Word: "120034", Token: "Int"},
		tokenTest{Word: "1", Token: "Int"},
		tokenTest{Word: "0", Token: "Int"},
		tokenTest{Word: "-1", Token: "Int"},
	}

	return testList
}


func idnData()([]tokenTest){
	testList := []tokenTest{
		tokenTest{Word: "1qwe", Token: "False"},
		tokenTest{Word: "1_asd", Token: "False"},
		tokenTest{Word: "qwe qwe", Token: "False"},
		tokenTest{Word: "_asdf sdf", Token: "False"},
		tokenTest{Word: "", Token: "False"},
		tokenTest{Word: " q", Token: "False"},
		tokenTest{Word: "1", Token: "False"},
		tokenTest{Word: "123a", Token: "False"},
		tokenTest{Word: "1_", Token: "False"},

		tokenTest{Word: "_", Token: "Identifier"},
		tokenTest{Word: "_1", Token: "Identifier"},
		tokenTest{Word: "_q", Token: "Identifier"},
		tokenTest{Word: "q", Token: "Identifier"},
		tokenTest{Word: "aA12", Token: "Identifier"},
		tokenTest{Word: "as_qwe", Token: "Identifier"},
		tokenTest{Word: "as12_qwe12", Token: "Identifier"},
		tokenTest{Word: "A", Token: "Identifier"},
	}

	return testList
}

func TestInt(t *testing.T){
	inD := intData()
	fmt.Println("testing Intiger")
	for i := 0; i < len(inD); i++{
		res := checkInt(IntDFATable,inD[i].Word)
		if res{
			if inD[i].Token != "False"{
				fmt.Println("success word:_", inD[i].Word, "_token", inD[i].Token)
			}else{
				res := checkInt(IntDFATable,inD[i].Word)
				t.Errorf("word: _%s_, token: _%s_. ans: %t ", inD[i].Word, inD[i].Token, res)
			}
		}else{
			if inD[i].Token == "False"{
				fmt.Println("success word:_", inD[i].Word, "_token", inD[i].Token)
			}else{
				res := checkInt(IntDFATable,inD[i].Word)
				t.Errorf("word: %s, token: %s. ans: %t ", inD[i].Word, inD[i].Token, res)
			}
		}
	}

}

func TestDFAs(t *testing.T) {

	Dlist := [][]tokenTest{
		cmtData(),
		wspcData(),
		strData(),
		fltData(),
		idnData(),
	}
	TList := []tableDFA{
		CmtDFATable,
		WspcDFATable,
		StrDFATable,
		FltDFATable,
		IdnDFATable,
	}

	for j := 0; j < len(Dlist); j++{
		fmt.Println("Testing: ", TList[j].Name)

		for i := 0; i < len(Dlist[j]); i++{
			res := runDFA(TList[j],Dlist[j][i].Word)
			if res{
				if Dlist[j][i].Token != "False"{
					fmt.Println("success word:", Dlist[j][i].Word, " token", Dlist[j][i].Token)
				}else{
					t.Errorf("word: _%s_, token: _%s_. ans: false", Dlist[j][i].Word, Dlist[j][i].Token)
				}
			}else{
				if Dlist[j][i].Token == "False"{
					fmt.Println("success word:", Dlist[j][i].Word, " token", Dlist[j][i].Token)
				}else{
					t.Errorf("word: _%s_, token: _%s_. ans: false", Dlist[j][i].Word, Dlist[j][i].Token)
				}
			}
		}
	}

}