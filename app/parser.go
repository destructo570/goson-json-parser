package app

import (
	"errors"
	"goson/enum"
	"goson/models"
)

type Parser struct {
	tokens  []models.Token
	current int
}

func (p *Parser) Parse(tokens []models.Token) (models.JsonValue, error) {
	p.tokens = tokens
	p.current = 0

	value, err := p.ParseValue()

	if err == nil {
		return value, err
	}

	return nil, err
}

func (p *Parser) ParseValue() (models.JsonValue, error) {
	switch p.Peek().Type {
	case enum.LEFT_BRACE:
		return p.ParseObject()
	case enum.LEFT_BRACKET:
		return p.ParseArray()
	case enum.STRING:
		token, err := p.Consume(enum.STRING, "expected string")
		if err != nil {
			return nil, err
		}
		str := token.Literal.(string)
		node := models.JsonString{Value: str}
		return node, nil
	case enum.NUMBER:
		token, err := p.Consume(enum.NUMBER, "expected number")
		if err != nil {
			return nil, err
		}
		num := token.Literal.(float64)
		node := models.JsonNumber{Value: num}
		return node, nil
	case enum.TRUE:
		_, err := p.Consume(enum.TRUE, "expected boolean")
		if err != nil {
			return nil, err
		}
		node := models.JsonBool{Value: true}
		return node, nil
	case enum.FALSE:
		_, err := p.Consume(enum.FALSE, "expected boolean")
		if err != nil {
			return nil, err
		}
		node := models.JsonBool{Value: false}
		return node, nil
	case enum.NULL:
		_, err := p.Consume(enum.NULL, "expected null")
		if err != nil {
			return nil, err
		}

		node := models.JsonNull{}
		return node, nil
	default:
		return nil, errors.New("Syntax Error")
	}
}

func (p *Parser) ParseObject() (models.JsonValue, error) {
	_, lerr := p.Consume(enum.LEFT_BRACE, "Expected left brace")
	if lerr != nil {
		return nil, lerr
	}

	obj := models.JsonObject{Fields: map[string]models.JsonValue{}}

	if p.Peek().Type == enum.RIGHT_BRACE {
		_, rerr := p.Consume(enum.RIGHT_BRACE, "Expected Right brace")
		if rerr != nil {
			return nil, rerr
		}
		return obj, nil
	} else {
		for p.Peek().Type != enum.EOF {

			keyToken, err := p.Consume(enum.STRING, "Expected String")
			if err != nil {
				return nil, err
			}
			_, err2 := p.Consume(enum.COLON, "Expected Colon")
			if err2 != nil {
				return nil, err2
			}
			val, err3 := p.ParseValue()
			if err3 != nil {
				return nil, err3
			}
			key := keyToken.Literal.(string)
			obj.Fields[key] = val

			if !p.Match(enum.COMMA) {
				break
			}
		}

		_, err := p.Consume(enum.RIGHT_BRACE, "Expected Right Brace")
		if err != nil {
			return nil, err
		}
		return obj, nil
	}
}

func (p *Parser) ParseArray() (models.JsonValue, error) {
	_, lerr := p.Consume(enum.LEFT_BRACKET, "Expected left bracket")
	if lerr != nil {
		return nil, lerr
	}

	obj := models.JsonArray{Elements: []models.JsonValue{}}

	if p.Peek().Type == enum.RIGHT_BRACKET {
		_, rerr := p.Consume(enum.RIGHT_BRACKET, "Expected Right bracket")
		if rerr != nil {
			return nil, rerr
		}
		return obj, nil
	} else {
		for p.Peek().Type != enum.EOF {

			val, err3 := p.ParseValue()
			if err3 != nil {
				return nil, err3
			}
			// _, err2 := p.Consume(enum.COMMA, "Expected Comma")
			// if err2 != nil {
			// 	return nil, err2
			// }

			obj.Elements = append(obj.Elements, val)

			if !p.Match(enum.COMMA) {
				break
			}
		}

		_, err := p.Consume(enum.RIGHT_BRACKET, "Expected Right Bracket")
		if err != nil {
			return nil, err
		}
		return obj, nil
	}
}

func (p *Parser) Peek() models.Token {
	currentToken := p.current
	return p.tokens[currentToken]
}

func (p *Parser) Advance() {
	if !p.IsAtEnd() {
		p.current = p.current + 1
	}
}

func (p *Parser) Check(t enum.TokenType) bool {
	if p.IsAtEnd() {
		return false
	}
	return p.Peek().Type == t
}

func (p *Parser) Match(types ...enum.TokenType) bool {
	for i := 0; i < len(types); i++ {
		if p.Check(types[i]) {
			p.Advance()
			return true
		}
	}
	return false
}

func (p *Parser) Consume(expected enum.TokenType, message string) (models.Token, error) {
	if p.Check(expected) {
		p.Advance()
		return p.tokens[p.current-1], nil
	} else {
		return p.tokens[p.current], errors.New(message)
	}
}

func (p *Parser) IsAtEnd() bool {
	return p.Peek().Type == enum.EOF
}

func (p *Parser) Previous() models.Token {
	return p.tokens[p.current-1]
}
