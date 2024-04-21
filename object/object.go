package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/WillBallentine/bark/ast"
)

type ObjectType string

const (
	FUNCTION_OBJ     = "FUNCTION"
	BUILTIN_OBJ      = "BUILTIN"
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	STRING_OBJ       = "STRING"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ARRAY_OBJ        = "ARRAY"
	ERROR_OBJ        = "ERROR"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

type BuiltIn struct {
	Fn BuiltInFunction
}

type BuiltInFunction func(args ...Object) Object

type String struct {
	Value string
}

type Integer struct {
	Value int64
}

type Boolean struct {
	Value bool
}

type Null struct{}

type Return_Value struct {
	Value Object
}

type Array struct {
	Elements []Object
}

type Error struct {
	Message string
}

func (ar *Array) Type() ObjectType { return ARRAY_OBJ }
func (ar *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ar.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (bi *BuiltIn) Type() ObjectType { return BUILTIN_OBJ }
func (bi *BuiltIn) Inspect() string  { return "builtin function" }

func (st *String) Type() ObjectType { return STRING_OBJ }
func (st *String) Inspect() string  { return st.Value }

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("trick")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n")

	return out.String()
}

func (er *Error) Type() ObjectType { return ERROR_OBJ }
func (er *Error) Inspect() string  { return "ERROR: " + er.Message }

func (rv *Return_Value) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *Return_Value) Inspect() string  { return rv.Value.Inspect() }

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
