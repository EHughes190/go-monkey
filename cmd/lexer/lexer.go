package lexer

import (
	"github.com/EHughes190/monkey/cmd/token"
)

type Lexer struct {
	input        string
	position     int // current position in our input (points to current char)
	readPosition int // current read position in input (next char after curr char)
	ch           byte
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// whitespace is not parsed in Monkey
	l.skipWhitespace()

	// determine what the token is
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
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// returns the word
			tok.Literal = l.readIdentifier()
			// checks if the literal is a keyword or not (and assigns a value either way)
			tok.Type = token.LookUpIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// advance our position
	l.readChar()
	return tok
}

// creates a new instance of a lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

// currently only supports ASCII not unicode (treats each ch as just 1 byte each)
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// Take a look at the next char but don't move our position
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	// whilst the char is a 'letter', advance our position
	for isLetter(l.ch) {
		l.readChar()
	}

	// the string from our start position to end of this subset of letters
	return l.input[position:l.position]
}

// deterines what our language classes as a 'letter'. We could add new ones here if we want!
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// if the char is whitespace, move our position forward
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position

	// whilst the char is a 'number', advance our position
	for isDigit(l.ch) {
		l.readChar()
	}

	// the string from our start position to end of this subset of letters
	return l.input[position:l.position]
}

// Monkey only supports INTs and not floats (yet)
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
