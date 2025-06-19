package enum

type TokenType string

const (
	NUMBER        TokenType = "NUMBER"
	STRING        TokenType = "STRING"
	TRUE          TokenType = "TRUE"
	FALSE         TokenType = "FALSE"
	LEFT_BRACE    TokenType = "LEFT_BRACE"
	RIGHT_BRACE   TokenType = "RIGHT_BRACE"
	LEFT_BRACKET  TokenType = "LEFT_BRACKET"
	RIGHT_BRACKET TokenType = "RIGHT_BRACKET"
	COMMA         TokenType = "COMMA"
	COLON         TokenType = "COLON"
	EOF           TokenType = "EOF"
)
