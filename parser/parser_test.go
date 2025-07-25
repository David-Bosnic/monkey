package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatments(t *testing.T) {
	input := `
    let x = 5;
    let y = 10;
    let foobar = 838383;
`
	l := lexer.New(input) //lexer
	p := New(l)           //parser

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statement) != 3 {
		t.Fatalf("programs.Statment does not contain 3 statments, got: %d", len(program.Statement))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statement[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral != 'let', got: %q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatment, got:%T", letStmt)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s', got: %s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.Value not '%s', got: %s", name, letStmt.Name.Value)
		return false
	}
	return true
}
