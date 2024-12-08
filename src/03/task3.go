package main

import (
	"common"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type TokenType string

const (
	FuncName    TokenType = "FuncName"
	Argument              = "Argument"
	LeftParen             = "LeftParen"
	RightParen            = "RightParen"
	Comma                 = "Comma"
	EOF                   = "EOF"
	IllegalChar           = "IllegalChar"
)

type ExpressionType string

const (
	Mul     ExpressionType = "Mul"
	Do                     = "Do"
	Dont                   = "Don't"
	Invalid                = "Invalid"
)

type Token struct {
	tokenType  TokenType
	tokenValue string
}

type Lexer struct {
	input   string
	current int
}

type Expression struct {
	name           string
	args           []string
	valid          bool
	errorTokenType TokenType
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:   strings.TrimSpace(input),
		current: 0,
	}
}

func parseExpression(lexer *Lexer, backtrackToken *Token) (Expression, error) {
	var name string
	args := make([]string, 2)

	hasArgs := true

	if backtrackToken != nil {
		name += backtrackToken.tokenValue
	}

	for {
		token, err := lexer.NextToken()

		if err != nil {
			panic(err)
		}

		if token.tokenType == LeftParen {
			break
		}

		if token.tokenType != FuncName {
			fmt.Errorf("expected expression name")
			return Expression{valid: false, errorTokenType: token.tokenType}, nil
		}

		name += token.tokenValue
	}

	for {
		token, err := lexer.NextToken()

		if err != nil {
			panic(err)
		}

		if token.tokenType == RightParen {
			hasArgs = false
			break
		}

		if token.tokenType == Comma {
			break
		}

		if token.tokenType == FuncName {
			return parseExpression(lexer, &token)
		}

		if token.tokenType != Argument {
			fmt.Errorf("expected argument or comma")
			return Expression{valid: false, errorTokenType: token.tokenType}, nil
		}

		args[0] += token.tokenValue
	}

	for hasArgs {
		token, err := lexer.NextToken()

		if err != nil {
			panic(err)
		}

		if token.tokenType == RightParen {
			break
		}

		if token.tokenType != Argument {
			fmt.Errorf("expected argument or ')'")
			return Expression{valid: false, errorTokenType: token.tokenType}, nil
		}

		args[1] += token.tokenValue
	}

	result := Expression{name: name, args: args, valid: true}
	fmt.Print("parsed expression: ")
	fmt.Println(result)

	return result, nil
}

func (lexer *Lexer) NextToken() (Token, error) {
	if lexer.current >= len(lexer.input) {
		return Token{tokenType: EOF, tokenValue: ""}, nil
	}

	switch currentChar := lexer.input[lexer.current]; currentChar {
	case '(':
		lexer.current++
		return Token{tokenType: LeftParen, tokenValue: "("}, nil
	case ')':
		lexer.current++
		return Token{tokenType: RightParen, tokenValue: ")"}, nil
	case ',':
		lexer.current++
		return Token{tokenType: Comma, tokenValue: ","}, nil
	default:
		lexer.current++
		if unicode.IsLetter(rune(currentChar)) || rune(currentChar) == '\'' {
			return Token{tokenType: FuncName, tokenValue: string(currentChar)}, nil
		} else if unicode.IsDigit(rune(currentChar)) {
			return Token{tokenType: Argument, tokenValue: string(currentChar)}, nil
		} else {
			return Token{tokenType: IllegalChar, tokenValue: string(currentChar)}, nil
		}
	}
}

func testLexer(lexer *Lexer) {
	for {
		token, err := lexer.NextToken()

		if err != nil {
			panic(err)
		}

		fmt.Println(token)

		if token.tokenType == EOF {
			return
		}
	}
}

func evalExpression(expression Expression) (int, bool, ExpressionType) {
	if strings.HasSuffix(expression.name, "mul") {
		arg0, err := strconv.Atoi(expression.args[0])
		if err != nil {
			panic(err)
		}
		arg1, err := strconv.Atoi(expression.args[1])
		if err != nil {
			panic(err)
		}

		result := arg0 * arg1
		return result, true, Mul
	} else if strings.HasSuffix(expression.name, "do") {
		return 0, true, Do
	} else if strings.HasSuffix(expression.name, "don't") {
		return 0, true, Dont
	}

	fmt.Printf("could not eval expr '%s'\n", expression.name)
	return 0, false, Invalid
}

func evalMuls(expressions []Expression) (int, int) {
	sum := 0
	numberExpressions := 0

	isEnabled := true

	for _, expr := range expressions {
		result, doAdd, expressionType := evalExpression(expr)

		if expressionType == Mul && doAdd && isEnabled {
			sum += result
			numberExpressions += 1
		} else if expressionType == Do {
			isEnabled = true
		} else if expressionType == Dont {
			isEnabled = false
		}
	}

	return sum, numberExpressions
}

func main() {
	//lines := common.ReadLinesFromFile("./src/03/testinput.txt")
	lines := common.ReadLinesFromFile("./src/03/input.txt")
	input := strings.Join(lines, "")

	fmt.Println("input")
	fmt.Println(input)

	lexer1 := NewLexer(input)
	testLexer(lexer1)

	lexer2 := NewLexer(input)

	expressions := make([]Expression, 10)
	for {
		expr, err := parseExpression(lexer2, nil)

		if err != nil {
			panic(err)
		}

		if expr.valid {
			expressions = append(expressions, expr)
		} else if expr.errorTokenType == EOF {
			fmt.Println("EOF found")
			break
		}
	}

	result, numberExpressions := evalMuls(expressions)

	fmt.Printf("number of expressions %d\n", numberExpressions)
	fmt.Printf("sum of muls %d", result)
}
