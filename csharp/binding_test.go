package csharp_test

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/csharp"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("using static System.Math;"), csharp.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(compilation_unit (using_directive (qualified_name (identifier) (identifier))))",
		n.String(),
	)
}

func TestGrammar_Complex(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`public class OrderServices(
    IMediator mediator,
    IOrderQueries queries,
    IIdentityService identityService,
    ILogger<OrderServices> logger) 
	: IOrderQueries
{
    public IMediator Mediator { get; set; } = mediator;
    public ILogger<OrderServices> Logger { get; } = logger;
}`), csharp.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(compilation_unit (class_declaration (modifier) name: (identifier) (parameter_list (parameter type: (identifier) name: (identifier)) (parameter type: (identifier) name: (identifier)) (parameter type: (identifier) name: (identifier)) (parameter type: (generic_name (identifier) (type_argument_list (identifier))) name: (identifier))) (base_list (identifier)) body: (declaration_list (property_declaration (modifier) type: (identifier) name: (identifier) accessors: (accessor_list (accessor_declaration) (accessor_declaration)) value: (identifier)) (property_declaration (modifier) type: (generic_name (identifier) (type_argument_list (identifier))) name: (identifier) accessors: (accessor_list (accessor_declaration)) value: (identifier)))))",
		n.String(),
	)
}
