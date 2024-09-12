package typescript_test

import (
	"context"
	"os"
	"testing"
	"time"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("let a : number = 1;"), typescript.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(program (lexical_declaration (variable_declarator name: (identifier) type: (type_annotation (predefined_type)) value: (number))))",
		n.String(),
	)
}

func TestGrammar_Performance(t *testing.T) {
	assert := assert.New(t)

	bytes, err := os.ReadFile("./testdata/index.d.ts")
	st := time.Now()
	assert.NoError(err)
	n, err := sitter.ParseCtx(context.Background(), bytes, typescript.GetLanguage())
	t.Logf("parse time: %v", time.Since(st))
	assert.NoError(err)
	assert.NotNil(n)
}
