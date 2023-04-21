package parser

import (
	"testing"

	"github.com/ehsaaniqbal/ma/ast"
	"github.com/ehsaaniqbal/ma/lexer"
)

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"god bless you 5;", 5},
		{"god bless you true;", true},
		{"god bless you foobar;", "foobar"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
        t.Log(l)
		p := New(l)
        t.Log(p)
		program := p.ParseProgram()
        t.Log(program)
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		stmt := program.Statements[0]
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("stmt not *ast.ReturnStatement. got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "god bless you" {
			t.Fatalf("returnStmt.TokenLiteral not 'god bless you', got %q",
				returnStmt.TokenLiteral())
		}
		// if testLiteralExpression(t, returnStmt.ReturnValue, tt.expectedValue) {
		// 	return
		// }
	}
}


func TestLetStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"ma x = 5;", "x", 5},
		{"ma y = true;", "y", true},
		{"ma foobar = y;", "foobar", "y"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
        t.Log(l)
		p := New(l)
        t.Log(p)
		program := p.ParseProgram()
        t.Log(program)
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		// val := stmt.(*ast.LetStatement).Value
		// if !testLiteralExpression(t, val, tt.expectedValue) {
		// 	return
		// }
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "ma" {
		t.Errorf("s.TokenLiteral not 'ma'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
