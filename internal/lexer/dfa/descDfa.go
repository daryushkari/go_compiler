package dfa

// comment DFA

var CmtAlph = []string{
	"Q","W","E","R","T","Y","U","I","O","P","A","S","D","F","G","H","J","K","L","Z","X","C","V","B","N","M",
	"q","w","e","r","t","y","u","i","o","p","a","s","d","f","g","h","j","k","l","z","x","c","v","b","n","m",
	"1","2","3","4","5","6","7","8","9","0",
	"!","@","#","$","%","^","&","*","(",")","{","}","[","]",":"," ","\t",":","?","<",">","'","\"","\n",",","\\",
	"/",
}

var CmtAcpt = 4
var CmtRjct = 5

var CmtTable = map[string][]int{
	"others":   {5,   5,   2,   2,   5,   5},
	"/"     :   {1,   2,   3,   4,   4,   5},
}

var CmtDFATable = tableDFA{
	Name: "Comment",
	Alphabet: CmtAlph,
	Reject: CmtRjct,
	Accept: CmtAcpt,
	Table: CmtTable,
}

// float DFA

var FltAlph = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	"-", ".",
}

var FltAcpt = 5
var FltRjct = 2

var FltTable = map[string][]int{
	"1":   {3,   3,   2,   3,   5,  5,  2},
	"2":   {3,   3,   2,   3,   5,  5,  2},
	"3":   {3,   3,   2,   3,   5,  5,  2},
	"4":   {3,   3,   2,   3,   5,  5,  2},
	"5":   {3,   3,   2,   3,   5,  5,  2},
	"6":   {3,   3,   2,   3,   5,  5,  2},
	"7":   {3,   3,   2,   3,   5,  5,  2},
	"8":   {3,   3,   2,   3,   5,  5,  2},
	"9":   {3,   3,   2,   3,   5,  5,  2},
	"0":   {6,   6,   2,   3,   5,  5,  2},
	"-":   {1,   2,   2,   2,   2,  2,  2},
	".":   {2,   2,   2,   4,   2,  2,  4},
}

var FltDFATable = tableDFA{
	Name: "Float",
	Alphabet: FltAlph,
	Reject: FltRjct,
	Accept: FltAcpt,
	Table: FltTable,
}

// Identifier DFA

var IdnAlph = []string{
	"Q","W","E","R","T","Y","U","I","O","P","A","S","D","F","G","H","J","K","L","Z","X","C","V","B","N","M",
	"q","w","e","r","t","y","u","i","o","p","a","s","d","f","g","h","j","k","l","z","x","c","v","b","n","m",
	"1","2","3","4","5","6","7","8","9","0",
	"_",
}

var IdnAcpt = 1
var IdnRjct = 2

var IdnTable = map[string][]int{
	"others":   {1,   1,   2},
	"1"     :   {2,   1,   2},
	"2"     :   {2,   1,   2},
	"3"     :   {2,   1,   2},
	"4"     :   {2,   1,   2},
	"5"     :   {2,   1,   2},
	"6"     :   {2,   1,   2},
	"7"     :   {2,   1,   2},
	"8"     :   {2,   1,   2},
	"9"     :   {2,   1,   2},
	"0"     :   {2,   1,   2},
}

var IdnDFATable = tableDFA{
	Name: "Identifier",
	Alphabet: IdnAlph,
	Reject: IdnRjct,
	Accept: IdnAcpt,
	Table: IdnTable,
}

// Int DFA

var IntAlph = []string {
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "-",
}

var IntAcpt = 1
var IntRjct = 2

var IntTable = map[string][]int{
	"1":   {1,   1,   2},
	"2":   {1,   1,   2},
	"3":   {1,   1,   2},
	"4":   {1,   1,   2},
	"5":   {1,   1,   2},
	"6":   {1,   1,   2},
	"7":   {1,   1,   2},
	"8":   {1,   1,   2},
	"9":   {1,   1,   2},
	"0":   {2,   1,   2},
	"-":   {1,   2,   2},
}

var IntDFATable = tableDFA{
	Name: "Int",
	Alphabet: IntAlph,
	Reject: IntRjct,
	Accept: IntAcpt,
	Table: IntTable,
}

// string DFA

var StrAlph = []string{
	"Q","W","E","R","T","Y","U","I","O","P","A","S","D","F","G","H","J","K","L","Z","X","C","V","B","N","M",
	"q","w","e","r","t","y","u","i","o","p","a","s","d","f","g","h","j","k","l","z","x","c","v","b","n","m",
	"1","2","3","4","5","6","7","8","9","0",
	"!","@","#","$","%","^","&","*","(",")","{","}","[","]",":"," ","\t",":","?","<",">",
	"\"",
	"'",
}

var StrAcpt = 3
var StrRjct = 4

var StrTable = map[string][]int{
	"others":   {4,   1,   2,   4,   4},
	"'"     :   {1,   3,   2,   4,   4},
	"\""    :   {2,   1,   3,   4,   4},
}

var StrDFATable = tableDFA{
	Name: "String",
	Alphabet: StrAlph,
	Reject: StrRjct,
	Accept: StrAcpt,
	Table: StrTable,
}

// whitespace DFA

var WspcAlph = []string{
	"\n", "\t", " ",
}

var WspcAcpt = 1
var WspcRjct = 2

var WspcTable = map[string][]int{
	" " :  {1,   1,   2},
	"\n":  {1,   1,   2},
	"\t":  {1,   1,   2},
}

var WspcDFATable = tableDFA{
	Name: "WhiteSpace",
	Alphabet: WspcAlph,
	Reject: WspcRjct,
	Accept: WspcAcpt,
	Table: WspcTable,
}
