package lexer

import (
	"go_compiler/internal/lexer/dfa"
	"go_compiler/internal/utilities/strUtl"
)

/*
lexer names:
assign
compare
delimiter
keyword
logical
mainType
operator
punctuation
shift

Float
Comment
Identifier
Int
String
WhiteSpace

*/

type LexToken struct {
	Word  string
	Token string
}

func Lexer(s string)(tkList []*LexToken){
	lexed := s
	tkList = []*LexToken{}


	for i := len(lexed); i > 0; i--{
		hasTk, tk := checkConst(lexed[:i])
		if hasTk{
			tkList = append(tkList, &LexToken{Token: tk, Word: lexed[:i]})
			lexed = lexed[i:]
			i = len(lexed)
			continue
		}

		hasTk, tk = dfa.ReturnFitDFA(lexed[:i])
		if hasTk{
			tkList = append(tkList, &LexToken{Token: tk, Word: lexed[:i]})
			lexed = lexed[i:]
			i = len(lexed)
			continue
		}

	}

	return tkList
}

func checkConst(s string) (hasTk bool, tk string) {
	for key, element := range lexConst {
		if strUtl.CheckHasString(s, element){
			return true, key
		}
	}
	return false, ""
}


