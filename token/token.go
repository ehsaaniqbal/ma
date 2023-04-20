// Package token implements types and constants to support tokenizing
// the input source before passing the stream of tokens to the parser.
package token

// Represents the type of Token
type TokenType string

// Holds a single Token type and its literal value
type Token struct {
	Type    TokenType
	Literal string
}

// Token type definitions
const (
	ILLEGAL = "ILLEGAL" // Represents an illegal token
	EOF     = "EOF"     // End of file

	// Identifiers + literals
	IDENT  = "IDENT"  // add, x, y, foobar
	INT    = "INT"    // 1, 2, 3
	STRING = "STRING" // "immersion"

	// Operators
	ASSIGN   = "ASSIGN"
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"initiative":    FUNCTION,
	"ma":            LET,
	"true":          TRUE,
	"false":         FALSE,
	"super dream":   IF,
	"regular":       ELSE,
	"god bless you": RETURN,
}

// Looks up the identifier in ident and returns the appropriate
// token type depending on whether the identifier is user-defined or a keyword
func LookupIdentifier(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	}

	return IDENT
}
