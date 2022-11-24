package pkg

type TokenType string

type Token struct {
	Type TokenType
	Value string
}

type Lexer struct {
	input string
	currPos int
	nextPos int
	ch byte
}

func NewLexer(contents string) *Lexer {
	l := &Lexer{
		input: contents,
	}

	return l
}
