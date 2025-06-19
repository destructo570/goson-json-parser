package app

import (
	"errors"
	"goson/enum"
	"goson/models"
	"strconv"
)

type Scanner struct {
	Source  string
	Tokens  []models.Token
	start   int
	Line    int
	current int
}

func (s *Scanner) ScanTokens() []models.Token {
	for !s.IsAtEnd() {
		s.start = s.current
		s.ScanToken()
	}

	return s.Tokens
}

func (s *Scanner) Advance() rune {
	runes := []rune(s.Source)
	currentRune := runes[s.current]
	s.current++
	return currentRune
}

func (s *Scanner) CurrentChar() rune {
	runes := []rune(s.Source)
	return runes[s.current]
}

func (s *Scanner) AddToken(tokenType enum.TokenType) {
	s.AddTokenWithLiteral(tokenType, nil)
}

func (s *Scanner) AddTokenWithLiteral(tokenType enum.TokenType, literal any) {
	text := ""

	if tokenType == enum.STRING {
		text = Substring(s.Source, s.start+1, s.current)
	} else {
		text = Substring(s.Source, s.start, s.current)
	}
	s.Tokens = append(s.Tokens, models.Token{
		Type:    tokenType,
		Lexeme:  text,
		Line:    s.Line,
		Literal: literal,
	})
}

func (s *Scanner) IsAtEnd() bool {
	return s.current >= len(s.Source)
}

func (s *Scanner) ScanToken() error {
	char := s.Advance()

	switch char {
	case '[':
		s.AddToken(enum.LEFT_BRACKET)
	case ']':
		s.AddToken(enum.RIGHT_BRACKET)
	case '{':
		s.AddToken(enum.LEFT_BRACE)
	case '}':
		s.AddToken(enum.RIGHT_BRACE)
	case ':':
		s.AddToken(enum.COLON)
	case ',':
		s.AddToken(enum.COMMA)
	case '"':
		s.MatchString()
	case ' ':
	case '\r':
	case '\t':
		// Ignore whitespace.
		break
	case '\n':
		s.Line++
	default:
		if s.isDigit(char) {
			s.MatchNumber()
		} else {
			return errors.New("unexpected character")
		}
	}

	return nil
}

func (s *Scanner) isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func (s *Scanner) MatchNumber() {
	for s.isDigit(s.peek()) {
		s.Advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.Advance()

		for s.isDigit(s.peek()) {
			s.Advance()
		}
	}
	s.Advance()
	valueStr := Substring(s.Source, s.start, s.current)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		// handle the error, or panic/log as per your needs
		panic("Invalid float number: " + valueStr)
	}

	s.AddTokenWithLiteral(enum.NUMBER, value)
}

func (s *Scanner) peekNext() rune {
	if s.current+2 >= len(s.Source) {
		return rune(0)
	}
	runes := []rune(s.Source)
	return runes[s.current+2]
}

func (s *Scanner) MatchString() error {
	for s.peek() != '"' && !s.IsAtEnd() {
		if s.peek() == '\n' {
			s.Line++
		}
		s.Advance()
	}

	if s.IsAtEnd() {
		return errors.New("Unterminated string")
	}

	s.Advance()
	value := Substring(s.Source, s.start+1, s.current)
	s.AddTokenWithLiteral(enum.STRING, value)
	s.Advance()
	return nil
}

func (s *Scanner) Match(expected rune) bool {
	if s.IsAtEnd() || s.CurrentChar() != expected {
		return false
	}

	return true
}

func (s *Scanner) peek() rune {
	if s.IsAtEnd() {
		return rune(0)
	}
	runes := []rune(s.Source)
	return runes[s.current+1]
}

// Helpers
func Substring(input string, start int, end int) string {
	runes := []rune(input)
	return string(runes[start:end])
}
