// Package lexer implements the lexical analysis that is used to
// transform the source code input into a stream of tokens for
// parsing by the parser.
package lexer

import (
	"github.com/ehsaaniqbal/ma/token"
)

// Represents the Lexer and contains the source input
// and internal state
type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position in input (points to the next char)
	ch           byte // character being analysed
}

// Returns a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

// Returns the next token read from the input stream
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Skip over whitespace
	l.eatWhitespace()

	// Match current character to Token
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Returns the next character present in the input string
// and advances the positon in the input string
func (l *Lexer) readChar() {
	// Check if end of input, else access the next character
	if l.readPosition >= len(l.input) {
		l.ch = 0 // NUL (ASCII)
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// Reads in an identifier and advances the lexer’s positions
// until it encounters a non-letter-character.
func (l *Lexer) readIdentifier() string {
	pos := l.position

	for isLetter(l.ch) {
		l.readChar()
		// fmt.Println(l.input[pos: l.position])
		if token.LookupIdentifier(l.input[pos:l.position]) == "IDENT" {
			if isWhitespace(l.ch) {
				// fmt.Println("WHITESPACE", l.peekChar())
				pc := l.peekChar()
				if pc == 'd' || pc == 'b' || pc == 'y' {
					l.readChar()
				}
			}
		}

	}

	return l.input[pos:l.position]
}

func (l *Lexer) readString() string {
	pos := l.position + 1

	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition > len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) eatWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func isWhitespace(ch byte) bool {
	if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
		return true
	}

	return false
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
