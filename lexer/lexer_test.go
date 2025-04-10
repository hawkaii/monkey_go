package lexer

import (
	"testing"

	"github.com/hawkaii/monkey_go/token"
)

func TestNextToken(t *testing.T) {

	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{

		{token.ASSIGN, "="},
		{token.PLUS, "+"},
	}

}
