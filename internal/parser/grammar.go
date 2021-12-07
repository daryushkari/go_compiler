package parser

// eps used as epsilon
var ter = map[string]bool{
	"a":   true,
	"b":   true,
	"d":   true,
	"e":   true,
	"f":   true,
	"eps": true,
}

var nTer = map[string]bool{
	"S": true,
	"A": true,
	"B": true,
	"C": true,
}

var grammars = map[string][][]string{
	"S": {
		{"A", "C", "B"},
		{"C", "b", "B"},
		{"B", "a"},
	},

	"A": {
		{"d", "a"},
		{"B", "C"},
	},

	"B": {{"e"}, {"eps"}},

	"C": {{"f"}, {"eps"}},
}
