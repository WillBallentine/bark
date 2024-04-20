package evaluator

import (
	"github.com/WillBallentine/bark/lexer"
	"github.com/WillBallentine/bark/object"
	"github.com/WillBallentine/bark/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}
func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return Eval(program, env)
}
func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
		return false
	}
	return true
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"goodboi", true},
		{"badboi", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"goodboi == goodboi", true},
		{"badboi == badboi", true},
		{"goodboi == badboi", false},
		{"goodboi != badboi", true},
		{"badboi != goodboi", true},
		{"(1 < 2) == goodboi", true},
		{"(1 < 2) == badboi", false},
		{"(1 > 2) == goodboi", false},
		{"(1 > 2) == badboi", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}

func testBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!badboi", true},
		{"!5", false},
		{"!!goodboi", true},
		{"!!badboi", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"borkf (goodboi) { 10 }", 10},
		{"borkf (badboi) { 10 }", nil},
		{"borkf (1) { 10 }", 10},
		{"borkf (1 < 2) { 10 }", 10},
		{"borkf (1 > 2) { 10 }", nil},
		{"borkf (1 > 2) { 10 } woofwise { 20 }", 20},
		{"borkf (1 < 2) { 10 } woofwise { 20 }", 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"fetchit 10;", 10},
		{"fetchit 10; 9;", 10},
		{"fetchit 2 * 5; 9;", 10},
		{"9; fetchit 2 * 5; 9;", 10},
		{
			`
			borkf (10 > 1) {
				borkf (10 > 1) {
					fetchit 10;
						}
				fetchit 1;
					}
			`,
			10,
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + goodboi;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + goodboi; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-goodboi",
			"unknown operator: -BOOLEAN",
		},
		{
			"goodboi + badboi;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; goodboi + badboi; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"borkf (10 > 1) { goodboi + badboi; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
			borkf (10 > 1) {
				borkf (10 > 1) {
					fetchit goodboi + badboi;
						}
				fetchit 1;
				}
				`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"toy a = 5; a;", 5},
		{"toy a = 5 * 5; a;", 25},
		{"toy a = 5; toy b = a; b;", 5},
		{"toy a = 5; toy b = a; toy c = a + b + 5; c;", 15},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}
