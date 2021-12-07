package lexer

var assign = []string{
	"=",
	":=",
}

var compare = []string{
	"==",
	"!=",
	"<",
	"<=",
	">",
	">=",
}

var delimiter = []string{
	"}",
	"{",
	"]",
	"[",
	")",
	"(",
}

var keyword = []string{
	"break",
	"default",
	"func",
	"interface",
	"select",
	"case",
	"defer",
	"go",
	"map",
	"struct",
	"chan",
	"else",
	"goto",
	"package",
	"switch",
	"const",
	"fallthrough",
	"if",
	"range",
	"type",
	"continue",
	"for",
	"import",
	"return",
	"var",
}

var logical = []string{
	"&&",
	"||",
	"!",
	"&",
	"|",
}

var mainType = []string{
	"bool",
	"byte",
	"complex",
	"complex64",
	"complex128",
	"uint16",
	"float32",
	"float64",
	"int",
	"int8",
	"int16",
	"uint32",
	"int32",
	"int64",
	"uint64",
	"string",
	"uint",
	"uint8",
	"*bool",
	"*byte",
	"*complex",
	"*complex64",
	"*complex128",
	"*uint16",
	"*float32",
	"*float64",
	"*int",
	"*int8",
	"*int16",
	"*uint32",
	"*int32",
	"*int64",
	"*uint64",
	"*string",
	"*uint",
	"*uint8",
}

var operator = []string{
	"-",
	"+",
	"*",
	"/",
	"%",
}

var punctuation = []string{
	".",
	";",
	",",
	"->",
	"<-",
	":",
}

var shift = []string{
	">>",
	"<<",
}

var lexConst = map[string][]string{
	"Assign": assign,
	"Compare": compare,
	"Delimiter": delimiter,
	"Keyword": keyword,
	"Logical": logical,
	"MainType": mainType,
	"Operator": operator,
	"Punctuation": punctuation,
	"Shift": shift,
}