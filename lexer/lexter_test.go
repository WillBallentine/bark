package lexer

import (
	"github.com/WillBallentine/bark/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `toy five = 5;
	toy ten = 10;

	toy add = trick(x, y) {
	x + y;
	};

	toy result = add(five, ten);

	!-/*5;
	5 < 10 > 5;
	borkf (5 < 10) {
		fetchit goodboi;
	} woofwise {
		fetchit badboi;
	}

	10 == 10;
	10 != 9;
	"foobar"
	"foo bar"
	`

	tests := []struct {
		expectedType    token.TokenType
		expextedLiteral string
	}{
		{token.LET, "toy"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "toy"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "toy"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "trick"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "toy"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "borkf"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "fetchit"},
		{token.TRUE, "goodboi"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "woofwise"},
		{token.LBRACE, "{"},
		{token.RETURN, "fetchit"},
		{token.FALSE, "badboi"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		//{token.LBRACKET, "["},
		//{token.INT, "1"},
		//{token.COMMA, ","},
		//{token.INT, "2"},
		//{token.RBRACKET, "]"},
		//{token.SEMICOLON, ";"},
		//{token.LBRACE, "{"},
		//{token.STRING, "foo"},
		//{token.COLON, ":"},
		//{token.STRING, "bar"},
		//{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expextedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expextedLiteral, tok.Literal)
		}
	}
}
