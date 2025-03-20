package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	//Unknown Char + End of File
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literals
	IDENT = "IDENT"
	INT   = "INT"

	//Operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LCURLY = "{"
	RCURLY = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
